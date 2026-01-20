package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const HeaderRequestID = "X-Request-Id"
type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type AnalyzeRequest struct {
	Text      string            `json:"text"`
	Source    string            `json:"source,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	EntryID   string            `json:"entry_id,omitempty"`
	CreatedAt string            `json:"created_at,omitempty"`
}


type AnalyzeResponse struct {
	EntryID    string                 `json:"entry_id"`
	CreatedAt  string                 `json:"created_at"`
	Analysis   map[string]any         `json:"analysis"`
	MockNotice string                 `json:"mock_notice,omitempty"`
}

func (c *Client) Analyze(ctx context.Context, requestID string, req AnalyzeRequest) (AnalyzeResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return AnalyzeResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/analyze", bytes.NewReader(body))
	if err != nil {
		return AnalyzeResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if requestID != "" {
		httpReq.Header.Set(HeaderRequestID, requestID)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return AnalyzeResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return AnalyzeResponse{}, fmt.Errorf("agent returned status %d", resp.StatusCode)
	}

	var out AnalyzeResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return AnalyzeResponse{}, err
	}

	return out, nil
}
