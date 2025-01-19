package exceptions

import (
	"github.com/nnniyaz/blog/pkg/core"
)

var ErrUserPasswordInvalid = core.NewI18NError(core.EINVALID, core.TXT_USER_PASSWORD_INVALID)
