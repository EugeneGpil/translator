package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_translation(t *testing.T) {
	lang := "testLang"
	key := "testKey"
	expectedTranslation := "testTranslation"

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
