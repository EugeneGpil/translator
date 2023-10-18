package should_return_translation

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

var lang = "testLang"
var key = "testKey"
var expectedTranslation = "testTranslation"

func Test_should_return_translation(t *testing.T) {
	base.AssertTranslation(base.AssertTranslationDto{
		Translations: map[string]map[string]string{
			lang: {
				key: expectedTranslation,
			},
		},
		Lang:                lang,
		FallbackLang:        "SomeLang",
		Key:                 key,
		ExpectedTranslation: expectedTranslation,
		T:                   t,
	})
}
