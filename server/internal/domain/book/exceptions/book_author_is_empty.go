package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrBookAuthorIsEmpty = core.NewI18NError(core.EINVALID, core.TXT_AUTHOR_OF_BOOK_IS_EMPTY)
