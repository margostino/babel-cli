package db

type Metadata struct {
	Category     string   `json:"category"`
	Highlights   []string `json:"highlights"`
	Keywords     []string `json:"keywords"`
	Path         string   `json:"path"`
	References   []string `json:"references"`
	RelatedLinks []string `json:"related_links"`
	Summary      string   `json:"summary"`
	Tags         []string `json:"tags"`
}

type Asset struct {
	Content  string
	Metadata Metadata
}

type SemanticSearchResult struct {
	Category string
	Path     string
	Summary  string
}

type GenerativeSearchResult struct {
	SingleResult string
	Error        string
}

type RawGenerativeSearchResult struct {
	Additional struct {
		Generate GenerativeSearchResult `json:"generate"`
	} `json:"_additional"`
	Category string `json:"category"`
	Path     string `json:"path"`
	Summary  string `json:"summary"`
}
