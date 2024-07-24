package content

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Sha256(pkl any) (string, error) {
	switch c := pkl.(type) {
	case *File:
		if c.Sha256 != nil {
			return *c.Sha256, nil
		}
		f, err := os.Open(c.File)
		if err != nil {
			return "", err
		}
		return readAndHash(f)
	case *Url:
		if c.Sha256 != nil {
			return *c.Sha256, nil
		}
		resp, err := http.Get(c.Url)
		if err != nil {
			return "", err
		}
		return readAndHash(resp.Body)
	case *Base64:
		if c.Sha256 != nil {
			return *c.Sha256, nil
		}
		panic("not implemented")
	case string:
		return readAndHash(noCloseReadCloser{strings.NewReader(c)})
	default:
		panic(fmt.Sprintf("unsupported content %T", c))
	}
}

func readAndHash(reader io.ReadCloser) (string, error) {
	defer reader.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", string(hasher.Sum(nil))), nil
}

func Open(pkl any) (io.ReadCloser, int64, error) {
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
	case *Base64:
		content, err := base64.StdEncoding.DecodeString(c.Base64)
		if err != nil {
			return nil, -1, err
		}
		return noCloseReadCloser{bytes.NewReader(content)}, int64(len(content)), nil
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
