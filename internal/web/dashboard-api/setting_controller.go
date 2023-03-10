package dashboardAPI

import (
	"encoding/json"
	"net/http"
	core "seyes-core/internal/core/dashboard"
	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

//*FIXME please refactor comment function
//*      and handle errors

// SettingController defines handler for http protocol
type SettingController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewSettingsController creates a new WebHandler
func NewSettingsController(sc *service.Container) *SettingController {
	return &SettingController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// GetSettingsHandler endpoint for get a Settings
func (h *SettingController) GetSettingsHandler(w http.ResponseWriter, r *http.Request) {

	res, err := core.GetSetting(h.db, &helper.UrlParams{
		ID: 1,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get a Settings", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// UpdateSettingsHandler endpoint for update a Settings
func (h *SettingController) UpdateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))
	var req core.SettingsParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	req.ID = ps.ID
	res, err := core.UpdatedSettings(h.db, &req)

	if err != nil {
		helper.ReturnError(w, err, "error update a Settings", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}
