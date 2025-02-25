package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrUserWithEmailExists = core.NewI18NError(core.EINVALID, core.TXT_USER_WITH_EMAIL_EXISTS)
