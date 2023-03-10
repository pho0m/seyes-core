package core

import (
	"seyes-core/internal/helper"

	"gorm.io/gorm"
)

// ReportsFilter define Reports filter
type ReportsFilter struct {
	Page    int64    `json:"page"`
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Active  []string `json:"active"`
	OrderBy string   `json:"order_by"`
	SortBy  string   `json:"sort_by"`
}

// ReportsParams define params for create Reports
type ReportsParams struct {
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

// GetAllReportss get all Reports product
func GetAllReportss(db *gorm.DB, filter *ReportsFilter) (map[string]interface{}, error) {
	// // var resPr []ReportsParams
	// var Reports []mo.Reports
	// var resReportss []ReportsParams

	// dbx := db.Model(&mo.Reports{})
	// pg := helper.FormatWebPaginate(dbx, filter.Page)

	// if err := pg.DB.Find(&Reports).Error; err != nil {
	// 	return nil, err
	// }

	// for _, r := range Reports {
	// 	resReportss = append(resReportss, ReportsParams{
	// 		ID:     int64(r.ID),
	// 		Label:  r.Label,
	// 		CamURL: r.CamURL,
	// 		Status: r.Status,
	// 		Active: r.Active,
	// 	})
	// }

	// if len(resReportss) == 0 {
	// 	resReportss = []ReportsParams{}
	// }

	return map[string]interface{}{
		"items":       "resReportss",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// GetReports get a Reports by Reports id
func GetReports(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	// var Reports mo.Reports

	// if err := db.Where("id = ?", ps.ID).
	// 	Where("deleted_at IS NULL").
	// 	First(&Reports).Error; err != nil {
	// 	return nil, err
	// }

	return map[string]interface{}{
		"items":       "resReportss",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil

}

// CreateReports create a Reports
func CreateReports(db *gorm.DB, ps *ReportsParams) (map[string]interface{}, error) {
	// Reports := &mo.Reports{
	// 	Label:  ps.Label,
	// 	CamURL: ps.CamURL,
	// 	Status: ps.Status,
	// 	Active: ps.Active,
	// }

	// if err := db.Create(&Reports).Error; err != nil {
	// 	return nil, err
	// }

	// res := map[string]interface{}{
	// 	"id":      Reports.ID,
	// 	"label":   Reports.Label,
	// 	"cam_url": Reports.CamURL,
	// 	"status":  Reports.Status,
	// 	"active":  Reports.Active,
	// }

	// return res, nil

	return map[string]interface{}{
		"items":       "resReportss",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// UpdatedReports update a Reports
func UpdatedReports(db *gorm.DB, ps *ReportsParams) (map[string]interface{}, error) {
	// var Reports mo.Reports

	// if err := db.Where("id = ?", ps.ID).
	// 	First(&Reports).Error; err != nil {
	// 	return nil, err
	// }

	// Reports.Label = ps.Label
	// Reports.CamURL = ps.CamURL
	// Reports.Status = ps.Status
	// Reports.Active = ps.Active

	// if err := db.Save(&Reports).Error; err != nil {
	// 	return nil, err
	// }

	// res := map[string]interface{}{
	// 	"id":      Reports.ID,
	// 	"label":   Reports.Label,
	// 	"cam_url": Reports.CamURL,
	// 	"status":  Reports.Status,
	// 	"active":  Reports.Active,
	// }

	// return res, nil

	return map[string]interface{}{
		"items":       "resReportss",
		"page":        "pg.Page",
		"total_pages": "pg.TotalPages",
		"total_count": "pg.TotalCount",
	}, nil
}

// DeletedReports delete a Reports
func DeletedReports(db *gorm.DB, id int64) error {
	// var Reports mo.Reports
	// t := time.Now()

	// if err := db.Where("id = ?", id).
	// 	First(&Reports).Error; err != nil {
	// 	return err
	// }
	// Reports.DeletedAt = &t

	// if err := db.Delete(&Reports).Error; err != nil {
	// 	return err
	// }

	// return nil

	return nil
}
