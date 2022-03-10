package handlers

//func TestHandlerCreateShortURL(t *testing.T) {
//	type want struct {
//		code int
//	}
//
//	tests := []struct {
//		name    string
//		request string
//		want    want
//	}{
//		// TODO: Add test cases.
//		{
//			name:    "positive test",
//			request: "/",
//			want: want{
//				code: 201,
//			},
//		},
//		{
//			name:    "negative test",
//			request: "/wrong",
//			want: want{
//				code: 404,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			router := httprouter.New()
//			router.POST("/", HandlerCreateShortURL)
//
//			req := httptest.NewRequest(http.MethodPost, tt.request, nil)
//
//			rr := httptest.NewRecorder()
//
//			router.ServeHTTP(rr, req)
//
//			result := rr.Result()
//			defer result.Body.Close()
//
//			assert.Equal(t, tt.want.code, result.StatusCode)
//		})
//	}
//}
//
//func TestHandlerGetURLByID(t *testing.T) {
//	type want struct {
//		contentType string
//		code        int
//	}
//	tests := []struct {
//		name    string
//		request string
//		long    string
//		id      int
//		mapURLs map[int]longShortURLs
//		want    want
//	}{
//		// TODO: Add test cases.
//		{
//			name:    "positive test #1",
//			request: "/1",
//			id:      1,
//			long:    "http://somesite.com",
//			want: want{
//				code:        307,
//				contentType: "text/plain",
//			},
//		},
//		// {
//		//     name: "not correct",
//		// 	request: "/2",
//		// 	id: 2,
//		// 	long: "http://anothersomesite.com",
//		//     want: want{
//		//         code:        400,
//		//         contentType: "text/plain",
//		//     },
//		// },
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			for _, tt := range tests {
//				short := Shorter(1)
//				mapURLs[tt.id] = longShortURLs{
//					Long:  tt.long,
//					Short: short,
//				}
//
//				router := httprouter.New()
//				router.GET("/:id", HandlerGetURLByID)
//
//				req := httptest.NewRequest(http.MethodGet, tt.request, nil)
//
//				rr := httptest.NewRecorder()
//
//				router.ServeHTTP(rr, req)
//
//				result := rr.Result()
//				defer result.Body.Close()
//				assert.Equal(t, tt.want.code, result.StatusCode)
//			}
//		})
//	}
//}

//func TestHandler_HandlerShortenURL(t *testing.T) {
//	type fields struct {
//		repo    memory.MemoryInterface
//		baseURL string
//		result  BodyResponse
//	}
//	type want struct {
//		contentType string
//		response    string
//		code        int
//	}
//	type args struct {
//		w   http.ResponseWriter
//		r   *http.Request
//		in2 httprouter.Params
//	}
//	tests := []struct {
//		name    string
//		query   string
//		body    string
//		rawData string
//		result  string
//		fields  fields
//		args    args
//		want    want
//	}{
//		// TODO: Add test cases.
//		{
//			name:    "correct POST",
//			query:   "api/shorten",
//			body:    `{"url": "http://iloverestaurant.ru/"}`,
//			rawData: "http://iloverestaurant.ru/",
//			result:  "98fv58Wr3hGGIzm2-aH2zA628Ng=",
//			want: want{
//				code:        201,
//				response:    `{"result": "http://localhost:8080/98fv58Wr3hGGIzm2-aH2zA628Ng="}`,
//				contentType: `application/json; charset=utf-8`,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := &Handler{
//				repo:    tt.fields.repo,
//				baseURL: tt.fields.baseURL,
//				result:  tt.fields.result,
//			}
//			h.HandlerShortenURL(tt.args.w, tt.args.r, tt.args.in2)
//		})
//	}
//}
