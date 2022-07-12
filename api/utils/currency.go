package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	dtos "github.com/truecoder34/user-balance-service/api/DTOs"
)

func ConvertFromRub(value int64, currencyConvertTo string) float64 {
	url := "https://api.apilayer.com/exchangerates_data/convert?to=" + currencyConvertTo + "&from=RUB&amount=" + strconv.Itoa(int(value))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	os.Getenv("API_LAYER_KEY")
	req.Header.Set("apikey", os.Getenv("API_LAYER_KEY"))

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)

	var exchangeRates dtos.ExchangeRatesDTO
	json.Unmarshal(body, &exchangeRates)
	fmt.Println(exchangeRates)

	return exchangeRates.Result

}
