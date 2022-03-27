//package handlers
//
//import (
//	"encoding/json"
//	"github.com/DelusionTea/praktikum-go/cmd/conf"
//	"github.com/DelusionTea/praktikum-go/internal/memory"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"regexp"
//	"strings"
//	"testing"
//)
//
//func setupRouter(repo memory.MemoryMap, baseURL string) *gin.Engine {
//	router := gin.Default()
//
//	handler := New(repo, baseURL)
//
//	router.GET("/:id", handler.HandlerGetURLByID)
//	router.POST("/", handler.HandlerCreateShortURL)
//	router.POST("/api/shorten", handler.HandlerShortenURL)
//
//	router.HandleMethodNotAllowed = true
//
//	return router
//}
//
//func SendTestRequest(t *testing.T, ts *httptest.Server, method, path, contentType string, body io.Reader) (*http.Response, string) {
//	client := &http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		}}
//
//	req, err := http.NewRequest(method, ts.URL+path, body)
//	require.NoError(t, err)
//	if contentType != "" {
//		req.Header.Add("Content-Type", contentType)
//	}
//	resp, err := client.Do(req)
//	require.NoError(t, err)
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	require.NoError(t, err)
//
//	defer resp.Body.Close()
//	return resp, string(respBody)
//}
//
//type ShortenerResponse struct {
//	Result string `json:"result"`
//}
//
//type want struct {
//	responseStatusCode  int
//	responseParams      map[string]string
//	responseContentType string
//	responseBody        string
//}
//type request struct {
//	url, method, contentType, body string
//}
//
//const testAddr = "http://localhost:8080/"
//
//func TestHandler_HandlerGetURLByID(t *testing.T) {
//	type fields struct {
//		repo    memory.MemoryMap
//		baseURL string
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name    string
//		request request
//		want    want
//	}{
//		{
//			name: "negative test #1. GET with empty url",
//			request: request{
//				url:    "/",
//				method: http.MethodGet,
//				body:   "",
//			},
//			want: want{
//				responseStatusCode: http.StatusMethodNotAllowed,
//				responseParams:     nil,
//				responseBody:       "405 method not allowed",
//			},
//		},
//	}
//	//cfg := conf.GetConfig()
//	r := setupRouter(memory.NewMemoryFile("sorter.txt"), testAddr)
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			resp, body := SendTestRequest(t, ts, tt.request.method, tt.request.url, "", nil)
//			defer resp.Body.Close()
//			assert.Equal(t, tt.want.responseStatusCode, resp.StatusCode)
//			assert.Equal(t, tt.want.responseBody, strings.TrimRight(body, "\n"))
//		})
//	}
//}
//
//func TestHandler_HandlerCreateShortURL(t *testing.T) {
//
//	tests := []struct {
//		name    string
//		request request
//		want    want
//	}{
//		{
//			name: "positive test #1. POST",
//			request: request{
//				url:    "/",
//				method: http.MethodPost,
//				body:   "body:http://test.ru",
//			},
//			want: want{
//				responseStatusCode: http.StatusCreated,
//				responseParams:     nil,
//				responseBody:       "http://([a-zA-Z1-9]{5})",
//			},
//		},
//	}
//
//	cfg := conf.GetConfig()
//	r := setupRouter(memory.NewMemoryFile(cfg.FilePath), cfg.BaseURL)
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			resp, body := SendTestRequest(t, ts, tt.request.method, tt.request.url, "", strings.NewReader(tt.request.body))
//			defer resp.Body.Close()
//
//			assert.Equal(t, tt.want.responseStatusCode, resp.StatusCode)
//
//			//assert.True(t, resp.StatusCode == tt.want.responseStatusCode)
//
//			matched, _ := regexp.MatchString(tt.want.responseBody, strings.TrimRight(body, "\n"))
//			assert.True(t, matched)
//			//assert.Equal(t, tt.want.responseBody, matched)
//		})
//	}
//}
//
//func TestHandler_HandlerShortenURL(t *testing.T) {
//
//	tests := []struct {
//		name    string
//		request request
//		want    want
//	}{
//		{
//			name: "positive test #1. POST /api/shorten",
//			request: request{
//				url:         "/api/shorten",
//				contentType: "application/json",
//				method:      http.MethodPost,
//				body:        `{"url": "http://test.ru"}`,
//			},
//			want: want{
//				responseStatusCode:  201,
//				responseParams:      nil,
//				responseContentType: "application/json",
//				responseBody:        "http://([a-zA-Z1-9]{5})",
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := setupRouter(memory.NewMemoryFile("sorter.log"), testAddr)
//			ts := httptest.NewServer(r)
//			defer ts.Close()
//			resp, body := SendTestRequest(t, ts, tt.request.method, tt.request.url, tt.request.contentType, strings.NewReader(tt.request.body))
//			var originalLink ShortenerResponse
//			if err := json.Unmarshal([]byte(body), &originalLink); err != nil {
//				assert.True(t, false)
//			}
//			defer resp.Body.Close()
//
//			//assert.Equal(t, resp.Header.Get("Content-Type"), tt.want.responseContentType)
//			assert.Equal(t, tt.want.responseStatusCode, resp.StatusCode)
//
//			matched, _ := regexp.MatchString(tt.want.responseBody, originalLink.Result)
//			assert.True(t, matched)
//		})
//	}
//}
