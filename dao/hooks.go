package dao

import (
	"context"
	"strings"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/duke-git/lancet/v2/strutil"
)

// Hooks satisfies the sqlhook.Hooks interface
type Hooks struct{}

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	query = strings.ReplaceAll(query, "--sql", "")
	query = strutil.RemoveWhiteSpace(query, false)
	logger.Info("SQL> ", query)
	logger.Info("SQL> ", args)
	return ctx, nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	// do something...
	return ctx, nil
}
