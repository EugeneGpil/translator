package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_translation_if_no_languages_set_at_all(t *testing.T) {
	key := "testKey"
	expectedTranslation := "testTranslation"
	defaultFallbackLang := "en"

	translations := map[string]map[string]string{
		defaultFallbackLang: {
			key: expectedTranslation,
		},
	}

	base.AssertTranslation(base.AssertTranslationDto{
		Translations:        translations,
		Key:                 key,
		ExpectedTranslation: expectedTranslation,
		T:                   t,
	})
}
