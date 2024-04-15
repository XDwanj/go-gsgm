package dao

import (
	"context"
	"strings"

	"github.com/XDwanj/go-gsgm/logger"
	"github.com/duke-git/lancet/v2/strutil"
)

type Hooks struct{}

func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	query = strings.Replace(query, "--sql", "", 1)
	query = strutil.RemoveWhiteSpace(query, false)
	logger.Info("SQL> ", query)
	logger.Info("SQL> args=", args)
	return ctx, nil
}

func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	// do something...
	return ctx, nil
}
