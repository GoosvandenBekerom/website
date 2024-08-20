package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/goosvandenbekerom/website/logger/colors"
)

const timeFormat = "[2006-01-02 15:04:05.000]"

// PrettyPrintHandler pretty prints a human-readable version of the structured logs to stdout
type PrettyPrintHandler struct {
	level      slog.Leveler
	attributes []slog.Attr
	group      string
}

func NewPrettyPrintHandler(level slog.Leveler) slog.Handler {
	return &PrettyPrintHandler{
		level: level,
	}
}

func (s *PrettyPrintHandler) Enabled(_ context.Context, l slog.Level) bool {
	minLevel := slog.LevelInfo
	if s.level != nil {
		minLevel = s.level.Level()
	}
	return l >= minLevel
}

func (s *PrettyPrintHandler) Handle(_ context.Context, r slog.Record) error {
	r.AddAttrs(s.attributes...)
	if s.group != "" {
		r.Add(slog.String("group", s.group))
	}

	var attrs strings.Builder
	r.Attrs(func(attr slog.Attr) bool {
		fmt.Fprintf(&attrs, " %s;", attr)
		return true
	})

	level := "[" + r.Level.String() + "]"

	switch r.Level {
	case slog.LevelDebug:
		level = colors.Convert(colors.DarkGray, level)
	case slog.LevelInfo:
		level = colors.Convert(colors.Cyan, level)
	case slog.LevelWarn:
		level = colors.Convert(colors.LightYellow, level)
	case slog.LevelError:
		level = colors.Convert(colors.LightRed, level)
	}

	_, err := fmt.Fprintf(os.Stdout, "%s %s %s%s\n",
		r.Time.Format(timeFormat),
		level,
		colors.Convert(colors.White, r.Message),
		colors.Convert(colors.LightGray, attrs.String()))

	return err
}

func (s *PrettyPrintHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyPrintHandler{
		level:      s.level,
		attributes: attrs,
		group:      s.group,
	}
}

func (s *PrettyPrintHandler) WithGroup(name string) slog.Handler {
	return &PrettyPrintHandler{
		level:      s.level,
		attributes: s.attributes,
		group:      name,
	}
}
