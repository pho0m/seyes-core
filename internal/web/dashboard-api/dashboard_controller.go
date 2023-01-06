package dashboardAPI

import (
	"net/http"
	rd "seyes-core/internal/core"
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

// NewDashboardController creates a new ShopWebHandler
func NewDashboardController(sc *service.Container) *DashboardController {
	return &DashboardController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// Notify endpoint for create shop
func (c *DashboardController) Notify(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024) // 2 Mb
	file, handler, err := r.FormFile("photo")

	if err != nil {
		c.Error(w, err, "error upload photo", http.StatusBadRequest)
		return
	}

	ps := helper.ParsingQueryUpload(r.Form)

	defer file.Close()
	u := helper.ParsingUploadFileParams(file, handler)

	data := rd.NotifyParam{
		Photo:    u.File,
		ID:       ps.ID,
		Person:   ps.Person,
		ComOn:    ps.ComOn,
		UploadAt: ps.UploadAt,
		Time:     ps.Time,
	}

	err = rd.SendToLineNotify(&data)

	if err != nil {
		c.Error(w, err, "send to error", http.StatusInternalServerError)
		return
	}

	c.JSON(w, "notify !")
}

// var ps rc.ShopParams
// err := json.NewDecoder(r.Body).Decode(&ps)

// if err != nil {
// 	c.Error(w, err, "decode param error", http.StatusBadRequest)
// 	return
// }
// br, err := rc.CreateShop(c.db, &ps)

// if err != nil {
// 	c.Error(w, err, "create error", http.StatusInternalServerError)
// 	return
// }
