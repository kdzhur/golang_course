package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rest/cmd/task2/internal/data"
)

func TranslateText(q, source, target, format string) (string, error) {

	apiKey, ok := os.LookupEnv("TRANSLATE_API_KEY")
	if !ok {
		return "", fmt.Errorf("no API key provided. TRANSLATE_API_KEY env variable should be set")
	}

	apiUrl := fmt.Sprintf("https://libretranslate.com/translate?key=%s", apiKey)

	requestBody := map[string]interface{}{
		"q":      q,
		"source": source,
		"target": target,
		"format": format,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var translationResponse data.TranslationResponse
		err := json.Unmarshal(body, &translationResponse)
		if err != nil {
			return "", err
		}
		return translationResponse.TranslatedText, nil
	}

	return "", fmt.Errorf("error making API request: %s", resp.Status)
}
