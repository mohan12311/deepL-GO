package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	MARSHAL_ERROR = "failed to marshal request body: "
)

type Client struct {
	HTTPClient *http.Client
	APIKey     string
	Host       string
}

type ClientConfig struct {
	APIKey          string
	IsPremiumMember bool
}

func NewClientConfig(apiKey string) *ClientConfig {
	isPremiumMember := true
	if len(apiKey) >= 3 && strings.HasSuffix(apiKey, ":fx") {
		isPremiumMember = false
	}

	return &ClientConfig{APIKey: apiKey, IsPremiumMember: isPremiumMember}

}

func NewClient(c *ClientConfig) *Client {
	host := getHost(c.IsPremiumMember)
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		APIKey:     c.APIKey,
		Host:       host,
	}
}

func getHost(isPremiumMember bool) string {
	if isPremiumMember {
		return PREMIUM_MEMBER
	}
	return FREE_MEMBER

}

func (c *Client) Translate(tr TranslationRequest) (*TranslationResponse, error) {

	reqBody, err := json.Marshal(tr)
	if err != nil {
		return nil, fmt.Errorf(MARSHAL_ERROR, err)
	}

	req, err := http.NewRequest("POST", c.Host+"/v2/translate", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	req.Header.Add("Authorization", "DeepL-Auth-Key "+c.APIKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	// 응답 상태 코드 확인
	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("received non-200 response code: %d, body: %s", res.StatusCode, string(bodyBytes))
	}

	// 응답 본문 출력 및 디코딩
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response TranslationResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)

	}

	return &response, nil
}
