package gotenberg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type GotenbergClient struct {
	config GotenbergConfig
}

type GotenbergConfig struct {
	ServiceOrigin      string
	WebhookMethod      string
	WebhookErrorUrl    string
	WebhookErrorMethod string
}

func NewGotenbergClient(cnf GotenbergConfig) *GotenbergClient {
	return &GotenbergClient{config: cnf}
}

func (c GotenbergClient) ConvertURLWithWebhook(ctx context.Context, htmlUrl string, webhookUrl string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormField("url")
	io.WriteString(part, htmlUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/forms/chromium/convert/url", c.config.ServiceOrigin), body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Gotenberg-Webhook-Url", webhookUrl)
	req.Header.Add("Gotenberg-Webhook-Method", c.config.WebhookMethod)
	req.Header.Add("Gotenberg-Webhook-Error-Url", c.config.WebhookErrorUrl)
	req.Header.Add("Gotenberg-Webhook-Error-Method", c.config.WebhookErrorMethod)

	client := &http.Client{}

	_, err := client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
