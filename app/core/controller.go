package core

import (
	"encoding/json"
	"fmt"
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

func (con *Controller) SendValidationError(err error, w http.ResponseWriter) {
	// translate all error at once
	errs := err.(validator.ValidationErrors)
	trans, _ := UniversalTranslator.GetTranslator("en")

	customerrors := make([]*jsonapi.ErrorObject, len(errs))
	for index, e := range errs {
		// can translate each error one at a time.
		fmt.Println(e.Translate(trans))
		customerror := &jsonapi.ErrorObject{
			Title:  e.Field(),
			Detail: e.Translate(trans),
			Status: "422",
			Code:   VALIDATIONERROR,
		}
		customerrors[index] = customerror
	}

	con.SendCustomError(w, customerrors, http.StatusUnprocessableEntity)
}
