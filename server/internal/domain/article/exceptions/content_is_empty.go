package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrContentIsEmpty = core.NewI18NError(core.EINVALID, core.TXT_ARTICLES_CONTENT_IS_EMPTY)
