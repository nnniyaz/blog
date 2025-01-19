package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrUserEmailNotFound = core.NewI18NError(core.ENOTFOUND, core.TXT_USER_WITH_THIS_EMAIL_NOT_FOUND)
