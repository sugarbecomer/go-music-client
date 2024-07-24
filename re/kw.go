package re

import "regexp"

var (
	KWTokenReg = regexp.MustCompile(`kw_token=(.*?);`)

	KWUrlReg = regexp.MustCompile(`url=.+`)
)
