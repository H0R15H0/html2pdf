package gotenberg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type GotenbergClient struct {
	Origin          string
	WebhookErrorUrl string
}

func NewGotenbergClient(origin string, wErrUrl string) *GotenbergClient {
	return &GotenbergClient{Origin: origin, WebhookErrorUrl: wErrUrl}
}

func (c GotenbergClient) ConvertURLWithWebhook(ctx context.Context, htmlUrl string, webhookUrl string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormField("url")
	io.WriteString(part, htmlUrl)
	writer.Close()

	// TODO: set Gotenberg-Webhook-Extra-Http-Headers to upload generated PDF to s3 bucket.
	// https://github.com/gotenberg/gotenberg/issues/495
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/forms/chromium/convert/url", c.Origin), body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("/app/hoge.pdf")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
