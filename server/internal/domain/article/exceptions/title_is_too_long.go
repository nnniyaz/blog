package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrTitleIsTooLong = core.NewI18NError(core.EINVALID, core.TXT_ARTICLES_TITLE_IS_TOO_LONG)
