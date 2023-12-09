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
	Placeholders        map[string]string
	ExpectedTranslation string
	T                   *testing.T
}

func AssertTranslation(dto AssertTranslationDto) {
	tester.SetTester(dto.T)

	translator := translator.New()

	if dto.Lang != "" {
		translator.SetLang(dto.Lang)
	}

	if dto.FallbackLang != "" {
		translator.SetFallbackLang(dto.FallbackLang)
	}

	translation := translator.
		SetTranslations(dto.Translations).
		SetKey(dto.Key).
		SetPlaceholders(dto.Placeholders).
		Run()

	tester.AssertSame(dto.ExpectedTranslation, translation)
}
