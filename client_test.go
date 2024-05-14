package deepl

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
	"time"
)

func TestTranslate(t *testing.T) {
	// httpmock 활성화
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// 모의 응답 설정
	httpmock.RegisterResponder("POST", "https://api.deepl.com/v2/translate",
		httpmock.NewStringResponder(200, `{"translations":[{"detected_source_language":"EN","text":"Hallo Welt"}]}`))

	client := &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		APIKey:     "test-api-key",
		Host:       "https://api.deepl.com",
	}

	translationRequest := TranslationRequest{
		Text:       []string{"Hello world"},
		TargetLang: "DE",
	}

	translatedText, err := client.Translate(translationRequest)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := "Hallo Welt"
	if translatedText.Translations[0].Text != expected {
		t.Errorf("expected %q, got %q", expected, translatedText.Translations[0].Text)
	}
}

func TestNewClientConfig(t *testing.T) {
	config := NewClientConfig("test-api-key:fx")
	if config.APIKey != "test-api-key:fx" {
		t.Errorf("expected APIKey to be 'test-api-key:fx', got %q", config.APIKey)
	}
	if config.IsPremiumMember {
		t.Errorf("expected IsPremiumMember to be false, got true")
	}

	config = NewClientConfig("test-api-key")
	if config.APIKey != "test-api-key" {
		t.Errorf("expected APIKey to be 'test-api-key', got %q", config.APIKey)
	}
	if !config.IsPremiumMember {
		t.Errorf("expected IsPremiumMember to be true, got false")
	}
}

func TestNewClient(t *testing.T) {
	config := NewClientConfig("test-api-key")
	client := NewClient(config)
	if client.APIKey != "test-api-key" {
		t.Errorf("expected APIKey to be 'test-api-key', got %q", client.APIKey)
	}
	if client.Host != "https://api.deepl.com" {
		t.Errorf("expected Host to be 'https://api.deepl.com', got %q", client.Host)
	}
	if client.HTTPClient == nil {
		t.Errorf("expected HTTPClient to be non-nil")
	}

	// Timeout 체크
	timeout := client.HTTPClient.Timeout
	if timeout != 10*time.Second {
		t.Errorf("expected Timeout to be 10 seconds, got %v", timeout)
	}

	config = NewClientConfig("test-api-key:fx")
	client = NewClient(config)
	if client.Host != "https://api-free.deepl.com" {
		t.Errorf("expected Host to be 'https://api-free.deepl.com', got %q", client.Host)
	}
}
