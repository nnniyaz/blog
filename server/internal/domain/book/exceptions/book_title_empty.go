package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrBookTitleEmpty = core.NewI18NError(core.EINVALID, core.TXT_TITLE_OF_BOOK_IS_EMPTY)
