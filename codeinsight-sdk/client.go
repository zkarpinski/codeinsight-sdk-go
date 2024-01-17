package codeinsightsdk

type CodeInsightClient struct {
	baseURL      string
	apiURL       string
	apiToken     string
	apiHeaders   map[string]string
	timeout      int
	verifySSL    bool
	experimental bool
}

// NewCodeInsightClient creates a new CodeInsightClient
func NewCodeInsightClient(baseURL string, apiURL string, apiToken string) *CodeInsightClient {

	client := &CodeInsightClient{
		baseURL:    baseURL,
		apiURL:     apiURL,
		apiToken:   apiToken,
		apiHeaders: map[string]string{"Authorization": "Bearer " + apiToken},
		timeout:    10,
		verifySSL:  true,
	}

	return client
}
