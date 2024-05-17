package deepl

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTranslate(t *testing.T) {
	// 테스트 서버 생성
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 모의 응답 설정
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"translations":[{"detected_source_language":"EN","text":"Hallo Welt"}]}`))
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer ts.Close()

	client := &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		APIKey:     "test-api-key",
		Host:       ts.URL,
	}

	translationRequest := NewTranslationRequest(
		[]string{"Hello world"},
		"DE",
		WithSourceLang("EN"),
		WithFormality("more"),
	)

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
