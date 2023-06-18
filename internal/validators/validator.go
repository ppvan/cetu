package validators

import (
	"net/url"
	"strings"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func AddError(v *Validator, field, message string) {

	if v == nil {
		v = New()
	}

	if _, ok := v.Errors[field]; !ok {
		v.Errors[field] = message
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(rule bool, field string, message string) {
	if !rule {
		AddError(v, field, message)
	}
}

func (v *Validator) IsBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}

func (v *Validator) IsURL(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)

	return err == nil
}
