package strings

import (
	lisette "github.com/ivov/lisette/prelude"
	"strings"
)

func After(text, last string) string {
	ret_1 := strings.Index(text, last)
	option_2 := lisette.OptionFromCommaOk[int](ret_1, ret_1 != -1)
	if option_2.Tag == lisette.OptionSome {
		return lisette.SubstringFrom(text, option_2.SomeVal+len(last))
	}
	return text
}

func Before(text, last string) string {
	ret_1 := strings.Index(text, last)
	option_2 := lisette.OptionFromCommaOk[int](ret_1, ret_1 != -1)
	if option_2.Tag == lisette.OptionSome {
		return lisette.SubstringTo(text, option_2.SomeVal)
	}
	return text
}

func BeforeLast(text, last string) string {
	ret_1 := strings.LastIndex(text, last)
	option_2 := lisette.OptionFromCommaOk[int](ret_1, ret_1 != -1)
	if option_2.Tag == lisette.OptionSome {
		return lisette.SubstringTo(text, option_2.SomeVal)
	}
	return text
}

func AfterLast(text, last string) string {
	ret_1 := strings.LastIndex(text, last)
	option_2 := lisette.OptionFromCommaOk[int](ret_1, ret_1 != -1)
	if option_2.Tag == lisette.OptionSome {
		return lisette.SubstringFrom(text, option_2.SomeVal+len(last))
	}
	return text
}
