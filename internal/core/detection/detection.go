package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlSeyesCam = "http://202.44.35.76:9093/image"

type DetectionParams struct {
	Uuid      string `json:"id"`
	Channel   string `json:"channel"`
	ImageData string `json:"image"`
}

func GetImageFromSeyesCam(ps *DetectionParams) (string, error) {

	resFromSCAM, err := http.Get(urlSeyesCam + "/" + ps.Uuid + "/channel" + "/" + ps.Channel) //+ "/" + ps.Uuid + "/channel" + ps.Channel
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(resFromSCAM.Body)
	if err != nil {
		return "", err
	}
	var responseObject DetectionParams
	json.Unmarshal(responseData, &responseObject)

	var jsonStr = []byte(`{"image":` + `"` + responseObject.ImageData + `"` + `}`)
	req, err := http.NewRequest("POST", "http://202.44.35.76:9094/detect", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return responseObject.ImageData, nil
}

func DetectionFromSeyesDetect() {

}
