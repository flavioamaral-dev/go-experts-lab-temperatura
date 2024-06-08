package services

import (
	"flavioamaral-dev/go-experts-lab-temperatura/adapters"
	"flavioamaral-dev/go-experts-lab-temperatura/dto"
	"flavioamaral-dev/go-experts-lab-temperatura/errors"
	"net/url"
)

func SearchWeather(zipCode string) (*dto.TemperaturaResponse, error) {
	resViaCep, err := adapters.SearchZipCode(zipCode)
	if resViaCep.Cep == "" {
		return nil, errors.NotFoundZipCode
	}

	if err != nil {
		return nil, errors.UnableToRetrieveZipCode
	}

	city := url.QueryEscape(resViaCep.Locale)
	resWeather, err := adapters.GetWeather(city)
	if err != nil {
		return nil, errors.UnableToRetrieveWeather
	}

	return &dto.TemperaturaResponse{
		TempC: resWeather.Current.TempC,
		TempK: getTemperatureKelvin(resWeather.Current.TempC),
		TempF: getTemperatureFahrenheit(resWeather.Current.TempC),
	}, nil
}

func getTemperatureFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func getTemperatureKelvin(celsius float64) float64 {
	return celsius + 273
}
