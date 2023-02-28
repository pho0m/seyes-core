package core

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const urlNotify = "https://notify-api.line.me/api/notify"

// NotifyParam define data to Notify template
type NotifyParam struct {
	ID        int64          `json:"id"`
	Person    int64          `json:"person"`
	ComOn     int64          `json:"com_on"`
	UploadAt  string         `json:"upload_at"`
	Time      string         `json:"time"`
	Photo     multipart.File `json:"photo"`
	Accurency string         `json:"accurency"`
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

	accessToken := "Bearer " + os.Getenv("NOTIFY_TOKEN")
	message := "Detection !" + "\n" +
		"Person : " + person + "\n" +
		"Com On : " + comOn + "\n" +
		"Upload at : " + ps.UploadAt + "\n" +
		"Time : " + ps.Time + "\n" +
		"Accurency : " + ps.Accurency + "%"

	body, contentType, err := makeMultipartBody(message, ps.Photo)

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
