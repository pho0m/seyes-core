package common

import (
	"context"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"

	"seyes-core/internal/service"
)

const RoleSuperAdmin = "super_admin"
const RoleAdmin = "admin"
const RoleManager = "manager"
const RoleCashier = "cashier"
const RoleIvtManager = "inventory_manager"
const RoleOwner = "owner"
const RoleAdminDelivery = "admin_delivery"
const RoleAccountant = "accountant"
const RoleStaff = "staff"

// MiddlewareAPI defines handler for http protocol
type MiddlewareAPI struct {
	db   *gorm.DB
	auth *Authenticator
}

// NewApiMiddleware creates a new NewApiMiddleware
func NewApiMiddleware(sc *service.Container) *MiddlewareAPI {

	return &MiddlewareAPI{
		db:   sc.DB,
		auth: sc.Auth.(*Authenticator),
	}
}

// Middleware define middleware for root path api
func (a *MiddlewareAPI) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()

			if err != nil {
				spew.Dump("error middleware")
				panic(err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// MiddlewareCore define verify token before access api core
func (a *MiddlewareAPI) MiddlewareCore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Token")

		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		} else {
			claims, err := a.auth.VerifyJWTToken(token)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			user, err := a.auth.ParseClaimsContext(a.db, claims)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_info", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
