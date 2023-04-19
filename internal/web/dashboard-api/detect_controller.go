package dashboardAPI

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	core "seyes-core/internal/core/dashboard"
	noti "seyes-core/internal/core/notifications"

	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// DetectController defines handler for http protocol
type DetectController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewDetectController creates a new WebHandler
func NewDetectController(sc *service.Container) *DetectController {
	return &DetectController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// GetDetectHandler endpoint for get all Detect
func (h *DetectController) GetDetectHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)
	var urlSeyesCam = os.Getenv("SEYES_CAM_URL") + "/image"
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	resRoom, err := core.GetRoom(h.db, &helper.UrlParams{
		ID: ps.ID,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get a room", http.StatusBadRequest)
		return
	}

	rr, _ := json.Marshal(resRoom)
	var room core.RoomParams
	json.Unmarshal(rr, &room)

	resFromSCAM, err := http.Get(urlSeyesCam + "/" + room.UuidCam + "/channel" + "/0") //+ "/" + ps.Uuid + "/channel" + ps.Channel
	if err != nil {
		helper.ReturnError(w, err, "error cannot get image form seyes cam", http.StatusBadRequest)
		return
	}

	responseData, err := ioutil.ReadAll(resFromSCAM.Body)
	if err != nil {
		helper.ReturnError(w, err, "error cannot read seyescam data", http.StatusBadRequest)
		return
	}
	var dp core.DetectionParams
	json.Unmarshal(responseData, &dp)

	var jsonStr = []byte(`{"image":` + `"` + dp.ImageData + `"` + `}`)
	req, err := http.NewRequest("POST", os.Getenv("SEYES_DETECT_URL")+"/detect", bytes.NewBuffer(jsonStr))
	if err != nil {
		helper.ReturnError(w, err, "error Detection form seyes detection", http.StatusBadRequest)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.ReturnError(w, err, "error client do", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &dp)

	personCount, _ := strconv.Atoi(dp.PersonCount)
	comCount, _ := strconv.Atoi(dp.ConOnCount)
	status := "undetected"

	if dp.Accurency != "0%" {
		status = "detected"
	}

	resReports, err := core.CreateReport(h.db, &core.ReportsParams{
		PersonCont: int64(personCount),
		ComOnCount: int64(comCount),
		Status:     status,
		Image:      dp.ImageData,
		Accurency:  dp.Accurency,
		RoomLabel:  room.Label,
		ReportTime: dp.Time,
		ReportDate: dp.Date,
	})

	if err != nil {
		helper.ReturnError(w, err, "error create report", http.StatusBadRequest)
		return
	}

	rreport, _ := json.Marshal(resReports)
	var report core.ReportsParams
	json.Unmarshal(rreport, &report)

	notifyData := noti.NotifyParamV2{
		Uuid:      room.UuidCam,
		Image:     report.Image,
		ID:        report.ID,
		Person:    strconv.Itoa(int(report.PersonCont)),
		ComOn:     strconv.Itoa(int(report.ComOnCount)),
		UploadAt:  report.ReportDate,
		Time:      report.ReportTime,
		Accurency: report.Accurency,
	}

	if notifyData.Accurency != "0%" {
		err = noti.SendToLineNotifyV2(&notifyData)

		if err != nil {
			helper.ReturnError(w, err, "error sent notify", http.StatusBadRequest)
			return
		}
	}

	h.JSON(w, notifyData)
}

// UpdateModelFileHandler endpoint for update a new model
func (h *DetectController) UpdateModelFileHandler(w http.ResponseWriter, r *http.Request) {

	h.JSON(w, "wip")
}
