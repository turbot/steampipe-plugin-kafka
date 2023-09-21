package kafka

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		errMsg := err.Error()
		for _, msg := range retryErrors {
			if strings.Contains(errMsg, msg) {
				return true
			}
		}
		return false
	}
}
