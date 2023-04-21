package negotiate_response

import (
	"github.com/unrolled/render"
	"io"
	"net/http"
)

// Negotiator ContentType 값에 따라 응답 포멧을 다르게 하는 구조체
type Negotiator struct {
	ContentType string
	*render.Render
}

func GetNegotiator(r *http.Request) *Negotiator {
	contentType := r.Header.Get("Content-Type")
	return &Negotiator{
		ContentType: contentType,
		Render:      render.New(),
	}
}

func (n *Negotiator) Respond(w io.Writer, status int, v interface{}) {
	switch n.ContentType {
	case render.ContentJSON:
		n.Render.JSON(w, status, v)
	case render.ContentXML:
		n.Render.XML(w, status, v)
	default:
		n.Render.JSON(w, status, v)
	}
}
