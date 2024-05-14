# deepL-GO

[![Go Version](https://img.shields.io/github/go-mod/go-version/mohan12311/deepL-GO)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/mohan12311/deepL-GO)](https://github.com/mohan12311/deepL-GO/blob/main/LICENSE)

deepL-GO는 DeepL API를 사용하여 번역 기능을 제공하는 Go 언어 라이브러리입니다.

## 소개

deepL-GO는 DeepL의 번역 기능을 간편하게 사용할 수 있도록 도와주는 라이브러리입니다. 이 라이브러리를 사용하면 다양한 언어의 텍스트를 간단하게 번역할 수 있습니다.

## 설치

Go Modules를 사용하여 간편하게 설치할 수 있습니다.

```sh
go get github.com/mohan12311/deepL-GO
```

## 사용방법

아래는 deepL-GO 라이브러리를 사용하는 간단한 예제입니다.

```go
package main

import (
    "fmt"
    "log"

    "github.com/mohan12311/deepL-GO/deepl"
)

func main() {
    apiKey := "YOUR_DEEPL_API_KEY"
    client := deepl.NewClient(apiKey)

    text := "Hello, world!"
    translatedText, err := client.Translate(text, "EN", "KO")
    if err != nil {
        log.Fatalf("Failed to translate text: %v", err)
    }

    fmt.Println("Translated Text:", translatedText)
}
```

