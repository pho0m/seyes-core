package dashboardAPI

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"

	auth "seyes-core/internal/web/common/auth"

	"seyes-core/internal/service"
	"seyes-core/internal/web/common"
)

// AuthController defines handler for http protocol
type AuthController struct {
	db *gorm.DB
	common.BaseRender
	auth   *auth.Authenticator
	sc     *service.Container
}

// NewAuthController creates a new WebHandler
func NewAuthController(sc *service.Container) *AuthController {
	return &AuthController{
		db:     sc.DB,
		sc:     sc,
		auth:   sc.Auth.(*auth.Authenticator),
	}
}

// Login endpoint login and return access token if login success
func (h *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var ps auth.User
	err := json.NewDecoder(r.Body).Decode(&ps)

	spew.Dump(ps)

	if err != nil {

		h.Error(w, err, "decode param error", http.StatusInternalServerError)
		return
	}
	user, err := auth.GetUserByEmail(h.db, ps.Email)

	if err != nil {

		h.Error(w, err, "user doesn't exist", http.StatusUnauthorized)
		return
	}

	spew.Dump(ps.Password)
	spew.Dump(user.Password)

	if !h.auth.CheckPassword(user.Password, ps.Password) {
spew.Dump("in if")

		h.Error(w, errors.New("uauthorized"), "uauthorized", http.StatusUnauthorized)
		return
	}

	token,  _, err := h.auth.SignJWTToken( map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"name":       user.FirstName,
	}, nil)

	if err != nil {
		h.Error(w, err, "sign jwt error", http.StatusUnauthorized)
		return
	}
	res := map[string]interface{}{
		"access_token":  token,
	}

	h.JSON(w, res)
}

// Me endpoint for get user credentail by token
func (h *AuthController) Me(w http.ResponseWriter, r *http.Request) {

	var err error
	token := r.Header.Get("X-Token")
	spew.Dump(token)

	if token == "" {
		h.Error(w, errors.New("required token"), "unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := h.auth.VerifyJWTToken(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userCtx, err := h.auth.ParseClaimsContext(h.db, claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err := auth.GetUserFromCtx(h.db, userCtx.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	h.JSON(w, user)
}

