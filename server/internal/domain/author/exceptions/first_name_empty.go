package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrFirstNameEmpty = core.NewI18NError(core.EINVALID, core.TXT_FIRST_NAME_IS_EMPTY)
