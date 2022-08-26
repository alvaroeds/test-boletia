package currency

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type getByCurrencyHandlerMock struct {
	getByCurrencyRes []GetAllCurrency
	getByCurrencyErr error
}

func (s getByCurrencyHandlerMock) Insert(ctx context.Context, req *InsertCurrency) error {
	// TODO implement me
	panic("implement me")
}

func (m getByCurrencyHandlerMock) GetByCurrency(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error) {
	return m.getByCurrencyRes, m.getByCurrencyErr
}

func TestHandler_GetCurrencyHandler(t *testing.T) {
	errNotFound := errors.New("not found")
	//q := "?finit=2022-08-14&fend=2022-08-17"
	tests := []struct {
		name       string
		statusCode int
		code       string
		end        string
		init       string
		resp       []GetAllCurrency
		err        error
	}{
		{
			name:       "Currency Failure",
			statusCode: http.StatusBadRequest,
			code:       "",
			resp:       nil,
			err:        nil,
		},
		{
			name:       "init date Failure",
			statusCode: http.StatusBadRequest,
			code:       "USD",
			end:        "",
			init:       "2022-08-x17",
			resp:       nil,
			err:        nil,
		},
		{
			name:       "end date Failure",
			statusCode: http.StatusBadRequest,
			code:       "USD",
			end:        "2022-08-x17",
			init:       "2022-08-17",
			resp:       nil,
			err:        nil,
		},
		{
			name: "success",
			//body:       []byte(`{"currency":"USD"}`),
			statusCode: http.StatusOK,
			code:       "USD",
			resp:       []GetAllCurrency{},
			err:        nil,
		},
		{
			name:       "failure",
			statusCode: http.StatusInternalServerError,
			code:       "d",
			resp:       nil,
			err:        errNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := getByCurrencyHandlerMock{
				getByCurrencyRes: test.resp,
				getByCurrencyErr: test.err,
			}

			w := httptest.NewRecorder()

			//q := "?finit=2022-08-14&fend=2022-08-17"
			r, err := http.NewRequest(http.MethodGet, "/currency/"+test.code+"?finit="+test.init+"&fend="+test.end, nil)
			if err != nil {
				require.NoError(t, err)
			}

			h := Handler{
				service: m,
			}

			mux := chi.NewMux()
			mux.Route("/currency", func(r chi.Router) {
				r.Get("/", h.GetCurrencyHandler)
				r.Get("/{currency}", h.GetCurrencyHandler)
			})
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}
}

func TestHandler(t *testing.T) {
	NewHandler(nil)
}
