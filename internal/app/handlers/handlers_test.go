package handlers

//
//import (
//	"encoding/json"
//	"github.com/go-chi/chi"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestHandler_HandlerGetURLByID(t *testing.T) {
//	store := make(map[string]string)
//
//	store["1"] = "http://test.com"
//
//	r := chi.NewRouter()
//	r.Get("/{articleID}", func(w http.ResponseWriter, r *http.Request) {
//		key := chi.URLParam(r, "articleID")
//		url := store[key]
//		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//		w.WriteHeader(307)
//		w.Write([]byte(url))
//	})
//
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	if res, _ := TestRequest(t, ts, "GET", "/1", nil); res.StatusCode != 307 {
//		t.Fatalf("want %d, got %d", 307, res.StatusCode)
//	}
//
//	if res, _ := TestRequest(t, ts, "GET", "/something/666", nil); res.StatusCode != 404 {
//		t.Fatalf("want %d, got %d", 404, res.StatusCode)
//	}
//
//	if res, _ := TestRequest(t, ts, "GET", "/1", nil); res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
//		t.Fatalf("want %s, got %s", "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
//	}
//
//	if _, body := TestRequest(t, ts, "GET", "/1", nil); string(body) != "http://test.com" {
//		t.Fatalf("want %s, got %s", "http://google.com", string(body))
//	}
//}
//
//func TestHandler_HandlerCreateShortURL(t *testing.T) {
//	r := chi.NewRouter()
//	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//		w.WriteHeader(201)
//		w.Write([]byte("http://localhost:8080/whTHc"))
//	})
//
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	if _, body := TestRequest(t, ts, "POST", "/", nil); string(body) != "http://localhost:8080/whTHc" {
//		t.Fatalf("want %s, got %s", "http://localhost:8080/whTHc", string(body))
//	}
//
//	if res, _ := TestRequest(t, ts, "POST", "/", nil); res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
//		t.Fatalf("want %s, got %s", "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
//	}
//
//	if res, _ := TestRequest(t, ts, "POST", "/", nil); res.StatusCode != 201 {
//		t.Fatalf("want %d, got %d", 201, res.StatusCode)
//	}
//
//	if res, _ := TestRequest(t, ts, "POST", "/somewrong", nil); res.StatusCode != 404 {
//		t.Fatalf("want %d, got %d", 404, res.StatusCode)
//	}
//}
//
//type shorttest struct {
//	Key string `json:"result"`
//}
//
//type bodytest struct {
//	Result string `json:"result"`
//}
//
//func TestHandler_HandlerShortenURL(t *testing.T) {
//	short := shorttest{
//		Key: "http://localhost:8080/1",
//	}
//
//	r := chi.NewRouter()
//	r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		w.WriteHeader(201)
//		result, err := json.Marshal(&short)
//		if err != nil {
//			t.Error(err.Error())
//			return
//		}
//
//		w.Write([]byte(result))
//	})
//
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	if res, _ := TestRequest(t, ts, "POST", "/api/shorten", nil); res.StatusCode != 201 {
//		t.Fatalf("want %d, got %d\n", 201, res.StatusCode)
//	}
//
//	if res, _ := TestRequest(t, ts, "POST", "/api/shorten", nil); res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
//		t.Fatalf("want %s, got %s", "application/json; charset=utf-8", res.Header.Get("Content-Type"))
//	}
//
//	if res, _ := TestRequest(t, ts, "POST", "/api/shorten/3", nil); res.StatusCode != 404 {
//		t.Fatalf("want %d, got %d", 404, res.StatusCode)
//	}
//
//	bodytest := bodytest{}
//
//	_, body := TestRequest(t, ts, "POST", "/api/shorten", nil)
//
//	err := json.Unmarshal(body, &bodytest)
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//
//	bodyRes := bodytest.Result
//	if string(bodyRes) != "http://localhost:8080/1" {
//		t.Fatalf("want %s, got %s", "http://localhost:8080/1", bodyRes)
//	}
//}
