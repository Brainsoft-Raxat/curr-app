package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Brainsoft-Raxat/curr-app/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/currency", SaveCurrencyHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func SaveCurrencyHandler(c echo.Context) error {
	target, err := GetCurrencies()
	if err != nil {
		return err
	}

	currency := models.Currency{}

	currency.EUR = fmt.Sprintf("%f", target.Rates.KZT)
	currency.USD = fmt.Sprintf("%f", target.Rates.KZT/target.Rates.USD)
	currency.GBP = fmt.Sprintf("%f", target.Rates.KZT/target.Rates.GBP)
	currency.RUR = fmt.Sprintf("%f", target.Rates.KZT/target.Rates.RUB)

	return c.JSON(http.StatusOK, currency)
}

func GetCurrencies() (models.Target, error) {
	url := "http://api.exchangeratesapi.io/v1/latest?access_key=db3cf960ce2f00c47486c7513971fee1&base=EUR"

	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("x-rapidapi-host", "currency-converter5.p.rapidapi.com")
	//req.Header.Add("x-rapidapi-key", "a3b3a8f12emshfa9afa2d1552d23p13abfajsn4681908f2c15")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	target := models.Target{}
	err := json.Unmarshal([]byte(body), &target)
	if err != nil {
		log.Printf("%v", err)
		return models.Target{}, err
	}

	return target, nil
}
