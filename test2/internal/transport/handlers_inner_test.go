package transport

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bejaneps/rollee-assignment/test2/internal/storage"
	"github.com/bejaneps/rollee-assignment/test2/internal/transport/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleServicePost(t *testing.T) {
	testCases := map[string]struct {
		setup              func() (*mocks.Service, *http.Request)
		expectedStatusCode int
		expectedBody       string
	}{
		"success": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"UpsertWord",
					"hello",
				).Return(
					nil,
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/service/word="hello"`,
					nil,
				)

				req = mux.SetURLVars(req, map[string]string{
					"word": "hello",
				})

				return svc, req
			},
			expectedStatusCode: 200,
			expectedBody:       "",
		},
		"fail": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"UpsertWord",
					"hello",
				).Return(
					errors.New("random error"),
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/service/word="hello"`,
					nil,
				)

				req = mux.SetURLVars(req, map[string]string{
					"word": "hello",
				})

				return svc, req
			},
			expectedStatusCode: 500,
			expectedBody:       "Internal Server Error\n",
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			svc, req := tt.setup()
			rr := httptest.NewRecorder()

			handler := handleServicePost(svc)
			handler(rr, req)

			body, err := io.ReadAll(rr.Result().Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatusCode, rr.Result().StatusCode)
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}

func TestHandleServiceGet(t *testing.T) {
	testCases := map[string]struct {
		setup              func() (*mocks.Service, *http.Request)
		expectedStatusCode int
		expectedBody       string
	}{
		"success": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"hello",
					nil,
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/service/word="hel"`,
					nil,
				)

				req = mux.SetURLVars(req, map[string]string{
					"word": "hel",
				})

				return svc, req
			},
			expectedStatusCode: 200,
			expectedBody:       "hello",
		},
		"fail": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"",
					errors.New("random error"),
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/service/word="hel"`,
					nil,
				)

				req = mux.SetURLVars(req, map[string]string{
					"word": "hel",
				})

				return svc, req
			},
			expectedStatusCode: 500,
			expectedBody:       "Internal Server Error\n",
		},
		"fail-word-not-found": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"",
					storage.ErrWordNotFound,
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/service/word="hel"`,
					nil,
				)

				req = mux.SetURLVars(req, map[string]string{
					"word": "hel",
				})

				return svc, req
			},
			expectedStatusCode: 400,
			expectedBody:       "null\n",
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			svc, req := tt.setup()
			rr := httptest.NewRecorder()

			handler := handleServiceGet(svc)
			handler(rr, req)

			body, err := io.ReadAll(rr.Result().Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatusCode, rr.Result().StatusCode)
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}

func TestHandleNotFound(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		`/path/doesnt/exist`,
		nil,
	)

	rr := httptest.NewRecorder()

	handler := handleNotFound()
	handler(rr, req)

	body, err := io.ReadAll(rr.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, 400, rr.Result().StatusCode)
	assert.Equal(
		t, 
		"Specified path doesn't exist or word is invalid.\n",
		string(body),
	)
}
