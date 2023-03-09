package dashboardAPI

import (
	"encoding/json"
	"net/http"
	core "seyes-core/internal/core/dashboard"
	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web/common"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// RoomController defines handler for http protocol
type RoomController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewRoomController creates a new WebHandler
func NewRoomController(sc *service.Container) *RoomController {
	return &RoomController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// IndexRoomHandler endpoint for get all room
func (h *RoomController) IndexRoomHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context().Value("user_info").(*auth.UserInfo)

	ps := helper.ParsingQueryString(r.URL.Query())

	drws, err := core.GetAllRoom(h.db, &core.RoomFilter{
		ID:      ps.ID,
		Page:    ps.Page,
		OrderBy: ps.OrderBy,
		SortBy:  ps.SortBy,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get all room", http.StatusBadRequest)
		return
	}

	h.JSON(w, drws)
}

// GetRoomHandler endpoint for get a room
func (h *RoomController) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	spew.Dump(ps)

	res, err := core.GetRoom(h.db, &helper.UrlParams{
		ID: ps.ID,
	})

	if err != nil {
		helper.ReturnError(w, err, "error get a room", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// CreateRoomHandler endpoint for create a room
func (h *RoomController) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var ps core.RoomParams

	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	res, err := core.CreateRoom(h.db, &ps)

	if err != nil {
		helper.ReturnError(w, err, "error create a room", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// UpdateRoomHandler endpoint for update a room
func (h *RoomController) UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))
	var req core.RoomParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ReturnError(w, err, "error decode params", http.StatusBadRequest)
		return
	}
	req.ID = ps.ID
	res, err := core.UpdatedRoom(h.db, &req)

	if err != nil {
		helper.ReturnError(w, err, "error update a room", http.StatusBadRequest)
		return
	}

	h.JSON(w, res)
}

// DeleteRoomHandler endpoint for delete a room
func (h *RoomController) DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	ps := helper.ParsingQueryString(chi.URLParam(r, "id"))

	err := core.DeletedRoom(h.db, ps.ID)

	if err != nil {
		helper.ReturnError(w, err, "error delete room", http.StatusBadRequest)
		return
	}

	h.JSON(w, "ok")
}
