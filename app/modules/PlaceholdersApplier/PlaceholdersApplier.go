package PlaceholdersApplier

import "strings"

type PlaceholdersApplier struct {
	placeholders map[string]string
	translation string
}

func New() PlaceholdersApplier {
	return PlaceholdersApplier{}
}

func (placeholdersApplier *PlaceholdersApplier) SetPlaceholders(placeholders map[string]string) *PlaceholdersApplier {
	placeholdersApplier.placeholders = placeholders

	return placeholdersApplier
}

func (placeholdersApplier *PlaceholdersApplier) SetTranslation(translation string) *PlaceholdersApplier {
	placeholdersApplier.translation = translation

	return placeholdersApplier
}

func (placeholdersApplier PlaceholdersApplier) Run() string {
	res := placeholdersApplier.translation

	for key, placeholder := range placeholdersApplier.placeholders {
		res = strings.Replace(res, ":" + key, placeholder, -1)
	}

	return res
}
