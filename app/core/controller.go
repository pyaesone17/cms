package core

import (
	"net/http"

	"github.com/google/jsonapi"
)

type Controller struct {
}

func (con *Controller) SendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := jsonapi.MarshalPayload(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (con *Controller) SendServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Title:  "Internal Server Error",
		Detail: "Something Went Wrong Inside server.",
		Status: "500",
		Code:   CODESERVERERROR,
	}})
}

func (con *Controller) SendCustomError(w http.ResponseWriter, errors []*jsonapi.ErrorObject) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	jsonapi.MarshalErrors(w, errors)
}
