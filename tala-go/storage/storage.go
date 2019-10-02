package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/tikivn/ops-delivery/util"
	"go.opencensus.io/plugin/ochttp"
)

type UploadedFile struct {
	Bucket       string `json:"bucket"`
	Key          string `json:"key"`
	Location     string `json:"location"`
	OriginalName string `json:"original_name"`
	PreviewUrl   string `json:"this_is_just_a_preview_url"`
	Url          string `json:"url"`
}

type Service interface {
	UploadFile(ctx context.Context, bucket, filename string, file io.Reader) (*UploadedFile, error)
}

func NewFileService(
	host string,
	transport http.RoundTripper,
) Service {
	// httpClient := &http.Client{
	// 	Transport: transport,
	// }

	// Use opencencus.io metrics
	httpClient := &http.Client{
		Transport: &ochttp.Transport{},
		Timeout:   5 * time.Minute,
	}

	return &service{
		Host:       host,
		httpClient: httpClient,
	}
}

type service struct {
	Host       string
	httpClient *http.Client
}

func (s *service) UploadFile(ctx context.Context, bucket, filename string, file io.Reader) (*UploadedFile, error) {
	req, err := newRequest(ctx, fmt.Sprintf("%s/v1/objects", s.Host), bucket, filename, file)
	if err != nil {
		return nil, err
	}

	resp, err := s.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	e := envelope{}
	if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return nil, err
	}
	if e.Err != nil {
		return nil, e.Err
	}

	return e.UploadedFile, nil
}

func newRequest(ctx context.Context, url, bucket, filename string, file io.Reader) (req *http.Request, err error) {
	body := &bytes.Buffer{}

	w := multipart.NewWriter(body)
	if err = w.WriteField("bucket", bucket); err != nil {
		w.Close()
		return nil, err
	}
	if part, err := w.CreateFormFile("file", filename); err != nil {
		w.Close()
		return nil, err
	} else {
		err = util.Copy(ctx, part, file)
		if err != nil {
			w.Close()
			return nil, err
		}
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	if req, err = http.NewRequest("POST", url, body); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	return req, nil
}

type storageError struct {
	Message string `json:"message"`
}

// Error implements error interface
func (r *storageError) Error() string {
	return "Storage: " + r.Message
}

type envelope struct {
	*UploadedFile
	Err *storageError `json:"error,omitempty"`
}
