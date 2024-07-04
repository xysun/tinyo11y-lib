package tinyo11ylib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TinyLogger struct {
	ApiKey  string
	LogHost string
	LogPort string
}

func (logger *TinyLogger) output(level string, msg string) {
	payload := map[string]string{"message": msg, "level": level}

	go func() {
		requestJson, _ := json.Marshal(payload)
		http.Post(fmt.Sprintf("http://%s:%s/v1/log/%s", logger.LogHost, logger.LogPort, logger.ApiKey),
			"application/json", bytes.NewBuffer(requestJson))
	}()
}

func (logger *TinyLogger) Info(msg string) {
	logger.output("info", msg)
}
