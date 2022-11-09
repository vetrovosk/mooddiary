// вот тут должен быть маленький мейн
package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {

	// err := errors.New("что-то пошло не так")
	// log.Fatalf("случилась ошибка, error: %s", err)

	logger, _ := zap.NewProduction()

	url := "http://blabla.ru"
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Debug("какой-то тупой дебаг лог")
	sugar.Warn("a это предупреждающий лог!")
}

// // в файле
// cfg:
// - city: Moscow
//   phone: 8495-123-234
// - city: Sochi
//   phone: 8263-74234-8989
// - city: London
//   phone: 8324-34-34
// - city: Varna
//   phone: 3742834-34-34

// // после того, как прочитаем
// var map[string]string{
// 	"Moscow": "8495-123-234",
// 	"Sochi": "8263-74234-8989",
// 	"London":  "8324-34-34",
// 	"Varna": "3742834-34-34"
// }
