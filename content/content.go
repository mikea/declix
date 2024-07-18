package content

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func OpenContent(pkl any) (io.ReadCloser, int64, error) {
	switch c := pkl.(type) {
	case *File:
		f, err := os.Open(c.File)
		if err != nil {
			return nil, 0, err
		}
		stat, err := f.Stat()
		if err != nil {
			return nil, 0, err
		}
		return f, stat.Size(), nil
	case string:
		return noCloseReadCloser{strings.NewReader(c)}, int64(len(c)), nil
	case *Url:
		resp, err := http.Get(c.Url)
		if err != nil {
			return nil, -1, err
		}
		return resp.Body, resp.ContentLength, nil
	default:
		panic(fmt.Sprintf("unsupported content %T", c))
	}
}

type noCloseReadCloser struct {
	reader io.Reader
}

// Close implements io.ReadCloser.
func (noCloseReadCloser) Close() error {
	return nil
}

// Read implements io.ReadCloser.
func (r noCloseReadCloser) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
