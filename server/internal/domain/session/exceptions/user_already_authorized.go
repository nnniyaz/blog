package exceptions

import (
	"github.com/nnniyaz/blog/pkg/core"
)

var ErrUserAlreadyAuthorized = core.NewI18NError(core.EFORBIDDEN, core.TXT_USER_ALREADY_AUTHORIZED)
