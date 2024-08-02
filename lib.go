package tinyo11ylib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TinyLogger struct {
	ApiKey string
	Host   string
	Port   string // TODO: remove when launch
}

func (logger *TinyLogger) output(level string, msg string) {
	payload := map[string]string{"message": msg, "level": level}

	go func() {
		requestJson, _ := json.Marshal(payload)
		// TODO: change to https when launch
		http.Post(fmt.Sprintf("http://%s:%s/v1/log/%s", logger.Host, logger.Port, logger.ApiKey),
			"application/json", bytes.NewBuffer(requestJson))
	}()
}

func (logger *TinyLogger) Info(msg string) {
	logger.output("info", msg)
}

func (logger *TinyLogger) Error(msg string, er error) {
	// by default golang does not show stacktrace....
	logger.output("error", fmt.Sprintf("%s: %s", msg, er.Error()))
}

func (logger *TinyLogger) Metric(metricName string, metricValue float64) {
	payload := map[string]string{"metric_name": metricName, "metric_value": fmt.Sprintf("%.6f", metricValue)}

	go func() {
		requestJson, _ := json.Marshal(payload)
		// TODO: change to https when launch
		http.Post(fmt.Sprintf("http://%s:%s/v1/metric/%s", logger.Host, logger.Port, logger.ApiKey),
			"application/json", bytes.NewBuffer(requestJson))

	}()
}
