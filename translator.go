package translator

import (
	TranslateModule "github.com/EugeneGpil/translator/app/modules/Translate"
)

var lang = "en"
var fallbackLang = "en"
var translations map[string]map[string]string

func SetLang(langToSet string) {
	lang = langToSet
}

func SetFallbackLang(langToSet string) {
	fallbackLang = langToSet
}

func SetTranslations(translationsToSet map[string]map[string]string) {
	translations = translationsToSet
}

// func Translate(key string, placeholders map[string]string) string {
func Translate(key string) string {
	return TranslateModule.
		New().
		SetTranslations(translations).
		SetLang(lang).
		SetFallbackLang(fallbackLang).
		SetKey(key).
		// SetPlaceholders(placeholders).
		Run()
}
