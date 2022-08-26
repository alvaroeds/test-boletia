package currency

import (
	"database/sql"
	"fmt"
	"github.com/alvaroeds/test-boletia/pkg/response"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

const dateFormat = "2006-01-02"

type Handler struct {
	service Service
}

func NewHandler(db *sql.DB) Handler {
	return Handler{
		service: NewService(db),
	}
}

func (h *Handler) GetCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currency := chi.URLParam(r, "currency")
	init := r.URL.Query().Get("finit")
	end := r.URL.Query().Get("fend")

	var (
		req CurrencyFilterRequest
		err error
	)

	if currency == "" {
		err := fmt.Errorf("Debe ingresar un código de currency")
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	req.Code = currency

	if init != "" {
		req.FInit, err = time.Parse(dateFormat, init)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		if end != "" {
			req.FEnd, err = time.Parse(dateFormat, end)
			if err != nil {
				_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
				return
			}
		}
	}

	resp, err := h.service.GetByCurrency(ctx, req)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return

	}

	_ = response.JSON(w, r, http.StatusOK, resp)
}
