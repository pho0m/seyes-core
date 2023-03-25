package core

import (
	"bytes"
	"encoding/base64"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"seyes-core/internal/helper"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const urlNotify = "https://notify-api.line.me/api/notify"

// NotifyParam define data to Notify template
type NotifyParam struct {
	ID        int64          `json:"id"`
	Person    int64          `json:"person"`
	ComOn     int64          `json:"com_on"`
	UploadAt  string         `json:"upload_at"`
	Time      string         `json:"time"`
	Image     multipart.File `json:"image"`
	Accurency string         `json:"accurency"`
}

// NotifyParam define data to Notify template
type NotifyParamV2 struct {
	ID        int64  `json:"id"`
	Person    string `json:"person"`
	ComOn     string `json:"com_on"`
	UploadAt  string `json:"upload_at"`
	Time      string `json:"time"`
	Image     string `json:"image"`
	Accurency string `json:"accurency"`
}

type SendDataToNotify struct {
	ImageFile multipart.File `json:"imageFile"`
	Message   string         `json:"message"`
}

type ResponseNotify struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

func SendToLineNotify(ps *NotifyParam) (*ResponseNotify, error) {

	person := strconv.Itoa(int(ps.Person))
	comOn := strconv.Itoa(int(ps.ComOn))

	accessToken := "Bearer " + os.Getenv("NOTIFY_TOKEN") //FIXME GET FROM SETTING
	message := "Detection !" + "\n" +
		"Person : " + person + "\n" +
		"Com On : " + comOn + "\n" +
		"Upload at : " + ps.UploadAt + "\n" +
		"Time : " + ps.Time + "\n" +
		"Accurency : " + ps.Accurency + "%"

	body, contentType, err := helper.MakeMultipartBody(message, ps.Image)

	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(&body)

	req, err := http.NewRequest("POST", urlNotify, buf)

	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("authorization", accessToken)
	response, err := client.Do(req)

	if response.StatusCode == 400 {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	res := ResponseNotify{
		Message:    "data have been send !",
		StatusCode: strconv.Itoa(int(response.StatusCode)),
	}

	return &res, nil
}

func SendToLineNotifyV2(ps *NotifyParamV2) error {
	logrus.Print(ps)

	accessToken := "Bearer " + os.Getenv("NOTIFY_TOKEN") //FIXME GET FROM SETTING
	message := "Detection !" + "\n" +
		"Person : " + ps.Person + "\n" +
		"Com On : " + ps.ComOn + "\n" +
		"Upload at : " + ps.UploadAt + "\n" +
		"Time : " + ps.Time + "\n" +
		"Accurency : " + ps.Accurency

	convertBase64ToFile(ps.Image)
	f, err := os.Open("./detected.jpeg")
	if err != nil {
		return err
	}

	body, contentType, err := helper.MakeMultipartBody(message, f)

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

	// res := ResponseNotify{
	// 	Message:    "data have been send !",
	// 	StatusCode: strconv.Itoa(int(response.StatusCode)),
	// }

	return nil
}

func convertBase64ToFile(data string) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		panic("InvalidImage")
	}
	ImageType := data[11:idx]
	log.Println(ImageType)

	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)
	switch ImageType {
	case "png":
		im, err := png.Decode(r)
		if err != nil {
			panic("Bad png")
		}

		f, err := os.OpenFile("detected.png", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, im)
	case "jpeg":
		im, err := jpeg.Decode(r)
		if err != nil {
			panic("Bad jpeg")
		}

		f, err := os.OpenFile("detected.jpeg", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(r)
		if err != nil {
			panic("Bad gif")
		}

		f, err := os.OpenFile("detected.gif", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		gif.Encode(f, im, nil)
	}

	return
}
