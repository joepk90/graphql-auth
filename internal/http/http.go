package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/joepk90/graphql-auth/internal/auth/iam"
)

type HttpService struct {
	baseURL string
	client  *http.Client
}

func NewHttpService(authURL string) *HttpService {
	return &HttpService{
		baseURL: authURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *HttpService) NewRequestWithContext(ctx context.Context, method, path string, reqBody any, resBody any) error {

	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, a.baseURL+path, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	principle := iam.PrincipalFromCtx(ctx)

	if principle.Token != "" {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+principle.Token)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("http error: %s - %s", resp.Status, string(data))
	}

	if resBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(resBody); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
