package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrUserAgentInvalid = core.NewI18NError(core.EINVALID, core.TXT_USER_AGENT_INVALID)
