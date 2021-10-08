package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Reporter interface {
	Alert(message string)
}

type reporter struct {
	httpDoer      *http.Client
	webhookURL    string
	mentionGroups []string
	mentionUsers  []string
}

type reporterRequest struct {
	Text string `json:"text"`
}

func (r reporter) Alert(message string) {
	go func() {
		mentions := ""
		for _, group := range r.mentionGroups {
			mentions += fmt.Sprintf("<!%s> ", group)
		}
		for _, user := range r.mentionUsers {
			mentions += fmt.Sprintf("<@%s> ", user)
		}
		if mentions != "" {
			mentions += "\n"
		}

		reporterRequest := reporterRequest{
			Text: fmt.Sprintf("%s%s", mentions, message),
		}

		body, err := json.Marshal(reporterRequest)
		if err != nil {
			logrus.Error(err)
			return
		}

		request, err := http.NewRequest("POST", r.webhookURL, bytes.NewBuffer(body))
		if err != nil {
			logrus.Error(err)
			return
		}
		request.Header.Add("Content-Type", "application/json")

		_, err = r.httpDoer.Do(request)
		if err != nil {
			logrus.Error(err)
		}
	}()
}

func NewReporter(webhookURL string, mentionGroups []string, mentionUsers []string) Reporter {
	return reporter{
		httpDoer: &http.Client{
			Timeout: 10 * time.Second,
		},
		webhookURL:    webhookURL,
		mentionGroups: mentionGroups,
		mentionUsers:  mentionUsers,
	}
}
