package exceptions

import (
	"github.com/nnniyaz/blog/pkg/core"
)

var ErrTitleEmpty = core.NewI18NError(core.EINVALID, core.TXT_ARTICLES_TITLE_IS_EMPTY)
