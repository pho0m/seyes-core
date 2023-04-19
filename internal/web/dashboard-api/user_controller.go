package dashboardAPI

import (
	"encoding/json"
	"net/http"
	core "seyes-core/internal/core/dashboard"
	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"
	auth "seyes-core/internal/web/common/auth"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// UserController defines handler for http protocol
type UserController struct {
	db *gorm.DB
	common.BaseRender
	sc *service.Container
	auth *auth.Authenticator
}

// NewUserController creates a new WebHandler
func NewUserController(sc *service.Container) *UserController {
	return &UserController{
		db: sc.DB,
		sc: sc,
		auth: sc.Auth.(*auth.Authenticator),
	}
}

// IndexUserHandler endpoint for get all User
func (h *UserController) IndexUserHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)

	// if ctx == nil {
	// 	helper.ReturnError(w, "err", "error get all User", http.StatusBadRequest)
	// 	return
	// }


	ps := helper.ParsingQueryString(r.URL.Query())

	Users, err := core.GetAllUser(h.db, &core.UserFilter{
		ID:      ps.ID,
		Page:    ps.Page,
		OrderBy: ps.OrderBy,
		SortBy:  ps.SortBy,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get all User", http.StatusBadRequest)
		return
	}

	h.JSON(w, Users)
}

// GetUserHandler endpoint for get a User
func (h *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	res, err := core.GetUser(h.db, &helper.UrlParams{
		ID: ps.ID,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get a User", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// CreateUserHandler endpoint for create a User
func (h *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var ps core.UserParams

	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	res, err := core.CreateUser(h.db, &ps)

	if err != nil {
		helper.ReturnError(w, err, "error create a User", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// UpdateUserHandler endpoint for update a User
func (h *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))
	var req core.UserParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	req.ID = ps.ID
	res, err := core.UpdatedUser(h.db, &req)

	if err != nil {
		helper.ReturnError(w, err, "error update a User", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// DeleteUserHandler endpoint for delete a User
func (h *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	err := core.DeletedUser(h.db, ps.ID)

	if err != nil {
		helper.ReturnError(w, err, "error delete User", http.StatusBadRequest)
		return
	}

	h.JSON(w, "ok")
}
