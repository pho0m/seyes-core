package core

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"seyes-core/internal/helper"
	"strings"

	"github.com/gofrs/uuid"
)

//APIS UPLOAD PHOTO V2
// PhotoFileParams defines File for Upload Photo
type PhotoFileParams struct {
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
