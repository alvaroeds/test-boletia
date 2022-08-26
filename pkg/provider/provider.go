package provider

import (
	"database/sql"
	"github.com/alvaroeds/test-boletia/pkg/currency"
	"github.com/alvaroeds/test-boletia/pkg/info"
	"github.com/alvaroeds/test-boletia/pkg/request"
	"github.com/roylee0704/gron"
	"net/http"
	"time"

	"context"
	"fmt"
	"github.com/alvaroeds/test-boletia/internal/config"
	"github.com/alvaroeds/test-boletia/pkg/provider/currencyapi"
	"log"
)

type Provider struct {
	cron        *gron.Cron
	currencyApi currencyapi.Service
	sCurrency   currency.Service
	sInfo       info.Service
	interval    int
}

func (p Provider) Load() {
	if p.cron == nil {
		return
	}

	go p.load()

	p.cron.AddFunc(gron.Every(time.Duration(p.interval)*time.Minute), p.load)
	p.cron.Start()
}

func (p Provider) load() {
	ctx := context.Background()

	status, err := p.currencyApi.GetStatus(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	if status.Quotas.Month.Remaining == 0 {
		err = fmt.Errorf("Quota api ended")
		log.Println(err)
		return
	}

	latestData, infoData, err := p.currencyApi.GetCurrencys(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, value := range latestData.Data {
		err = p.sCurrency.Insert(ctx, &currency.InsertCurrency{
			Code:  value.Code,
			Value: value.Value,
		})
		if err != nil {
			log.Println(err)
			continue
		}
	}

	err = p.sInfo.InsertRequest(ctx, &info.InsertRequest{
		Method: infoData.Method,
		Path:   infoData.Path,
		Code:   infoData.Code,
		Time:   infoData.Time,
	})
	if err != nil {
		log.Println(err)
	}
}

func New(conf *config.Config, db *sql.DB) *Provider {
	return &Provider{
		currencyApi: currencyapi.NewService(
			request.NewHTTP(
				conf.CurrencyApiURL,
				&http.Client{Timeout: time.Duration(conf.CurrencyApiTimeout) * time.Second},
				nil,
			),
			conf.CurrencyApiKEy),
		interval:  conf.CurrencyApiInterval,
		cron:      gron.New(),
		sCurrency: currency.NewService(db),
		sInfo:     info.NewService(db),
	}
}
