package core

import (
	"seyes-core/internal/helper"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	mo "seyes-core/internal/model/user"
)

// UserFilter define user filter
type UserFilter struct {
	Page    int64    `json:"page"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Active  []string `json:"active"`
	OrderBy string   `json:"order_by"`
	SortBy  string   `json:"sort_by"`
}

// UserParams define params for create user
type UserParams struct {
	ID        int64  `json:"id"`
	Active    bool   `json:"active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	Password  string `json:"password"` //FIXME
	Email     string `json:"email" gorm:"uniqueIndex:idx_user"`
}

// GetAllUser get all user product
func GetAllUser(db *gorm.DB, filter *UserFilter) (map[string]interface{}, error) {
	// var resPr []RoomParams
	var users []mo.User
	var resUsers []UserParams

	dbx := db.Model(&mo.User{})
	pg := helper.FormatWebPaginate(dbx, filter.Page)

	if err := pg.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	for _, u := range users {
		resUsers = append(resUsers, UserParams{
			ID:        int64(u.ID),
			Active:    u.Active,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Tel:       u.Tel,
			Email:     u.Email,
		})
	}

	if len(resUsers) == 0 {
		resUsers = []UserParams{}
	}

	return map[string]interface{}{
		"items":       resUsers,
		"page":        pg.Page,
		"total_pages": pg.TotalPages,
		"total_count": pg.TotalCount,
	}, nil
}

// GetUser get a user by user id
func GetUser(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	var user mo.User

	if err := db.Where("id = ?", ps.ID).
		Where("deleted_at IS NULL").
		First(&user).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":         user.ID,
		"active":     user.Active,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"tel":        user.Tel,
		"email":      user.Email,
	}

	return res, nil
}

// CreateUser create a user
func CreateUser(db *gorm.DB, ps *UserParams) (map[string]interface{}, error) {
	bcPass, _ := bcrypt.GenerateFromPassword([]byte(ps.Password), 14)
	
	user := &mo.User{
		Active:    ps.Active,
		FirstName: ps.FirstName,
		LastName:  ps.LastName,
		Tel:       ps.Tel,
		Password: string(bcPass),
		Email:     ps.Email,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":         user.ID,
		"active":     user.Active,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"tel":        user.Tel,
		"email":      user.Email,
	}

	return res, nil
}

// UpdatedUser update a room
func UpdatedUser(db *gorm.DB, ps *UserParams) (map[string]interface{}, error) {
	var user mo.User

	if err := db.Where("id = ?", ps.ID).
		First(&user).Error; err != nil {
		return nil, err
	}

	user.Active = ps.Active
	user.FirstName = ps.FirstName
	user.LastName = ps.LastName
	user.Tel = ps.Tel

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":         user.ID,
		"active":     user.Active,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"tel":        user.Tel,
		"email":      user.Email,
	}

	return res, nil
}

// DeletedUser delete a user
func DeletedUser(db *gorm.DB, id int64) error {
	var user mo.User
	// t := time.Now()

	if err := db.Where("id = ?", id).
		First(&user).Error; err != nil {
		return err
	}
	// room.DeletedAt = &t

	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
