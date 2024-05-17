package deepl

type TranslationRequest struct {
	Text               []string `json:"text"`                          // Text to be translated. Required.
	SourceLang         string   `json:"source_lang,omitempty"`         // Language of the text to be translated. Optional.
	TargetLang         string   `json:"target_lang"`                   // The language into which the text should be translated. Required.
	Context            string   `json:"context,omitempty"`             // Additional context that can influence a translation. Optional.
	SplitSentences     string   `json:"split_sentences,omitempty"`     // Sets whether the translation engine should first split the input into sentences. Optional.
	PreserveFormatting bool     `json:"preserve_formatting,omitempty"` // Sets whether the translation engine should respect the original formatting. Optional.
	Formality          string   `json:"formality,omitempty"`           // Sets whether the translated text should lean towards formal or informal language. Optional.
	GlossaryID         string   `json:"glossary_id,omitempty"`         // Specify the glossary to use for the translation. Optional.
	TagHandling        string   `json:"tag_handling,omitempty"`        // Sets which kind of tags should be handled. Optional.
	OutlineDetection   bool     `json:"outline_detection,omitempty"`   // Disables the automatic detection of the XML structure. Optional.
	NonSplittingTags   []string `json:"non_splitting_tags,omitempty"`  // Comma-separated list of XML or HTML tags which never split sentences. Optional.
	SplittingTags      []string `json:"splitting_tags,omitempty"`      // Comma-separated list of XML or HTML tags which always cause splits. Optional.
	IgnoreTags         []string `json:"ignore_tags,omitempty"`         // Comma-separated list of XML or HTML tags that indicate text not to be translated. Optional.
}

func NewTranslationRequest(text []string, targetLang string, options ...func(*TranslationRequest)) *TranslationRequest {
	req := &TranslationRequest{
		Text:       text,
		TargetLang: targetLang,
	}

	for _, opt := range options {
		opt(req)
	}

	return req
}

// WithSourceLang sets the SourceLang field of the TranslationRequest.
func WithSourceLang(sourceLang string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.SourceLang = sourceLang
	}
}

// WithContext sets the Context field of the TranslationRequest.
func WithContext(context string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.Context = context
	}
}

// WithSplitSentences sets the SplitSentences field of the TranslationRequest.
func WithSplitSentences(splitSentences string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.SplitSentences = splitSentences
	}
}

// WithPreserveFormatting sets the PreserveFormatting field of the TranslationRequest.
func WithPreserveFormatting(preserveFormatting bool) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.PreserveFormatting = preserveFormatting
	}
}

// WithFormality sets the Formality field of the TranslationRequest.
func WithFormality(formality string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.Formality = formality
	}
}

// WithGlossaryID sets the GlossaryID field of the TranslationRequest.
func WithGlossaryID(glossaryID string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.GlossaryID = glossaryID
	}
}

// WithTagHandling sets the TagHandling field of the TranslationRequest.
func WithTagHandling(tagHandling string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.TagHandling = tagHandling
	}
}

// WithOutlineDetection sets the OutlineDetection field of the TranslationRequest.
func WithOutlineDetection(outlineDetection bool) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.OutlineDetection = outlineDetection
	}
}

// WithNonSplittingTags sets the NonSplittingTags field of the TranslationRequest.
func WithNonSplittingTags(nonSplittingTags []string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.NonSplittingTags = nonSplittingTags
	}
}

// WithSplittingTags sets the SplittingTags field of the TranslationRequest.
func WithSplittingTags(splittingTags []string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.SplittingTags = splittingTags
	}
}

// WithIgnoreTags sets the IgnoreTags field of the TranslationRequest.
func WithIgnoreTags(ignoreTags []string) func(*TranslationRequest) {
	return func(req *TranslationRequest) {
		req.IgnoreTags = ignoreTags
	}
}
