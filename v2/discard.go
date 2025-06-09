package pango

import (
	"context"
	"log/slog"
)

// discardHandler is an slog handler which is always disabled.
//
// This slog handler implementation is always disabled, and therefore
// logs nothing. It is used to filter out log categories not
// explicitly enabled.
type discardHandler struct{}

func (d discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (d discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (d discardHandler) WithAttrs(attrs []slog.Attr) slog.Handler  { return d }
func (d discardHandler) WithGroup(name string) slog.Handler        { return d }
