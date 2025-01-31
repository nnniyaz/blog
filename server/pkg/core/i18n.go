package core

// ----------------------------------------------------------------------------
// Lang
// ----------------------------------------------------------------------------

type Lang string

const (
	KZ Lang = "KZ"
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

func (m MlString) Clean() (MlString, error) {
	if m.IsEmpty() {
		return nil, ErrEmptyMlString
	}
	mMap := map[Lang]string(m)
	cleaned := make(MlString)
	for k, v := range mMap {
		cleaned[k] = v
	}
	return cleaned, nil
}

// ----------------------------------------------------------------------------
// TxtResource
// ----------------------------------------------------------------------------

type TxtResource map[TxtKey]MlString
