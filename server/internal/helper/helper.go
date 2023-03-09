package helper

import (
	"math"
	"math/rand"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
)

const digit = 100

// Layout time
const defTime = "0001-01-01T00:00:00Z"
const layoutDateTime = "2006-01-02T15:04:05Z"
const layoutDate = "2006-01-02"
const layoutTime = "15:04"

// Size constants
const (
	MB = 1 << 20
)

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

// ParsingUploadFileParams parsing file data multipart.file to return UploadFileParams
func ParsingUploadFileParams(file multipart.File, handler *multipart.FileHeader) *UploadFileParams {
	mediaType := handler.Header.Get("Content-Type")

	return &UploadFileParams{
		MediaType:  mediaType,
		File:       file,
		FileHeader: handler,
	}
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
