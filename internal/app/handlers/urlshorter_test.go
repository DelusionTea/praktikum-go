
import (
	"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCreateShortURL(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		request string
	}{
		// TODO: Add test cases.
		{
            name: "positive test",
			request "/",
            want: want{
                code:        201,
            },
        },
		{
            name: "negative test",
			request "/wrong",
            want: want{
                code:        404,
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.POST("/", HandlerCreateShortURL)

			req := httptest.NewRequest(http.MethodPost, tt.request, nil)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			result := rr.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}

func TestHandlerGetURLByID(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name string
		request string
		long string
		id int
		mapURLs map[int]longShortURLs
		args args
	}{
		// TODO: Add test cases.
		{
            name: "positive test #1",
			request "/1",
			id: 1,
			long: "http://somesite.com"
            want: want{
                code:        307,
                contentType: "text/plain",
            },
        },
		{
            name: "not correct",
			request "/2",
			id: 2,
			long: "http://anothersomesite.com"
            want: want{
                code:        400,
                contentType: "text/plain",
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				short := Shorter(1)
				shorts[tt.id] = URLs{
					Long: tt.long,
					Short: short,
				}

				router := httprouter.New()
				router.GET("/:id", HandlerGetURLByID)

				req := httptest.NewRequest(http.MethodGet, tt.request, nil)

				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				result := rr.Result()
				defer result.Body.Close()
				assert.Equal(t, tt.want.statusCode, result.StatusCode)
			}
		})
	}
}