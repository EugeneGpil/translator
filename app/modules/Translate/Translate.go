package Translate

type Translator struct {
	translations map[string]map[string]string
	lang         string
	fallbackLang string
	key          string
	placeholders map[string]string
}

func New() Translator {
	return Translator{}
}

func (translator Translator) SetTranslations(translations map[string]map[string]string) Translator {
	translator.translations = translations

	return translator
}

func (translator Translator) SetLang(lang string) Translator {
	translator.lang = lang

	return translator
}

func (translator Translator) SetFallbackLang(fallbackLang string) Translator {
	translator.fallbackLang = fallbackLang

	return translator
}

func (translator Translator) SetKey(key string) Translator {
	translator.key = key

	return translator
}

func (translator Translator) SetPlaceholders(placeholders map[string]string) Translator {
	translator.placeholders = placeholders

	return translator
}

func (translator Translator) Run() string {
	if translation := translator.getTranslation(); translation != "" {
		return translation
	}

	return translator.getFallbackTranslation()
}

func (translator Translator) getTranslation() string {
	if langTranslations, isset := translator.translations[translator.lang]; isset {
		if keyTranslation, isset := langTranslations[translator.key]; isset {
			return keyTranslation
		}
	}

	return ""
}

func (translator Translator) getFallbackTranslation() string {
	if translation := translator.getTranslation(); translation != "" {
		return translation
	}

	return translator.key
}
