package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrLastNameEmpty = core.NewI18NError(core.EINVALID, core.TXT_LAST_NAME_IS_EMPTY)
