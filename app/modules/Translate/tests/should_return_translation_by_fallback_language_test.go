package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_translation_by_fallback_language(t *testing.T) {
	key := "testKey"
	fallbackLanguage := "fallbackLanguage"
	expectedTranslation := "testTranslation"

	translations := map[string]map[string]string{
		fallbackLanguage: {
			key: expectedTranslation,
		},
	}

	base.AssertTranslation(base.AssertTranslationDto{
		Translations:        translations,
		Key:                 key,
		FallbackLang:        fallbackLanguage,
		ExpectedTranslation: expectedTranslation,
		T:                   t,
	})
}
