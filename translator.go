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

func Translate(key string) string {
	return TranslateModule.Translate(translations, lang, fallbackLang, key)
}
