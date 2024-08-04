package render

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

// A struct that represents renderable Templ template.
type Templ struct {
	ctx       context.Context
	component templ.Component
}

func New(ctx context.Context, component templ.Component) *Templ {
	return &Templ{
		ctx:       ctx,
		component: component,
	}
}

func (t *Templ) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	if t.component != nil {
		return t.component.Render(t.ctx, w)
	}
	return nil
}

func (t *Templ) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
