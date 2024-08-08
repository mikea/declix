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

func CachedSha256(pkl Content) string {
	if render, ok := pkl.(Render); ok {
		return render.GetSha256()
	}

	switch c := pkl.(type) {
	case *Hashed:
		return c.Sha256
	case *File:
		if c.Sha256 != nil {
			return *c.Sha256
		}
		return "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	case *Url:
		if c.Sha256 != nil {
			return *c.Sha256
		}
		return "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	case *Base64:
		if c.Sha256 != nil {
			return *c.Sha256
		}
		panic("not implemented")
	case string:
		hash, err := readAndHash(noCloseReadCloser{strings.NewReader(c)})
		if err != nil {
			panic(fmt.Errorf("error computing string sha256 %w", err))
		}
		return hash
	default:
		panic(fmt.Sprintf("unsupported content %T", c))
	}
}

func Sha256(pkl Content) (hash string, err error) {
	if render, ok := pkl.(Render); ok {
		return render.GetSha256(), nil
	}

	switch c := pkl.(type) {
	case *Hashed:
		return c.Sha256, nil
	case *File:
		if c.Sha256 != nil {
			return *c.Sha256, nil
		}
		f, err := os.Open(c.File)
		if err != nil {
			return "", err
		}
		hash, err = readAndHash(f)
		if err != nil {
			c.Sha256 = &hash
		}
		return hash, err
	case *Url:
		if c.Sha256 != nil {
			return *c.Sha256, nil
		}
		resp, err := http.Get(c.Url)
		if err != nil {
			return "", err
		}
		hash, err = readAndHash(resp.Body)
		if err != nil {
			c.Sha256 = &hash
		}
		return hash, err
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
	if render, ok := pkl.(Render); ok {
		s := render.GetResult()
		return noCloseReadCloser{strings.NewReader(s)}, int64(len(s)), nil
	}

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

func Equal(a Content, b Content) (bool, error) {
	aHash, err := Sha256(a)
	if err != nil {
		return false, err
	}
	bHash, err := Sha256(b)
	if err != nil {
		return false, err
	}
	return aHash == bHash, nil
}
