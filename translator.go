package translator

import "github.com/EugeneGpil/translator/app/modules/PlaceholdersApplier"

var defaultFallbackLang = "en";

type Translator struct {
	translations map[string]map[string]string
	lang         string
	fallbackLang string
	key          string
	placeholders map[string]string
}

func New() Translator {
	return Translator{
		fallbackLang: defaultFallbackLang,
	}
}

func (translator *Translator) SetTranslations(translations map[string]map[string]string) *Translator {
	translator.translations = translations

	return translator
}

func (translator *Translator) SetLang(lang string) *Translator {
	translator.lang = lang

	return translator
}

func (translator *Translator) SetFallbackLang(fallbackLang string) *Translator {
	translator.fallbackLang = fallbackLang

	return translator
}

func (translator *Translator) SetKey(key string) *Translator {
	translator.key = key

	return translator
}

func (translator *Translator) SetPlaceholders(placeholders map[string]string) *Translator {
	translator.placeholders = placeholders

	return translator
}

func (translator Translator) Run() string {
	translation := translator.getTranslation()

	placeholdersApplier := PlaceholdersApplier.New()

	return placeholdersApplier.
		SetPlaceholders(translator.placeholders).
		SetTranslation(translation).
		Run()
}

func (translator Translator) getTranslation() string {
	if translation := translator.getTranslationByLang(translator.lang); translation != "" {
		return translation
	}

	return translator.getFallbackTranslation()
}

func (translator Translator) getTranslationByLang(lang string) string {
	if langTranslations, isset := translator.translations[lang]; isset {
		if keyTranslation, isset := langTranslations[translator.key]; isset {
			return keyTranslation
		}
	}

	return ""
}

func (translator Translator) getFallbackTranslation() string {
	if translation := translator.getTranslationByLang(translator.fallbackLang); translation != "" {
		return translation
	}

	return translator.key
}
