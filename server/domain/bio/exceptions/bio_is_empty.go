package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrBioIsEmpty = core.NewI18NError(core.EINVALID, core.TXT_BIO_IS_EMPTY)
