package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrUserNotAdmin = core.NewI18NError(core.EFORBIDDEN, core.TXT_USER_IS_NOT_ADMIN)
