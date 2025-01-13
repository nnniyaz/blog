package core

import "fmt"

type ErrCode string

const (
	EINTERNAL     ErrCode = "internal"
	ENOTFOUND     ErrCode = "not_found"
	EINVALID      ErrCode = "invalid"
	ECONFLICT     ErrCode = "conflict"
	EUNAUTHORIZED ErrCode = "unauthorized"
	EFORBIDDEN    ErrCode = "forbidden"
)

type I18NError struct {
	Code   ErrCode
	TxtKey TxtKey
	Args   []interface{}
}

func NewI18NError(code ErrCode, txtKey TxtKey, args ...interface{}) *I18NError {
	if code == "" {
		code = EINTERNAL
	}
	if txtKey == 0 {
		txtKey = TXT_UNKNOWN_ERROR
	}
	return &I18NError{code, txtKey, args}
}

func (i *I18NError) Error() string {
	return fmt.Sprintf("%q[%s]", fmt.Sprintf(GetTxtKeyAsString(i.TxtKey), i.Args...), string(i.Code))
}

func (i *I18NError) ErrMsg() string {
	return fmt.Sprintf(Txts[i.TxtKey].GetByLangOrEmpty(EN), i.Args...)
}

type I18NErrors struct {
	Errors []*I18NError
}

func NewI18NErrors(errs []*I18NError) *I18NErrors {
	return &I18NErrors{errs}
}

func (i *I18NErrors) Error() string {
	var err string
	for _, e := range i.Errors {
		err = fmt.Sprintf("%s %s", err, e.Error())
	}
	return err
}

func IsNotFoundErr(err error) bool {
	if err == nil {
		return false
	}
	if err, ok := err.(*I18NError); ok {
		return err.Code == ENOTFOUND
	}
	return false
}
