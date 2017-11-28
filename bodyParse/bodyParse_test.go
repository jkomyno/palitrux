package bodyParse

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func TestIsMultipart(t *testing.T) {
	// recorder := httptest.NewRecorder()
	h := http.Header{}
	h.Set("Content-Type", "multipart/formdata")
	req := &http.Request{
		Method:        "POST",
		URL:           &url.URL{Path: "/"},
		MultipartForm: &multipart.Form{},
		Header:        h,
	}

	res := isMultipart(req)

	if !res {
		t.Error("res should be multipart")
	}
}
