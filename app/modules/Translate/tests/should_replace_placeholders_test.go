package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_replace_placeholders(t *testing.T) {
	key := "key"
	lang := "lang"
	translationWithPlaceholder := "Translation :placeholder. Translation."
	translations := map[string]map[string]string {
		lang: {
			key: translationWithPlaceholder,
		},
	}
	placeholders := map[string]string {
		"placeholder": "replaced",
	}
	expectedTranslation := "Translation replaced. Translation."

	base.AssertTranslation(base.AssertTranslationDto{
		Key: key,
		Lang: lang,
		Translations: translations,
		T: t,
		Placeholders: placeholders,
		ExpectedTranslation: expectedTranslation,
	})
}
