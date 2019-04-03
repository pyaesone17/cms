package core

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/google/jsonapi"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	UniversalTranslator *ut.UniversalTranslator
	validate            *validator.Validate
)

type Controller struct {
}

func init() {
	en := en.New()
	UniversalTranslator = ut.New(en, en)
	trans, _ := UniversalTranslator.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func (con *Controller) SendJSON(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (con *Controller) SendServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Title:  "Internal Server Error",
		Detail: "Something Went Wrong Inside server.",
		Status: "500",
		Code:   CODESERVERERROR,
	}})
}

func (con *Controller) SendCustomError(w http.ResponseWriter, errors []*jsonapi.ErrorObject, statusCode int) {
	w.WriteHeader(statusCode)

	jsonapi.MarshalErrors(w, errors)
}

func (con *Controller) Validate(data interface{}) error {
	return validate.Struct(data)
}

// GetContent of the request inside given struct
func (con *Controller) GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		log.Printf("GetContent Decode error: %s", err.Error())
		return err
	}

	return nil
}
