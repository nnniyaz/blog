package exceptions

import (
	"github.com/nnniyaz/blog/pkg/core"
)

var ErrUserRoleInvalid = core.NewI18NError(core.EINVALID, core.TXT_USER_ROLE_INVALID)
