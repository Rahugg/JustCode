package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hw8/config/crm_core"
	"hw8/internal/crm_core/entity"
	"io"
	"net/http"
	"time"
)

type ValidateTransport struct {
	config *crm_core.Configuration
}

func NewTransport(config *crm_core.Configuration) *ValidateTransport {
	return &ValidateTransport{
		config: config,
	}
}

type (
	GetUserResponse struct {
		CurrentUser *entity.User
		CurrentRole *entity.Role
	}
	SendValidateRequest struct {
		roles []interface{}
	}
)

func (vt *ValidateTransport) Validate(ctx context.Context, accessToken string, roles ...interface{}) (*GetUserResponse, error) {
	var response *GetUserResponse

	responseBody, err := vt.makeRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/user/validate/%s", accessToken),
		vt.config.Validate.Timeout,
		roles,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to makeRequest err: %w", err)
	}

	if err = json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshall response err: %w", err)
	}

	return response, nil
}

func (vt *ValidateTransport) makeRequest(
	ctx context.Context,
	httpMethod string,
	endpoint string,
	timeout time.Duration,
	roles ...interface{},
) (b []byte, err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	requestURL := vt.config.Validate.Host + endpoint

	requestJSON, err := json.Marshal(SendValidateRequest{
		roles: roles,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, httpMethod, requestURL, bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, fmt.Errorf("failed to NewRequestWithContext err: %w", err)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client making http request err: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body err: %w", err)
	}

	return body, nil
}
