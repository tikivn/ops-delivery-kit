package email_service

import (
	"context"
	"io"
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"
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

type storageError struct {
	Message string `json:"message"`
}

type StorageService interface {
	UploadFile(ctx context.Context, bucket, filename string, file io.Reader) (*UploadedFile, error)
}

var _ = StorageService(&storageService{})

func NewStorageService(
	host string,
	_ http.RoundTripper,
) StorageService {
	httpClient := &http.Client{
		Transport: &ochttp.Transport{},
		Timeout:   3 * time.Minute,
	}

	return &storageService{
		client:     resty.New().SetHostURL(host),
		Host:       host,
		httpClient: httpClient,
	}
}

type storageService struct {
	client     *resty.Client
	Host       string
	httpClient *http.Client
}

func (r *storageError) Error() string {
	return "Storage: " + r.Message
}

func (s *storageService) UploadFile(ctx context.Context, bucket, filename string, reader io.Reader) (*UploadedFile, error) {
	req := s.client.R().
		SetContext(ctx).
		SetQueryParam("bucket", bucket).
		// SetQueryParam("filename", filename).
		// SetQueryParam("url", "https://salt.tikicdn.com/ts/banner/7e/55/26/anduc/data.zip").
		SetFileReader("file", filename, reader).
		SetResult(envelope{})

	resp, err := req.Post("v1/objects")

	if err != nil {
		return nil, err
	}

	body := resp.Result().(*envelope)

	// fmt.Println(body)
	if body.Err != nil {
		return nil, body.Err
	}

	return body.UploadedFile, nil
}

type envelope struct {
	*UploadedFile
	Err *storageError `json:"error,omitempty"`
}
