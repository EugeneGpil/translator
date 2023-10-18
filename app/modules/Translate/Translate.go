package Translate

func Translate(
	translations map[string]map[string]string,
	lang string,
	fallbackLang string,
	key string,
) string {
	langTranslations, isset := translations[lang]
	if isset {
		keyTranslation, isset := langTranslations[key]
		if isset {
			return keyTranslation
		}

		return getFallbackTranslation(translations, fallbackLang, key)
	}

	return getFallbackTranslation(translations, fallbackLang, key)
}

func getFallbackTranslation(
	translations map[string]map[string]string,
	fallbackLang string,
	key string,
) string {
	fallbackTranslations, isset := translations[fallbackLang]
	if isset {
		keyTranslation, isset := fallbackTranslations[key]
		if isset {
			return keyTranslation
		}
	}

	return key
}
