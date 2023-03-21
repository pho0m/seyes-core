package core

import (
	"seyes-core/internal/helper"

	"gorm.io/gorm"
)

// FIXME relate with room
// ScheduleFilter define Schedule filter
type ScheduleFilter struct {
	Page    int64    `json:"page"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Active  []string `json:"active"`
	OrderBy string   `json:"order_by"`
	SortBy  string   `json:"sort_by"`
}

// ScheduleParams define params for create Schedule
type ScheduleParams struct {
	ID        int64  `json:"id"`
	StartTime string `json:"start_time"`
	DueTime   string `json:"due_time"`
	Period    string `json:"period"`
	Day       string `json:"day"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
	Label     string `json:"label"`
	Prefix    string `json:"prefix"`
	Active    bool   `json:"active"`
}

// GetAllSchedules get all Schedule product
func GetAllSchedules(db *gorm.DB, filter *ScheduleFilter) (map[string]interface{}, error) {
	// // var resPr []ScheduleParams
	// var Schedule []mo.Schedule
	// var resSchedules []ScheduleParams

	// dbx := db.Model(&mo.Schedule{})
	// pg := helper.FormatWebPaginate(dbx, filter.Page)

	// if err := pg.DB.Find(&Schedule).Error; err != nil {
	// 	return nil, err
	// }

	// for _, r := range Schedule {
	// 	resSchedules = append(resSchedules, ScheduleParams{
	// 		ID:     int64(r.ID),
	// 		Label:  r.Label,
	// 		CamURL: r.CamURL,
	// 		Status: r.Status,
	// 		Active: r.Active,
	// 	})
	// }

	// if len(resSchedules) == 0 {
	// 	resSchedules = []ScheduleParams{}
	// }

	return map[string]interface{}{
		"items":       "resSchedules",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// GetSchedule get a Schedule by Schedule id
func GetSchedule(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	// var Schedule mo.Schedule

	// if err := db.Where("id = ?", ps.ID).
	// 	Where("deleted_at IS NULL").
	// 	First(&Schedule).Error; err != nil {
	// 	return nil, err
	// }

	return map[string]interface{}{
		"items":       "resSchedules",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil

}

// CreateSchedule create a Schedule
func CreateSchedule(db *gorm.DB, ps *ScheduleParams) (map[string]interface{}, error) {
	// Schedule := &mo.Schedule{
	// 	Label:  ps.Label,
	// 	CamURL: ps.CamURL,
	// 	Status: ps.Status,
	// 	Active: ps.Active,
	// }

	// if err := db.Create(&Schedule).Error; err != nil {
	// 	return nil, err
	// }

	// res := map[string]interface{}{
	// 	"id":      Schedule.ID,
	// 	"label":   Schedule.Label,
	// 	"cam_url": Schedule.CamURL,
	// 	"status":  Schedule.Status,
	// 	"active":  Schedule.Active,
	// }

	// return res, nil

	return map[string]interface{}{
		"items":       "resSchedules",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// UpdatedSchedule update a Schedule
func UpdatedSchedule(db *gorm.DB, ps *ScheduleParams) (map[string]interface{}, error) {
	// var Schedule mo.Schedule

	// if err := db.Where("id = ?", ps.ID).
	// 	First(&Schedule).Error; err != nil {
	// 	return nil, err
	// }

	// Schedule.Label = ps.Label
	// Schedule.CamURL = ps.CamURL
	// Schedule.Status = ps.Status
	// Schedule.Active = ps.Active

	// if err := db.Save(&Schedule).Error; err != nil {
	// 	return nil, err
	// }

	// res := map[string]interface{}{
	// 	"id":      Schedule.ID,
	// 	"label":   Schedule.Label,
	// 	"cam_url": Schedule.CamURL,
	// 	"status":  Schedule.Status,
	// 	"active":  Schedule.Active,
	// }

	// return res, nil

	return map[string]interface{}{
		"items":       "resSchedules",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// DeletedSchedule delete a Schedule
func DeletedSchedule(db *gorm.DB, id int64) error {
	// var Schedule mo.Schedule
	// t := time.Now()

	// if err := db.Where("id = ?", id).
	// 	First(&Schedule).Error; err != nil {
	// 	return err
	// }
	// Schedule.DeletedAt = &t

	// if err := db.Delete(&Schedule).Error; err != nil {
	// 	return err
	// }

	// return nil

	return nil
}
