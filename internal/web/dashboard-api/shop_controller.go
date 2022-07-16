package dashboardAPI

import (
	"mns-core/internal/service"
	"mns-core/internal/web/common"
	"net/http"

	"gorm.io/gorm"
)

// ShopController defines handler for http protocol
type ShopController struct {
	db *gorm.DB
	common.BaseRender
	// auth *auth.Authenticator
	sc *service.Container
}

// NewShopsController creates a new ShopWebHandler
func NewShopsController(sc *service.Container) *ShopController {
	return &ShopController{
		db: sc.DB,
		sc: sc,
		// auth: sc.Auth.(*auth.Authenticator),
	}
}

// NewShop endpoint for create shop
func (c *ShopController) NewShop(w http.ResponseWriter, r *http.Request) {
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

	c.JSON(w, "pass")
}
