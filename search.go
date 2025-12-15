package resultx

import (
	"strings"
)

type SearchRequest struct {
	Query string `json:"query"`
}

func NewSearchRequest(query string) SearchRequest {
	return SearchRequest{
		Query: query,
	}
}

func (s SearchRequest) QueryValue() string {
	return strings.TrimSpace(s.Query)
}
