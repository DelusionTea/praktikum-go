package urlshorter_test

import "testing"

package handlers

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
		args args
	}{
		// TODO: Add test cases.
		{
            name: "positive test #1",
            want: want{
                code:        200,
                response:    `{"status":"ok"}`,
                contentType: "application/json",
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlerCreateShortURL(tt.args.w, tt.args.r, tt.args.in2)
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
		args args
	}{
		// TODO: Add test cases.
		{
            name: "positive test #1",
            want: want{
                code:        200,
                response:    `{"status":"ok"}`,
                contentType: "application/json",
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlerGetURLByID(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}