package dashboardAPI

import (
	"encoding/json"
	"log"
	"net/http"
	dt "seyes-core/internal/core/detections"
	noti "seyes-core/internal/core/notifications"

	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"

	"gorm.io/gorm"
)

// DashboardController defines handler for http protocol
type DashboardController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewDashboardController creates a new DashboardController
func NewDashboardController(sc *service.Container) *DashboardController {
	return &DashboardController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// HealthCheck endpoint
func (c *DashboardController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	c.JSON(w, "hi ! from seyes server")
}

// Notify endpoint for notify in line
func (c *DashboardController) Notify(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 4*1024*1024) // 4 Mb
	file, handler, err := r.FormFile("photo")

	if err != nil {
		c.Error(w, err, "error upload photo", http.StatusBadRequest)
		return
	}

	ps := helper.ParsingQueryUpload(r.Form)

	defer file.Close()
	u := helper.ParsingUploadFileParams(file, handler)

	data := noti.NotifyParam{
		Photo:     u.File,
		ID:        ps.ID,
		Person:    ps.Person,
		ComOn:     ps.ComOn,
		UploadAt:  ps.UploadAt,
		Time:      ps.Time,
		Accurency: ps.Accurency,
	}

	res, err := noti.SendToLineNotify(&data)

	if err != nil {
		c.Error(w, err, "send to error", http.StatusInternalServerError)
		return
	}

	c.JSON(w, res)
}

// ReadModelFile endpoint read tensorflow model for object detection
func (c *DashboardController) ReadModelFile(w http.ResponseWriter, r *http.Request) {

	res, err := dt.ReadTensorflowModel()

	if err != nil {
		c.Error(w, err, "cannot read file", http.StatusInternalServerError)
		return
	}
	var m interface{}

	err = json.Unmarshal([]byte(res), &m)
	if err != nil {
		c.Error(w, err, "cannot Unmarshal JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(res)); err != nil {
		log.Println("Cannot write a response:", err.Error())
	}

	c.JSON(w, res)
}
