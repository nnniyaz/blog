package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrLastNameIsEmpty = core.NewI18NError(core.EINVALID, core.TXT_LAST_NAME_IS_EMPTY)
