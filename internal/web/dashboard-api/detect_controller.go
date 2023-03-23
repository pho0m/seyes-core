package dashboardAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"

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

	uuid := chi.URLParam(r, "uuid")
	channel := chi.URLParam(r, "channel")

	resFromSCAM, err := http.Get(urlSeyesCam + "/" + uuid + "/channel" + "/" + channel) //+ "/" + ps.Uuid + "/channel" + ps.Channel
	if err != nil {
		helper.ReturnError(w, err, "error get all Detect", http.StatusBadRequest)
		return
	}

	responseData, err := ioutil.ReadAll(resFromSCAM.Body)
	if err != nil {
		helper.ReturnError(w, err, "error get all Detect", http.StatusBadRequest)
		return
	}
	var responseObject DetectionParams
	json.Unmarshal(responseData, &responseObject)

	var jsonStr = []byte(`{"image":` + `"` + responseObject.ImageData + `"` + `}`)
	req, err := http.NewRequest("POST", "http://202.44.35.76:9094/detect", bytes.NewBuffer(jsonStr))
	if err != nil {
		helper.ReturnError(w, err, "error get all Detect", http.StatusBadRequest)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.ReturnError(w, err, "error get all Detect", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	json.Unmarshal(body, &responseObject)

	// de, err := core.GetImageFromSeyesCam(&core.DetectionParams{ //detects, err :=
	// 	Uuid:    uuid,
	// 	Channel: channel,
	// })

	// if err != nil {
	// helper.ReturnError(w, err, "error get all Detect", http.StatusBadRequest)
	// return
	// }

	h.JSON(w, responseObject)
}

// // GetDetectHandler endpoint for get a Detect
// func (h *DetectController) GetDetectHandler(w http.ResponseWriter, r *http.Request) {
// 	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

// 	res, err := core.GetDetect(h.db, &helper.UrlParams{
// 		ID: ps.ID,
// 	})

// 	if err != nil {
// 		helper.ReturnError(w, err, "error get a Detect", http.StatusBadRequest)
// 		return
// 	}

// 	h.JSON(w, res)
// }

// // CreateDetectHandler endpoint for create a Detect
// func (h *DetectController) CreateDetectHandler(w http.ResponseWriter, r *http.Request) {
// 	var ps core.DetectParams

// 	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
// 		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
// 		return
// 	}
// 	res, err := core.CreateDetect(h.db, &ps)

// 	if err != nil {
// 		helper.ReturnError(w, err, "error create a Detect", http.StatusBadRequest)
// 		return
// 	}

// 	h.JSON(w, res)
// }

// // UpdateDetectHandler endpoint for update a Detect
// func (h *DetectController) UpdateDetectHandler(w http.ResponseWriter, r *http.Request) {
// 	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))
// 	var req core.DetectParams

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
// 		return
// 	}
// 	req.ID = ps.ID
// 	res, err := core.UpdatedDetect(h.db, &req)

// 	if err != nil {
// 		helper.ReturnError(w, err, "error update a Detect", http.StatusBadRequest)
// 		return
// 	}

// 	h.JSON(w, res)
// }

// // DeleteDetectHandler endpoint for delete a Detect
// func (h *DetectController) DeleteDetectHandler(w http.ResponseWriter, r *http.Request) {
// 	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

// 	err := core.DeletedDetect(h.db, ps.ID)

// 	if err != nil {
// 		helper.ReturnError(w, err, "error delete Detect", http.StatusBadRequest)
// 		return
// 	}

// 	h.JSON(w, "ok")
// }
