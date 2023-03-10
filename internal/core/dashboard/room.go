package core

import (
	"seyes-core/internal/helper"
	"time"

	"gorm.io/gorm"

	mo "seyes-core/internal/model/room"
)

// RoomFilter define room filter
type RoomFilter struct {
	Page    int64    `json:"page"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Active  []string `json:"active"`
	OrderBy string   `json:"order_by"`
	SortBy  string   `json:"sort_by"`
}

// RoomParams define params for create room
type RoomParams struct {
	ID     int64  `json:"id"`
	Label  string `json:"label"`
	CamURL string `json:"cam_url"`
	Status string `json:"status"`
	Active string `json:"active"`
}

// GetAllRoom get all room product
func GetAllRoom(db *gorm.DB, filter *RoomFilter) (map[string]interface{}, error) {
	// var resPr []RoomParams
	var room []mo.Room
	var resRooms []RoomParams

	dbx := db.Model(&mo.Room{})
	pg := helper.FormatWebPaginate(dbx, filter.Page)

	if err := pg.DB.Find(&room).Error; err != nil {
		return nil, err
	}

	for _, r := range room {
		resRooms = append(resRooms, RoomParams{
			ID:     int64(r.ID),
			Label:  r.Label,
			CamURL: r.CamURL,
			Status: r.Status,
			Active: r.Active,
		})
	}

	if len(resRooms) == 0 {
		resRooms = []RoomParams{}
	}

	return map[string]interface{}{
		"items":       resRooms,
		"page":        pg.Page,
		"total_pages": pg.TotalPages,
		"total_count": pg.TotalCount,
	}, nil
}

// GetRoom get a room by room id
func GetRoom(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	var room mo.Room

	if err := db.Where("id = ?", ps.ID).
		Where("deleted_at IS NULL").
		First(&room).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":      room.ID,
		"label":   room.Label,
		"cam_url": room.CamURL,
		"status":  room.Status,
		"active":  room.Active,
	}

	return res, nil
}

// CreateRoom create a room
func CreateRoom(db *gorm.DB, ps *RoomParams) (map[string]interface{}, error) {
	room := &mo.Room{
		Label:  ps.Label,
		CamURL: ps.CamURL,
		Status: ps.Status,
		Active: ps.Active,
	}

	if err := db.Create(&room).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":      room.ID,
		"label":   room.Label,
		"cam_url": room.CamURL,
		"status":  room.Status,
		"active":  room.Active,
	}

	return res, nil
}

// UpdatedRoom update a room
func UpdatedRoom(db *gorm.DB, ps *RoomParams) (map[string]interface{}, error) {
	var room mo.Room

	if err := db.Where("id = ?", ps.ID).
		First(&room).Error; err != nil {
		return nil, err
	}

	room.Label = ps.Label
	room.CamURL = ps.CamURL
	room.Status = ps.Status
	room.Active = ps.Active

	if err := db.Save(&room).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":      room.ID,
		"label":   room.Label,
		"cam_url": room.CamURL,
		"status":  room.Status,
		"active":  room.Active,
	}

	return res, nil
}

// DeletedRoom delete a room
func DeletedRoom(db *gorm.DB, id int64) error {
	var room mo.Room
	t := time.Now()

	if err := db.Where("id = ?", id).
		First(&room).Error; err != nil {
		return err
	}
	room.DeletedAt = &t

	if err := db.Delete(&room).Error; err != nil {
		return err
	}

	return nil
}