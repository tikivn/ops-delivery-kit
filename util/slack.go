package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func SendSlack(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	values := map[string]string{"text": message}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://hooks.slack.com/services/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", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		logrus.WithError(err).Infof("Slack has an error!")
	}

	defer resp.Body.Close()
}
