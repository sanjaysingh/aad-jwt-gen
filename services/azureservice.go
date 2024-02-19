package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GenerateToken(clientId, tenantId, clientSecret, scope string) (string, error) {
	oAuthUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantId)
	body := []byte(fmt.Sprintf("client_id=%s&scope=%s&client_secret=%s&grant_type=client_credentials", clientId, url.QueryEscape(scope), clientSecret))

	req, _ := http.NewRequest("POST", oAuthUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("error parsing access token")
	}

	return token, nil
}
