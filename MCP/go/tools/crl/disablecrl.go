package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/iam-roles-anywhere/mcp-server/config"
	"github.com/iam-roles-anywhere/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DisablecrlHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		crlIdVal, ok := args["crlId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: crlId"), nil
		}
		crlId, ok := crlIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: crlId"), nil
		}
		url := fmt.Sprintf("%s/crl/%s/disable", cfg.BaseURL, crlId)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CrlDetailResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDisablecrlTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_crl_crlId_disable",
		mcp.WithDescription("<p>Disables a certificate revocation list (CRL).</p> <p> <b>Required permissions: </b> <code>rolesanywhere:DisableCrl</code>. </p>"),
		mcp.WithString("crlId", mcp.Required(), mcp.Description("The unique identifier of the certificate revocation list (CRL).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DisablecrlHandler(cfg),
	}
}
