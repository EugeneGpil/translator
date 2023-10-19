package Translate

func Translate(
	translations map[string]map[string]string,
	lang string,
	fallbackLang string,
	key string,
) string {
	if translation := getTranslation(translations, lang, key); translation != "" {
		return translation
	}

	return getFallbackTranslation(translations, fallbackLang, key)
}

func getTranslation(
	translations map[string]map[string]string,
	lang string,
	key string,
) string {
	if langTranslations, isset := translations[lang]; isset {
		if keyTranslation, isset := langTranslations[key]; isset {
			return keyTranslation
		}
	}

	return "";
}

func getFallbackTranslation(
	translations map[string]map[string]string,
	fallbackLang string,
	key string,
) string {
	if translation := getTranslation(translations, fallbackLang, key); translation != "" {
		return translation
	}

	return key
}
