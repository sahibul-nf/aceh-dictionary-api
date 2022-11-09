package dictionary

type DictionaryFormatter struct {
	ID        int    `json:"id"`
	Aceh      string `json:"aceh"`
	Indonesia string `json:"indonesia"`
	English   string `json:"english"`
}

func FormatDictionary(dictionary Dictionary) DictionaryFormatter {
	formatter := DictionaryFormatter{}
	formatter.ID = dictionary.ID
	formatter.Aceh = dictionary.Aceh
	formatter.Indonesia = dictionary.Indonesia
	formatter.English = dictionary.English

	return formatter
}

func FormatDictionaries(dictionaries []Dictionary) []DictionaryFormatter {
	if len(dictionaries) == 0 {
		return []DictionaryFormatter{}
	}

	var formatter []DictionaryFormatter
	for _, dictionary := range dictionaries {
		formatter = append(formatter, FormatDictionary(dictionary))
	}

	return formatter
}
