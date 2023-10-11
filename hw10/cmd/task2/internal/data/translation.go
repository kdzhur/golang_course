package data

type TranslationResponse struct {
	DetectedLanguage `json:"detectedLanguage"`
	TranslatedText   string `json:"translatedText"`
}

type DetectedLanguage struct {
	Confidence int    `json:"confidence"`
	Language   string `json:"language"`
}
