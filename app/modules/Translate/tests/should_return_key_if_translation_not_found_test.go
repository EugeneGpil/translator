package tests

import (
	"testing"

	"github.com/EugeneGpil/translator/app/modules/Translate/tests/base"
)

func Test_should_return_key_if_translation_not_found(t *testing.T) {
	key := "Some Key"

	translations := map[string]map[string]string{}

	base.AssertTranslation(base.AssertTranslationDto{
		Key:                 key,
		Translations:        translations,
		ExpectedTranslation: key,
	})
}
