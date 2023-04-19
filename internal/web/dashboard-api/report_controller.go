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

// ReportController defines handler for http protocol
type ReportController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewReportController creates a new WebHandler
func NewReportController(sc *service.Container) *ReportController {
	return &ReportController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// IndexReportHandler endpoint for get all Report
func (h *ReportController) AnalyticsReportsHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)

	ps := helper.ParsingQueryString(r.URL.Query())

	reports, err := core.AnalyticsReports(h.db, &core.ReportsFilter{
		ID:      ps.ID,
		Page:    ps.Page,
		OrderBy: ps.OrderBy,
		SortBy:  ps.SortBy,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get all AnalyticsReports Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, reports)
}

// IndexReportHandler endpoint for get all Report
func (h *ReportController) IndexReportHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)

	ps := helper.ParsingQueryString(r.URL.Query())

	reports, err := core.GetAllReports(h.db, &core.ReportsFilter{
		ID:      ps.ID,
		Page:    ps.Page,
		OrderBy: ps.OrderBy,
		SortBy:  ps.SortBy,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get all Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, reports)
}

// GetReportHandler endpoint for get a Report
func (h *ReportController) GetReportHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	res, err := core.GetReport(h.db, &helper.UrlParams{
		ID: ps.ID,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get a Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// CreateReportHandler endpoint for create a Report
func (h *ReportController) CreateReportHandler(w http.ResponseWriter, r *http.Request) {
	var ps core.ReportsParams

	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	res, err := core.CreateReport(h.db, &ps)

	if err != nil {
		helper.ReturnError(w, err, "error create a Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// UpdateReportHandler endpoint for update a Report
func (h *ReportController) UpdateReportHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))
	var req core.ReportsParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	req.ID = ps.ID
	res, err := core.UpdatedReport(h.db, &req)

	if err != nil {
		helper.ReturnError(w, err, "error update a Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// DeleteReportHandler endpoint for delete a Report
func (h *ReportController) DeleteReportHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	err := core.DeletedReport(h.db, ps.ID)

	if err != nil {
		helper.ReturnError(w, err, "error delete Report", http.StatusBadRequest)
		return
	}

	h.JSON(w, "ok")
}
