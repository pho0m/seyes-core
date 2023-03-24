package core

import (
	"seyes-core/internal/helper"

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
	ID             int64  `json:"id"`
	Label          string `json:"label"`
	CamURL         string `json:"cam_url"`
	UuidCam        string `json:"uuid_cam"`
	Status         string `json:"status"`
	Active         bool   `json:"active"`
	MqttTopicLamp1 string `json:"mqtt_topic_lamp_1"`
	MqttTopicLamp2 string `json:"mqtt_topic_lamp_2"`
	MqttTopicLamp3 string `json:"mqtt_topic_lamp_3"`
	MqttTopicLamp4 string `json:"mqtt_topic_lamp_4"`
	MqttTopicLamp5 string `json:"mqtt_topic_lamp_5"`
	MqttTopicLamp6 string `json:"mqtt_topic_lamp_6"`
	MqttTopicDoor  string `json:"mqtt_topic_door"`
	MqttTopicAir   string `json:"mqtt_topic_air"`
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
			ID:             int64(r.ID),
			Label:          r.Label,
			CamURL:         r.CamURL,
			Status:         r.Status,
			Active:         r.Active,
			MqttTopicLamp1: r.MqttTopicLamp1,
			MqttTopicLamp2: r.MqttTopicLamp2,
			MqttTopicLamp3: r.MqttTopicLamp3,
			MqttTopicLamp4: r.MqttTopicLamp4,
			MqttTopicLamp5: r.MqttTopicLamp5,
			MqttTopicLamp6: r.MqttTopicLamp6,
			MqttTopicDoor:  r.MqttTopicDoor,
			MqttTopicAir:   r.MqttTopicAir,
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
		"id":                room.ID,
		"label":             room.Label,
		"cam_url":           room.CamURL,
		"uuid_cam":          room.UudiCam,
		"status":            room.Status,
		"active":            room.Active,
		"mqtt_topic_lamp_1": room.MqttTopicLamp1,
		"mqtt_topic_lamp_2": room.MqttTopicLamp2,
		"mqtt_topic_lamp_3": room.MqttTopicLamp3,
		"mqtt_topic_lamp_4": room.MqttTopicLamp4,
		"mqtt_topic_lamp_5": room.MqttTopicLamp5,
		"mqtt_topic_lamp_6": room.MqttTopicLamp6,
		"mqtt_topic_door":   room.MqttTopicDoor,
		"mqtt_topic_air":    room.MqttTopicAir,
	}

	return res, nil
}

// CreateRoom create a room
func CreateRoom(db *gorm.DB, ps *RoomParams) (map[string]interface{}, error) {
	room := &mo.Room{
		Label:          ps.Label,
		CamURL:         ps.CamURL,
		UudiCam:        ps.UuidCam,
		Status:         ps.Status,
		Active:         ps.Active,
		MqttTopicLamp1: ps.MqttTopicLamp1,
		MqttTopicLamp2: ps.MqttTopicLamp2,
		MqttTopicLamp3: ps.MqttTopicLamp3,
		MqttTopicLamp4: ps.MqttTopicLamp4,
		MqttTopicLamp5: ps.MqttTopicLamp5,
		MqttTopicLamp6: ps.MqttTopicLamp6,
		MqttTopicDoor:  ps.MqttTopicDoor,
		MqttTopicAir:   ps.MqttTopicAir,
	}

	if err := db.Create(&room).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                room.ID,
		"label":             room.Label,
		"cam_url":           room.CamURL,
		"uuid_cam":          room.UudiCam,
		"status":            room.Status,
		"active":            room.Active,
		"mqtt_topic_lamp_1": room.MqttTopicLamp1,
		"mqtt_topic_lamp_2": room.MqttTopicLamp2,
		"mqtt_topic_lamp_3": room.MqttTopicLamp3,
		"mqtt_topic_lamp_4": room.MqttTopicLamp4,
		"mqtt_topic_lamp_5": room.MqttTopicLamp5,
		"mqtt_topic_lamp_6": room.MqttTopicLamp6,
		"mqtt_topic_door":   room.MqttTopicDoor,
		"mqtt_topic_air":    room.MqttTopicAir,
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
	room.UudiCam = ps.UuidCam

	room.MqttTopicLamp1 = ps.MqttTopicLamp1
	room.MqttTopicLamp2 = ps.MqttTopicLamp2
	room.MqttTopicLamp3 = ps.MqttTopicLamp3
	room.MqttTopicLamp4 = ps.MqttTopicLamp4
	room.MqttTopicLamp5 = ps.MqttTopicLamp5
	room.MqttTopicLamp6 = ps.MqttTopicLamp6
	room.MqttTopicDoor = ps.MqttTopicDoor
	room.MqttTopicAir = ps.MqttTopicAir

	if err := db.Save(&room).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                room.ID,
		"label":             room.Label,
		"cam_url":           room.CamURL,
		"uuid_cam":          room.UudiCam,
		"status":            room.Status,
		"active":            room.Active,
		"mqtt_topic_lamp_1": room.MqttTopicLamp1,
		"mqtt_topic_lamp_2": room.MqttTopicLamp2,
		"mqtt_topic_lamp_3": room.MqttTopicLamp3,
		"mqtt_topic_lamp_4": room.MqttTopicLamp4,
		"mqtt_topic_lamp_5": room.MqttTopicLamp5,
		"mqtt_topic_lamp_6": room.MqttTopicLamp6,
		"mqtt_topic_door":   room.MqttTopicDoor,
		"mqtt_topic_air":    room.MqttTopicAir,
	}

	return res, nil
}

// DeletedRoom delete a room
func DeletedRoom(db *gorm.DB, id int64) error {
	var room mo.Room
	// t := time.Now()

	if err := db.Where("id = ?", id).
		First(&room).Error; err != nil {
		return err
	}
	// room.DeletedAt = &t

	if err := db.Delete(&room).Error; err != nil {
		return err
	}

	return nil
}
