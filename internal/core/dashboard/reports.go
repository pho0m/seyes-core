package core

import (
	"fmt"
	"seyes-core/internal/helper"
	mo "seyes-core/internal/model/room"
	"strconv"
	"strings"
	"time"

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

type DetectionParams struct {
	Uuid        string `json:"id"`
	Channel     string `json:"channel"`
	ImageData   string `json:"image"`
	Accurency   string `json:"accuracy"`
	ConOnCount  string `json:"com_on_count"`
	Date        string `json:"date"`
	PersonCount string `json:"person_count"`
	Status      string `json:"status_detec"`
	Time        string `json:"time"`
}

// ReportsParams define params for create Reports
type ReportsParams struct {
	ID         int64  `json:"id"`
	PersonCont int64  `json:"person_count"`
	ComOnCount int64  `json:"com_on_count"`
	Accurency  string `json:"accurency"`
	RoomLabel  string `json:"room_label"`
	ReportTime string `json:"report_time"`
	ReportDate string `json:"report_date"`
	Image      string `json:"image"`
	Status     string `json:"status"`
	Lamp1      string `json:"lamp_1_status"`
	Lamp2      string `json:"lamp_2_status"`
	Lamp3      string `json:"lamp_3_status"`
	Lamp4      string `json:"lamp_4_status"`
	Lamp5      string `json:"lamp_5_status"`
	Lamp6      string `json:"lamp_6_status"`
	Door       string `json:"door_status"`
	Air        string `json:"air_status"`
}

// GetAllReports get all Reports product
func GetAllReports(db *gorm.DB, filter *ReportsFilter) (map[string]interface{}, error) {
	// var resPr []ReportsParams
	var reports []mo.Report
	var resReportss []ReportsParams

	dbx := db.Model(&mo.Report{})
	pg := helper.FormatWebPaginate(dbx, filter.Page)

	if err := pg.DB.Order("id desc").Find(&reports).Error; err != nil {
		return nil, err
	}

	for _, r := range reports {
		resReportss = append(resReportss, ReportsParams{
			ID:         int64(r.ID),
			PersonCont: r.PersonCont,
			ComOnCount: r.ComOnCount,
			Accurency:  r.Accurency,
			RoomLabel:  r.RoomLabel,
			ReportTime: r.ReportTime,
			ReportDate: r.ReportDate,
			Status:     r.Status,
			Lamp1:      r.Lamp1,
			Lamp2:      r.Lamp2,
			Lamp3:      r.Lamp3,
			Lamp4:      r.Lamp4,
			Lamp5:      r.Lamp5,
			Lamp6:      r.Lamp6,
			Door:       r.Door,
			Air:        r.Air,
		})
	}

	if len(resReportss) == 0 {
		resReportss = []ReportsParams{}
	}

	return map[string]interface{}{
		"items":       resReportss,
		"page":        pg.Page,
		"total_pages": pg.TotalPages,
		"total_count": pg.TotalCount,
	}, nil
}

// AnalyticsReports get all Reports product
func AnalyticsReports(db *gorm.DB, filter *ReportsFilter) (map[string]interface{}, error) {
	// var resPr []ReportsParams
	var reports []mo.Report

	dbx := db.Model(&mo.Report{})
	pg := helper.FormatWebPaginate(dbx, filter.Page)

	if err := pg.DB.Order("id desc").Find(&reports).Error; err != nil {
		return nil, err
	}

	var sumPerson []int64
	var sumComon []int64
	var sumAccurency []string
	var arr []string
	for _, r := range reports {
		sumPerson = append(sumPerson, r.PersonCont)
		sumComon = append(sumComon, r.ComOnCount)

		if r.Accurency == "0%" || r.Accurency == "" {
			continue
		} else {
			sumAccurency = strings.Split(r.Accurency, "%")
			arr = append(arr, sumAccurency...)
		}
	}

	var removeEmpty []string
	for _, str := range arr {
		if str != "" {
			removeEmpty = append(removeEmpty, str)
		}
	}

	var idx = 0
	var accFloat []float64

	for _, e := range removeEmpty {
		if s, err := strconv.ParseFloat(e, 64); err == nil {
			fmt.Println(s)
			accFloat = append(accFloat, s)
			idx++
		}
	}

	resPerson := helper.SumArr(sumPerson)
	resComOn := helper.SumArr(sumComon)
	sumAcc := helper.SumArrFloat(accFloat)

	resAcc := (sumAcc / float64(idx))

	return map[string]interface{}{
		"comon_count":  resComOn,
		"person_count": resPerson,
		"accurency":    resAcc,
	}, nil
}

// GetReport get a Reports by Reports id
func GetReport(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	var report mo.Report

	if err := db.Where("id = ?", ps.ID).
		Where("deleted_at IS NULL").
		First(&report).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":            report.ID,
		"update_at":     report.UpdatedAt,
		"person_count":  report.PersonCont,
		"com_on_count":  report.ComOnCount,
		"accurency":     report.Accurency,
		"room_label":    report.RoomLabel,
		"report_time":   report.ReportTime,
		"report_date":   report.ReportDate,
		"image":         report.Image,
		"status":        report.Status,
		"lamp_1_status": report.Lamp1,
		"lamp_2_status": report.Lamp2,
		"lamp_3_status": report.Lamp3,
		"lamp_4_status": report.Lamp4,
		"lamp_5_status": report.Lamp5,
		"lamp_6_status": report.Lamp6,
		"door_status":   report.Door,
		"air_status":    report.Air,
	}, nil

}

