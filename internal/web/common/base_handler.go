package common

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// ErrParams define return err handler
type ErrParams struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

// Handler defines interface for web handler
type Handler interface {
	Register(r chi.Router)
}

// BaseRender provides base functionality for web handler
type BaseRender struct{}

// JSON return a JSON response
func (h *BaseRender) JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	if _, err := w.Write(d); err != nil {
		log.Println("Cannot write a response:", err.Error())
	}
}

// Error return a Error response
func (h *BaseRender) Error(w http.ResponseWriter, err interface{}, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	switch statusCode {
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
	case *ErrParams:
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
