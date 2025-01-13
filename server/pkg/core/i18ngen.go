// code generated by gen.go
// DO NOT EDIT.
package core

const (
	TXT_ARTICLES_TITLE_IS_EMPTY    TxtKey = 1
	TXT_ARTICLES_TITLE_IS_TOO_LONG TxtKey = 2
	TXT_ARTICLES_CONTENT_IS_EMPTY  TxtKey = 3
)

var Txts = TxtResource{
	TXT_ARTICLES_TITLE_IS_EMPTY: MlString{
		KZ: `Заголовок статьи пуст`,
		RU: `Заголовок статьи пуст`,
		EN: `Article title is empty`,
	},
	TXT_ARTICLES_TITLE_IS_TOO_LONG: MlString{
		KZ: `Заголовок статьи слишком длинный`,
		RU: `Заголовок статьи слишком длинный`,
		EN: `Article title is too long`,
	},
	TXT_ARTICLES_CONTENT_IS_EMPTY: MlString{
		KZ: `Содержимое статьи пустое`,
		RU: `Содержимое статьи пустое`,
		EN: `Article content is empty`,
	},
}

func GetTxtKeyAsString(k TxtKey) string {
	switch k {
	case TXT_ARTICLES_TITLE_IS_EMPTY:
		return "articles_title_is_empty"
	case TXT_ARTICLES_TITLE_IS_TOO_LONG:
		return "articles_title_is_too_long"
	case TXT_ARTICLES_CONTENT_IS_EMPTY:
		return "articles_content_is_empty"

	default:
		return ""
	}
}
