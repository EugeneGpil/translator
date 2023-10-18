package base

import (
	"testing"

	"github.com/EugeneGpil/tester"
	"github.com/EugeneGpil/translator"
)

type AssertTranslationDto struct {
	Translations        map[string]map[string]string
	Lang                string
	FallbackLang        string
	Key                 string
	ExpectedTranslation string
	T                   *testing.T
}

func AssertTranslation(dto AssertTranslationDto) {
	tester.SetTester(dto.T)

	translator.SetLang(dto.Lang)
	translator.SetFallbackLang(dto.FallbackLang)
	translator.SetTranslations(dto.Translations)

	translation := translator.Translate(dto.Key)

	tester.AssertSame(dto.ExpectedTranslation, translation)
}
