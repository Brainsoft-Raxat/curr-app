package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Brainsoft-Raxat/curr-app/models"
	"github.com/labstack/echo/v4"
)

func SaveCurrencyHandler(c echo.Context) error {
	target, err := getCurrencies()
	if err != nil {
		return err
	}

	rates := target.Rates

	currency := models.Currency{
		GBP: fmt.Sprintf("%.2f", rates.KZT / rates.GBP),
		RUB: fmt.Sprintf("%.2f", rates.KZT / rates.RUB),
		USD: fmt.Sprintf("%.2f", rates.KZT / rates.USD),
		EUR: fmt.Sprintf("%.2f", rates.KZT),
	}

	return c.JSON(http.StatusOK, currency)
}

func getCurrencies() (models.Target, error) {
	url := "http://api.exchangeratesapi.io/v1/latest?access_key=db3cf960ce2f00c47486c7513971fee1&base=EUR"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	target := models.Target{}
	err := json.Unmarshal([]byte(body), &target)
	if err != nil {
		log.Printf("%v", err)
		return models.Target{}, err
	}

	return target, nil
}