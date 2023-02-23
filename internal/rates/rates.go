package rates

import (
	"context"
	"encoding/json"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	"github.com/kjasuquo/usdngn-exchange/internal/services/web"
	"github.com/zeebo/errs"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	Http   http.Client
	Config config.Config
}

func NewRatesClient(conf config.Config) *Client {
	return &Client{Config: conf}
}

func (client *Client) GetRates(ctx context.Context) (models.Rates, error) {

	var rateResponse models.Rates

	apiUrl := client.Config.ExchangeRate

	request, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Println("error calling request: ", err)
		return rateResponse, err
	}

	// Add Basic headers
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Http.Do(request.WithContext(ctx))
	if err != nil {
		log.Println("error getting response: ", err)
		return rateResponse, err
	}

	defer func() {
		err = errs.Combine(err, response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated &&
		response.StatusCode != http.StatusAccepted {
		errorResp := web.ErrorResponse{}

		if err = json.NewDecoder(response.Body).Decode(&errorResp); err != nil {
			log.Println("error decoding response: ", err)
			return rateResponse, err
		}

		return rateResponse, errorResp
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error reading response.body to byte: ", err)
		return rateResponse, err
	}

	err = json.Unmarshal(resp, &rateResponse)
	if err != nil {
		log.Println("error unmarshalling response: ", err)
		return rateResponse, err
	}

	return rateResponse, nil
}
