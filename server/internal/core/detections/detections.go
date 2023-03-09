package core

import (
	"io/ioutil"
)

func ReadTensorflowModel() (string, error) {
	data, err := ioutil.ReadFile("internal/service/tensorflow_model/model.json")

	if err != nil {
		return "", err
	}

	return string(data), nil
}
