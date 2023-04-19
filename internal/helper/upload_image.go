package helper

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

//*FIXME please refactor comment function

// ImageFileParams defines File for Upload image
type ImageFileParams struct {
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

// UploadImage uploads file to web server // FIXME
func UploadImage(ps *UploadFileParams) (*Media, error) {
	mediaName := "media/temp" //os.Getenv("STORAGE_MEDIA_NAME")
	padding, err := GeneratePadding()

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

func MakeMultipartBody(message string, image multipart.File) (body bytes.Buffer, contentType string, err error) {
	writer := multipart.NewWriter(&body)

	if err = WriteBodyInMessage(writer, message); err != nil {
		return
	}
	if err = WriteBodyInImageFile(writer, image); err != nil {
		return
	}
	writer.Close()

	contentType = writer.FormDataContentType()
	return
}

func WriteBodyInMessage(writer *multipart.Writer, message string) (err error) {
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

func WriteBodyInImageFile(writer *multipart.Writer, image multipart.File) (err error) {

	mediaName := "media/" //os.Getenv("STORAGE_MEDIA_NAME")
	padding, err := GeneratePadding()

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

func WriteImageToRespone(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
