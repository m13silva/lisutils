package strings

import (
	lisette "github.com/ivov/lisette/prelude"
	"strings"
)

func After(text, last string) string {
	ret_2 := strings.Index(text, last)
	option_3 := lisette.OptionFromCommaOk[int](ret_2, ret_2 != -1)
	if option_3.Tag == lisette.OptionSome {
		return lisette.SubstringFrom(text, option_3.SomeVal+len(last))
	}
	return text
}

func Before(text, last string) string {
	ret_2 := strings.Index(text, last)
	option_3 := lisette.OptionFromCommaOk[int](ret_2, ret_2 != -1)
	if option_3.Tag == lisette.OptionSome {
		return lisette.SubstringTo(text, option_3.SomeVal)
	}
	return text
}

func BeforeLast(text, last string) string {
	ret_2 := strings.LastIndex(text, last)
	option_3 := lisette.OptionFromCommaOk[int](ret_2, ret_2 != -1)
	if option_3.Tag == lisette.OptionSome {
		return lisette.SubstringTo(text, option_3.SomeVal)
	}
	return text
}

func AfterLast(text, last string) string {
	ret_2 := strings.LastIndex(text, last)
	option_3 := lisette.OptionFromCommaOk[int](ret_2, ret_2 != -1)
	if option_3.Tag == lisette.OptionSome {
		return lisette.SubstringFrom(text, option_3.SomeVal+len(last))
	}
	return text
}
