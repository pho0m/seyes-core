package helper

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

const PERPAGEWEB = 20

// const digit = 100

// // Layout time
// const defTime = "0001-01-01T00:00:00Z"
// const layoutDateTime = "2006-01-02T15:04:05Z"
// const layoutDate = "2006-01-02"
// const layoutTime = "15:04"

// Size constants
const (
	MB = 1 << 20
)

// ResponseData define response data
type ResponseData struct {
	Items      []interface{} `json:"items"`
	Page       int64         `json:"page"`
	NextPage   int64         `json:"next_page"`
	TotalPages float64       `json:"total_pages"`
	TotalCount int64         `json:"total_count"`
}

// ErrParams define return err handler
type ErrParams struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

// UrlParams define filter component
type UrlParams struct {
	Status  string `json:"status"`
	Slug    string `json:"slug"`
	Page    int64  `json:"pages"`
	ID      int64  `json:"id"`
	OrderNo string `json:"order_no"`
}

// QueryStringParams defines struct for r.Query()
type QueryStringParams struct {
	ID        int64
	Slug      string
	Page      int64
	Name      string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Date      string
	Time      string
	DateStart string
	DateEnd   string
	UpdateAt  string
	CreateAt  string
	OrderBy   string
	SortBy    string
}

type PaginateParams struct {
	DB         *gorm.DB
	Page       int64
	TotalPages float64
	TotalCount int64
}

// FormatWebPaginate format db data into pagination
func FormatWebPaginate(db *gorm.DB, page int64) PaginateParams {
	var count int64
	var totalPage float64

	if page != 0 {
		db.Count(&count)
		offset := (page - 1) * PERPAGEWEB
		db = db.Offset(int(offset)).Limit(PERPAGEWEB)
		totalPage = math.Ceil(float64(count) / float64(PERPAGEWEB))
	}

	return PaginateParams{
		DB:         db,
		Page:       page,
		TotalPages: totalPage,
		TotalCount: count,
	}
}

// ParsingUploadFileParams parsing file data multipart.file to return UploadFileParams
func ParsingUploadFileParams(file multipart.File, handler *multipart.FileHeader) *UploadFileParams {
	mediaType := handler.Header.Get("Content-Type")

	return &UploadFileParams{
		MediaType:  mediaType,
		File:       file,
		FileHeader: handler,
	}
}

// QueryUploadParams define data to query upload params template
type QueryUploadParams struct {
	ID        int64  `json:"id"`
	Person    int64  `json:"person"`
	ComOn     int64  `json:"com_on"`
	UploadAt  string `json:"upload_at"`
	Time      string `json:"time"`
	Accurency string `json:"accurency"`
}

// UploadFileParams defines parameters fro UploadFile
type UploadFileParams struct {
	MediaType  string
	File       multipart.File
	FileHeader *multipart.FileHeader
}

// ParsingQueryUpload parse query type,id and return QueryUploadParams
func ParsingQueryUpload(query interface{}) *QueryUploadParams {
	var (
		ID        = 0
		person    = 0
		conOn     = 0
		upload_at = ""
		time      = ""
		accurency = ""
	)

	switch v := query.(type) {
	case url.Values:
		p := v.Get("person")
		person, _ = strconv.Atoi(p)

		c := v.Get("com_on")
		conOn, _ = strconv.Atoi(c)

		upload_at = v.Get("upload_at")
		time = v.Get("time")
		accurency = v.Get("accurency")

		i := v.Get("id")
		ID, _ = strconv.Atoi(i)
	}

	return &QueryUploadParams{
		ID:        int64(ID),
		Person:    int64(person),
		ComOn:     int64(conOn),
		UploadAt:  upload_at,
		Time:      time,
		Accurency: accurency,
	}
}

// ParsingQueryString parse query string and return type int
func ParsingQueryString(query interface{}) *QueryStringParams {
	var (
		page      = 0
		ID        = 0
		Slug      = ""
		firstName = ""
		lastName  = ""
		name      = ""
		phone     = ""
		email     = ""
		date      = ""
		time      = ""
		dateStart = ""
		dateEnd   = ""
		createdAt = ""
		updatedAt = ""
		orderBy   = ""
		sortBy    = ""
	)

	switch v := query.(type) {
	case url.Values:
		p := v.Get("page")

		page, _ = strconv.Atoi(p)

		name = v.Get("name")
		firstName = v.Get("first_name")
		lastName = v.Get("last_name")
		phone = v.Get("phone")
		email = v.Get("email")
		date = v.Get("date")
		time = v.Get("time")
		dateStart = v.Get("date_start")
		dateEnd = v.Get("date_end")
		orderBy = v.Get("order_by")
		sortBy = v.Get("sort_by")
		createdAt = v.Get("created_at")
		updatedAt = v.Get("updated_at")

	case string:
		ID, _ = strconv.Atoi(v)
		Slug = v

	}

	return &QueryStringParams{
		Page:      int64(page),
		ID:        int64(ID),
		Slug:      Slug,
		Name:      name,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Date:      date,
		Time:      time,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		OrderBy:   orderBy,
		SortBy:    sortBy,
		CreateAt:  createdAt,
		UpdateAt:  updatedAt,
	}
}

// Padding padding digit
func Padding(v int64, length int) string {
	abs := math.Abs(float64(v))
	var padding int
	if v != 0 {
		min := math.Pow10(length - 1)

		if min-abs > 0 {
			l := math.Log10(abs)
			if l == float64(int64(l)) {
				l++
			}
			padding = length - int(math.Ceil(l))
		}
	} else {
		padding = length - 1
	}
	builder := strings.Builder{}
	if v < 0 {
		length = length + 1
	}
	builder.Grow(length * 4)
	if v < 0 {
		builder.WriteRune('-')
	}
	for i := 0; i < padding; i++ {
		builder.WriteRune('0')
	}
	builder.WriteString(strconv.FormatInt(int64(abs), 10))
	return builder.String()
}

// GeneratePadding generate padding name
func GeneratePadding() (string, error) {
	randDomLeft := rand.Intn(999)
	iLeft := int64(randDomLeft)
	padLeft := Padding(iLeft, 3)
	randDomRight := rand.Intn(999)
	iRight := int64(randDomRight)
	padRight := Padding(iRight, 3)
	pDigit := padLeft + padRight

	return pDigit, nil
}

// ReturnError return StatusForbidden, StatusInternalServerError, NotFound
func ReturnError(w http.ResponseWriter, err interface{}, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)

	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)

	case http.StatusForbidden:
		w.WriteHeader(http.StatusForbidden)

	case http.StatusUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)

	case http.StatusBadGateway:
		w.WriteHeader(http.StatusBadGateway)
	}

	switch newEr := err.(type) {

	case error:
		if newEr != nil || msg != "" {
			er := ErrParams{
				Error: newEr.Error(),
				Msg:   msg,
			}
			d, err := json.Marshal(er)

			if err != nil {
				panic(err)
			}

			if _, err := w.Write(d); err != nil {
				log.Println("Cannot write a response:", err.Error())
			}

			return
		}

	case ErrParams:
		d, err := json.Marshal(newEr)

		if err != nil {
			panic(err)
		}

		if _, err := w.Write(d); err != nil {
			log.Println("Cannot write a response:", err.Error())
		}

		return

	default:
		er := ErrParams{
			Msg: msg,
		}
		d, err := json.Marshal(er)

		if err != nil {
			panic(err)
		}

		if _, err := w.Write(d); err != nil {
			log.Println("Cannot write a response:", err.Error())
		}

		return
	}

	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println("Cannot write a response:", err.Error())
	}
}
