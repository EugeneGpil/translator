package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_translation_by_fallback_language_if_main_language_translations_set(t *testing.T) {
	key := "someKey"
	lang := "mainLang"
	fallbackLang := "fallbackLang"
	expectedTranslation := "expectedTranslation"

	translations := map[string]map[string]string{
		lang: {},
		fallbackLang: {
			key: expectedTranslation,
		},
	}

	base.AssertTranslation(base.AssertTranslationDto{
		Translations:        translations,
		Lang:                lang,
		FallbackLang:        fallbackLang,
		Key:                 key,
		ExpectedTranslation: expectedTranslation,
		T:                   t,
	})
}
