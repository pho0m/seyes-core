package dashboardAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	core "seyes-core/internal/core/dashboard"
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

type DetectionParams struct {
	Uuid        string `json:"id"`
	Channel     string `json:"channel"`
	ImageData   string `json:"image"`
	Accurency   string `json:"accuracy"`
	ConOnCount  string `json:"com_on_count"`
	Date        string `json:"date"`
	PersonCount string `json:"person_count"`
	Status      string `json:"status_detec"`
	Time        string `json:"time"`
}

// GetDetectHandler endpoint for get all Detect
func (h *DetectController) GetDetectHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)
	const urlSeyesCam = "http://202.44.35.76:9093/image"
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
	var ro DetectionParams
	json.Unmarshal(responseData, &ro)

	var jsonStr = []byte(`{"image":` + `"` + ro.ImageData + `"` + `}`)
	req, err := http.NewRequest("POST", "http://202.44.35.76:9094/detect", bytes.NewBuffer(jsonStr))
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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &ro)

	personCount, _ := strconv.ParseInt(ro.PersonCount, 10, 64)
	comCount, _ := strconv.ParseInt(ro.ConOnCount, 10, 64)
	acc, _ := strconv.ParseFloat(ro.Accurency, 64)

	resReports, err := core.CreateReport(h.db, &core.ReportsParams{
		PersonCont: personCount,
		ComOnCount: comCount,
		Status:     "detected",
		Image:      ro.ImageData,
		Accurency:  acc,
		RoomLabel:  room.Label,
		ReportTime: ro.Time,
		ReportDate: ro.Date,
	})

	if err != nil {
		helper.ReturnError(w, err, "error create report", http.StatusBadRequest)
		return
	}

	h.JSON(w, resReports)
}

// UpdateModelFileHandler endpoint for update a new model
func (h *DetectController) UpdateModelFileHandler(w http.ResponseWriter, r *http.Request) {
	// ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	// res, err := core.GetDetect(h.db, &helper.UrlParams{
	// 	ID: ps.ID,
	// })

	// if err != nil {
	// 	helper.ReturnError(w, err, "error get a Detect", http.StatusBadRequest)
	// 	return
	// }

	h.JSON(w, "wip")
}
