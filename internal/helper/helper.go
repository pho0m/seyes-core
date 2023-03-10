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
	Status     string `json:"status"`
	Slug       string `json:"slug"`
	ShopID     int64  `json:"shop_id"`
	BranchID   int64  `json:"branch_id"`
	CatID      int64  `json:"cat_id"`
	AttrID     int64  `json:"attr_id"`
	Page       int64  `json:"pages"`
	ID         int64  `json:"id"`
	CusID      int64  `json:"cus_id"`
	Keyword    string `json:"keyword"`
	OrderNo    string `json:"order_no"`
	PrintType  string `json:"print_type"`
	Channel    string `json:"channel"`
	Level      int64  `json:"level"`
	MchOrderNo string `json:"mch_order_no"`
	Type       string `json:"type"`
	OrderRef   string `json:"order_ref"`
}

// QueryStringParams defines struct for r.Query()
type QueryStringParams struct {
	ID               int64
	Slug             string
	Page             int64
	ShopID           int64
	BranchID         int64
	SKU              string
	OrderNo          string
	Name             string
	FirstName        string
	LastName         string
	Phone            string
	Email            string
	Date             string
	Time             string
	PreorderDatetime string
	DateStart        string
	DateEnd          string
	UpdateAt         string
	CreateAt         string
	OrderBy          string
	SortBy           string
	BranchIDs        []int64
	Categories       []int64
	Attributes       []int64
	AttributeItems   []int64
	Status           []string
	Channel          []string
	ShippingMethod   []string
	ShippingStatus   []string
	PaymentStatus    []string
	PaymentMethod    []string
	Type             []string
	Actives          []string
	SellOnPreorder   []string
	UserID           []int64
	Roles            []string
	UserIDs          []int64
	PreorderDate     string
	PreorderTime     string
	OrderRef         string
	Referral         string
	ReconcileIDs     []int64
	IncludeIvt       string
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
		page            = 0
		sID             = 0
		bID             = 0
		ID              = 0
		Slug            = ""
		orderNo         = ""
		users           []string
		usrsIDs         []int64
		status          = ""
		st              []string
		channel         = ""
		ch              []string
		shippingMethod  = ""
		sm              []string
		shippingStatus  = ""
		ss              []string
		paymentStatus   = ""
		pys             []string
		paymentMethod   = ""
		pym             []string
		typo            = ""
		typs            []string
		category        = ""
		categories      []int64
		active          = ""
		actives         []string
		sellOnPreorder  = ""
		sellOnPreorders []string
		branchName      = ""
		branchesName    []string
		branchIDs       []int64
		role            = ""
		roles           []string
		attribute       = ""
		attributes      []int64
		attributeItem   = ""
		attributeItems  []int64
		firstName       = ""
		lastName        = ""
		name            = ""
		phone           = ""
		email           = ""
		date            = ""
		time            = ""
		dateStart       = ""
		dateEnd         = ""
		sku             = ""
		createdAt       = ""
		updatedAt       = ""
		orderBy         = ""
		sortBy          = ""
		preDt           = ""
		preDate         = ""
		preTime         = ""
		orderRef        = ""
		referral        = ""
		reconcile       = ""
		includeIvt      = ""
		reconciles      []string
		reconcileIDs    []int64
	)

	switch v := query.(type) {
	case url.Values:
		p := v.Get("page")
		s := v.Get("shop_id")
		b := v.Get("branch_id")

		page, _ = strconv.Atoi(p)
		sID, _ = strconv.Atoi(s)
		bID, _ = strconv.Atoi(b)

		orderNo = v.Get("order_no")

		status = v.Get("status")
		if status != "" {
			st = strings.Split(status, ",")
		}

		channel = v.Get("channel")
		if channel != "" {
			ch = strings.Split(channel, ",")
		}

		shippingMethod = v.Get("shipping_method")
		if shippingMethod != "" {
			sm = strings.Split(shippingMethod, ",")
		}

		shippingStatus = v.Get("shipping_status")
		if shippingStatus != "" {
			ss = strings.Split(shippingStatus, ",")
		}

		paymentStatus = v.Get("payment_status")
		if paymentStatus != "" {
			pys = strings.Split(paymentStatus, ",")
		}

		paymentMethod = v.Get("payment_method")
		if paymentMethod != "" {
			pym = strings.Split(paymentMethod, ",")
		}

		category = v.Get("category")
		if category != "" {
			cts := strings.Split(category, ",")

			for _, i := range cts {
				ctID, _ := strconv.Atoi(i)
				categories = append(categories, int64(ctID))
			}

		}

		attribute = v.Get("attribute")
		if attribute != "" {
			ats := strings.Split(attribute, ",")

			for _, i := range ats {
				atID, _ := strconv.Atoi(i)
				attributes = append(attributes, int64(atID))
			}
		}

		attributeItem = v.Get("attribute_item")
		if attributeItem != "" {
			ats := strings.Split(attributeItem, ",")

			for _, i := range ats {
				atID, _ := strconv.Atoi(i)
				attributeItems = append(attributeItems, int64(atID))
			}
		}

		active = v.Get("active")
		if active != "" {
			actives = strings.Split(active, ",")
		}

		sellOnPreorder = v.Get("sell_on_preorder")
		if sellOnPreorder != "" {
			sellOnPreorders = strings.Split(sellOnPreorder, ",")
		}

		branchName = v.Get("branch_name")
		if branchName != "" {
			branchesName = strings.Split(branchName, ",")

			for _, i := range branchesName {
				bID, _ := strconv.Atoi(i)
				branchIDs = append(branchIDs, int64(bID))
			}
		}

		reconcile = v.Get("reconcile")
		if reconcile != "" {
			reconciles = strings.Split(reconcile, ",")

			for _, i := range reconciles {
				recID, _ := strconv.Atoi(i)
				reconcileIDs = append(reconcileIDs, int64(recID))
			}
		}

		role = v.Get("role")
		if role != "" {
			roles = strings.Split(role, ",")
		}

		typo = v.Get("type")
		if typo != "" {
			typs = strings.Split(typo, ",")
		}

		u := v.Get("user")
		if u != "" {
			users = strings.Split(u, ",")

			for _, i := range users {
				uID, _ := strconv.Atoi(i)
				usrsIDs = append(usrsIDs, int64(uID))
			}
		}

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
		sku = v.Get("sku")
		preDt = v.Get("preorder_datetime")
		preDate = v.Get("preorder_date")
		preTime = v.Get("preorder_time")
		orderRef = v.Get("order_ref")
		referral = v.Get("referral")
		includeIvt = v.Get("include_ivt")

	case string:
		ID, _ = strconv.Atoi(v)
		Slug = v

	}

	return &QueryStringParams{
		Page:             int64(page),
		ShopID:           int64(sID),
		BranchID:         int64(bID),
		ID:               int64(ID),
		Slug:             Slug,
		OrderNo:          orderNo,
		Status:           st,
		Channel:          ch,
		Name:             name,
		FirstName:        firstName,
		LastName:         lastName,
		Phone:            phone,
		Email:            email,
		ShippingMethod:   sm,
		ShippingStatus:   ss,
		PaymentStatus:    pys,
		PaymentMethod:    pym,
		Type:             typs,
		Date:             date,
		Time:             time,
		DateStart:        dateStart,
		DateEnd:          dateEnd,
		OrderBy:          orderBy,
		SortBy:           sortBy,
		Categories:       categories,
		Attributes:       attributes,
		AttributeItems:   attributeItems,
		Actives:          actives,
		SellOnPreorder:   sellOnPreorders,
		BranchIDs:        branchIDs,
		CreateAt:         createdAt,
		UpdateAt:         updatedAt,
		SKU:              sku,
		Roles:            roles,
		PreorderDatetime: preDt,
		UserIDs:          usrsIDs,
		PreorderDate:     preDate,
		PreorderTime:     preTime,
		OrderRef:         orderRef,
		Referral:         referral,
		ReconcileIDs:     reconcileIDs,
		IncludeIvt:       includeIvt,
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
