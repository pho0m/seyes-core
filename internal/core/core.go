package core

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"seyes-core/internal/helper"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

const urlNotify = "https://notify-api.line.me/api/notify"

// NotifyParam define data to Notify template
type NotifyParam struct {
	ID       int64          `json:"id"`
	Person   int64          `json:"person"`
	ComOn    int64          `json:"com_on"`
	UploadAt string         `json:"upload_at"`
	Time     string         `json:"time"`
	Photo    multipart.File `json:"photo"`
}

type SendDataToNotify struct {
	ImageFile multipart.File `json:"imageFile"`
	Message   string         `json:"message"`
}

//APIS UPLOAD PHOTO V2
// PhotoFileParamsV2 defines File for Upload Photo
type PhotoFileParamsV2 struct {
	Bucket     string
	PublicRead bool
	MediaType  string
	FileByte   []byte
	Size       int64
}

// Media define data to Media
type Media struct {
	Key      string `json:"key"`
	MimeType string `json:"mime_type"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
}

// UploadPhoto uploads file to cloud storage // FIXME
func UploadPhoto(ps *helper.UploadFileParams) (*Media, error) {
	mediaName := "media_dev/temp" //os.Getenv("STORAGE_MEDIA_NAME")
	padding, err := helper.GeneratePadding()

	if err != nil {
		return nil, err
	}

	filename := uuid.Must(uuid.NewV4()).String() + path.Ext(ps.FileHeader.Filename)
	buffer := make([]byte, ps.FileHeader.Size)
	ps.File.Read(buffer)
	key := fmt.Sprintf("%s/%s/%s", mediaName, padding, filename)

	size := ps.FileHeader.Size
	mime := http.DetectContentType(buffer)

	m := &Media{
		Key:      key,
		MimeType: mime,
		Size:     size,
		Type:     ps.MediaType,
	}

	return m, nil
}

func SendToLineNotify(ps *NotifyParam) error {
	// id := strconv.Itoa(int(ps.ID))
	person := strconv.Itoa(int(ps.Person))
	comOn := strconv.Itoa(int(ps.ComOn))

	accessToken := "Bearer " + os.Getenv("NOTIFY_TOKEN")
	message := "Detection !" + "\n" +
		"Person : " + person + "\n" +
		"Com On : " + comOn + "\n" +
		"Upload at : " + ps.UploadAt + "\n" +
		"Time : " + ps.Time

	body, contentType, err := makeMultipartBody(message, ps.Photo)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(&body)

	req, err := http.NewRequest("POST", urlNotify, buf)

	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("authorization", accessToken)
	response, err := client.Do(req)

	if response.StatusCode == 400 {
		return err
	}

	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func makeMultipartBody(message string, image multipart.File) (body bytes.Buffer, contentType string, err error) {
	writer := multipart.NewWriter(&body)

	if err = writeBodyInMessage(writer, message); err != nil {
		return
	}
	if err = writeBodyInImageFile(writer, image); err != nil {
		return
	}
	writer.Close()

	contentType = writer.FormDataContentType()
	return
}

func writeBodyInMessage(writer *multipart.Writer, message string) (err error) {
	var messageWriter io.Writer
	messageWriter, err = writer.CreateFormField("message")
	if err != nil {
		return
	}
	if _, err = io.Copy(messageWriter, strings.NewReader(message)); err != nil {
		return
	}
	return
}

func writeBodyInImageFile(writer *multipart.Writer, image multipart.File) (err error) {

	mediaName := "media_dev/" //os.Getenv("STORAGE_MEDIA_NAME")
	padding, err := helper.GeneratePadding()

	if err != nil {
		return err
	}

	filename := uuid.Must(uuid.NewV4()).String()

	keyName := fmt.Sprintf("%s/%s/%s", mediaName, padding, filename)

	var imageWriter io.Writer
	imageWriter, err = writer.CreateFormFile("imageFile", keyName)
	if err != nil {
		return
	}
	if _, err = io.Copy(imageWriter, image); err != nil {
		return
	}
	image.Close()
	return
}
