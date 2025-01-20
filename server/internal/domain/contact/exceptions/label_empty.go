package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrLabelEmpty = core.NewI18NError(core.EINVALID, core.TXT_NAME_OF_CONTACT_OR_SOCIAL_IS_EMPTY)
