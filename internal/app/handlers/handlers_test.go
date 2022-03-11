package handlers

import (
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHandler_HandlerGetURLByID(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	type fields struct {
		repo    memory.MemoryInterface
		baseURL string
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		fields fields
		args   args
		name   string
		query  string
		err    error
		result string
		want   want
	}{
		{
			name:   "GET with correct id",
			query:  "98fv58Wr3hGGIzm2-aH2zA628Ng=",
			result: "98fv58Wr3hGGIzm2-aH2zA628Ng=",
			err:    nil,
			want: want{
				code:        307,
				response:    ``,
				contentType: `text/plain; charset=utf-8`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				repo:    tt.fields.repo,
				baseURL: tt.fields.baseURL,
			}
			h.HandlerGetURLByID(tt.args.c)
		})
	}
}
