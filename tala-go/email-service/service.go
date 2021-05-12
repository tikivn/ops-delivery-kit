// Copyright (c) 2020 TNSL Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package email_service

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultMailgunServer string = mailgun.APIBase
)

var (
	sendMailTimeout = 30
)

type Service interface {
	SendEmailWithMailgunService(params *MailgunMailParams) error
}

func NewService(
	emailHost string,
	mailgunApiBase string,
	mailgunApiKey string,
	mailgunDomain string,
	storage StorageService,
) (Service, error) {
	mg := mailgun.NewMailgun(mailgunDomain, mailgunApiKey)

	mg.SetAPIBase(mailgunApiBase)

	return &service{
		client:         resty.New().SetHostURL(emailHost),
		mailgunApiBase: mailgunApiBase,
		emailHost:      emailHost,
		mg:             mg,
		storage:        storage,
	}, nil
}

type (
	service struct {
		mailgunApiBase string
		emailHost      string
		client         *resty.Client
		mg             mailgun.Mailgun
		storage        StorageService
	}

	MailgunMailParams struct {
		Sender            string
		Recipients        []string
		CCs               []string
		BCCs              []string
		MailRegexValidate string
		Content           string
		Subject           string
		Files             []*File
	}

	File struct {
		Name string
		Body io.ReadCloser
	}
)

//SendEmailWithMailgunService send email to mailgun api, with email logic (transform data, make title,...)
func (s *service) SendEmailWithMailgunService(params *MailgunMailParams) error {
	if params == nil {
		return errors.New("param is nil")
	}

	if params.MailRegexValidate != "" {
		validate := regexp.MustCompile(params.MailRegexValidate)

		for _, r := range params.Recipients {
			if validate.FindString(r) == "" {
				return errors.New("invalid recipient address")
			}
		}

		// validate all cc, bcc

		for _, cc := range params.CCs {
			if validate.FindString(cc) == "" {
				return errors.New("invalid cc address")
			}
		}

		for _, bcc := range params.BCCs {
			if validate.FindString(bcc) == "" {
				return errors.New("invalid bcc address")
			}
		}
	}

	msg := s.mg.NewMessage(params.Sender, params.Subject, "")

	msg.SetHtml(params.Content)

	for _, r := range params.Recipients {
		msg.AddRecipient(r)
	}

	for _, cc := range params.CCs {
		msg.AddCC(cc)
	}

	for _, bcc := range params.BCCs {
		msg.AddBCC(bcc)
	}

	for _, file := range params.Files {
		msg.AddReaderAttachment(file.Name, file.Body)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sendMailTimeout))
	defer cancel()

	_, id, err := s.mg.Send(ctx, msg)

	if err != nil {
		return errors.Wrap(err, "unknown error when sending to mail gun")
	}

	fmt.Printf("ID: %s\n", id)

	return nil

}
