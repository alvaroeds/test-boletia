package currencyapi

import (
	"context"
	"encoding/json"
	"github.com/alvaroeds/test-boletia/pkg/request"
	"log"
	"net/http"
	"time"
)

type Repository interface {
	GetCurrencys(ctx context.Context) (*Latest, *QueryInfo, error)
	GetStatus(ctx context.Context) (*Status, error)
}

type repository struct {
	req   request.Request
	token string
}

func (r repository) GetCurrencys(ctx context.Context) (*Latest, *QueryInfo, error) {

	query := map[string]string{
		"apikey": r.token,
	}

	req, err := request.NewHttpRequest(http.MethodGet, "/v3/latest", nil, query, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	timeInit := time.Now()

	resp, code, err := r.req.Do(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	duration := time.Since(timeInit)

	info := &QueryInfo{
		Method: req.Method,
		Path:   req.URL.Path,
		Code:   code,
		Time:   duration.Seconds(),
	}

	data := Latest{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return nil, info, err
	}

	return &data, info, nil
}

func (r repository) GetStatus(ctx context.Context) (*Status, error) {

	query := map[string]string{
		"apikey": r.token,
	}

	req, err := request.NewHttpRequest(http.MethodGet, "/v3/status", nil, query, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, _, err := r.req.Do(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	data := Status{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func newRepository(req request.Request, token string) Repository {
	return repository{
		req:   req,
		token: token,
	}
}
