package services

type Horoscope interface {
	GetDailyHoroscope(sign string) (*HoroscopeResponse, error)
	Close() error
}

type HoroscopeResponse struct {
	Date          string `json:"date"`
	HoroscopeData string `json:"horoscope_data"`
	ZodiacSign    string `json:"zodiac_sign"`
}
