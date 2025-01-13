package core

// ----------------------------------------------------------------------------
// Lang
// ----------------------------------------------------------------------------

type Lang string

const (
	RU Lang = "RU"
	EN Lang = "EN"
)

// ----------------------------------------------------------------------------
// TxtKey
// ----------------------------------------------------------------------------

type TxtKey int

// ----------------------------------------------------------------------------
// MlString
// ----------------------------------------------------------------------------

var ErrEmptyMlString = NewI18NError(EINVALID, TXT_WRONG_MLSTRING_FORMAT)

type MlString map[Lang]string

func (m MlString) IsEmpty() bool {
	if m == nil {
		return true
	}
	mMap := map[Lang]string(m)
	for _, v := range mMap {
		if v != "" {
			return false
		}
	}
	return true
}

func (m MlString) GetByLangOrEmpty(lang Lang) string {
	if m == nil {
		return ""
	}
	if v, ok := m[lang]; ok {
		return v
	}
	return ""
}

// ----------------------------------------------------------------------------
// TxtResource
// ----------------------------------------------------------------------------

type TxtResource map[TxtKey]MlString
