package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_translation_if_no_languages_set_at_all(t *testing.T) {
	key := "testKey"
	expectedTranslation := "testTranslation"

	base.AssertTranslation(base.AssertTranslationDto{
		Translations: map[string]map[string]string{
			"en": {
				key: expectedTranslation,
			},
		},
		Key:                 key,
		ExpectedTranslation: expectedTranslation,
		T:                   t,
	})
}