// CreateReport create a Reports
func CreateReport(db *gorm.DB, ps *ReportsParams) (map[string]interface{}, error) {

	report := &mo.Report{
		PersonCont: ps.PersonCont,
		ComOnCount: ps.ComOnCount,
		Accurency:  ps.Accurency,
		RoomLabel:  ps.RoomLabel,
		ReportTime: ps.ReportTime,
		ReportDate: ps.ReportDate,
		Image:      ps.Image,
		Status:     ps.Status,
		Lamp1:      ps.Lamp1,
		Lamp2:      ps.Lamp2,
		Lamp3:      ps.Lamp3,
		Lamp4:      ps.Lamp4,
		Lamp5:      ps.Lamp5,
		Lamp6:      ps.Lamp6,
		Door:       ps.Door,
		Air:        ps.Air,
	}

	if err := db.Create(&report).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":            report.ID,
		"update_at":     report.UpdatedAt,
		"person_count":  report.PersonCont,
		"com_on_count":  report.ComOnCount,
		"accurency":     report.Accurency,
		"room_label":    report.RoomLabel,
		"report_time":   report.ReportTime,
		"report_date":   report.ReportDate,
		"image":         report.Image,
		"status":        report.Status,
		"lamp_1_status": report.Lamp1,
		"lamp_2_status": report.Lamp2,
		"lamp_3_status": report.Lamp3,
		"lamp_4_status": report.Lamp4,
		"lamp_5_status": report.Lamp5,
		"lamp_6_status": report.Lamp6,
		"door_status":   report.Door,
		"air_status":    report.Air,
	}

	return res, nil

}

// UpdatedReport update a Reports
func UpdatedReport(db *gorm.DB, ps *ReportsParams) (map[string]interface{}, error) {
	var report mo.Report

	if err := db.Where("id = ?", ps.ID).
		First(&report).Error; err != nil {
		return nil, err
	}

	report.PersonCont = ps.PersonCont
	report.ComOnCount = ps.ComOnCount
	report.Accurency = ps.Accurency
	report.RoomLabel = ps.RoomLabel
	report.ReportTime = ps.ReportTime
	report.ReportDate = ps.ReportDate
	report.Image = ps.Image
	report.Status = ps.Status
	report.Lamp1 = ps.Lamp1
	report.Lamp2 = ps.Lamp2
	report.Lamp3 = ps.Lamp3
	report.Lamp4 = ps.Lamp4
	report.Lamp5 = ps.Lamp5
	report.Lamp6 = ps.Lamp6
	report.Door = ps.Door
	report.Air = ps.Air

	if err := db.Save(&report).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":            report.ID,
		"update_at":     report.UpdatedAt,
		"person_count":  report.PersonCont,
		"com_on_count":  report.ComOnCount,
		"accurency":     report.Accurency,
		"room_label":    report.RoomLabel,
		"report_time":   report.ReportTime,
		"report_date":   report.ReportDate,
		"image":         report.Image,
		"status":        report.Status,
		"lamp_1_status": report.Lamp1,
		"lamp_2_status": report.Lamp2,
		"lamp_3_status": report.Lamp3,
		"lamp_4_status": report.Lamp4,
		"lamp_5_status": report.Lamp5,
		"lamp_6_status": report.Lamp6,
		"door_status":   report.Door,
		"air_status":    report.Air,
	}

	return res, nil

}

// DeletedReport delete a Reports
func DeletedReport(db *gorm.DB, id int64) error {
	var Reports mo.Report
	t := time.Now()

	if err := db.Where("id = ?", id).
		First(&Reports).Error; err != nil {
		return err
	}
	Reports.DeletedAt = &t

	if err := db.Save(&Reports).Error; err != nil {
		return err
	}

	return nil
}
