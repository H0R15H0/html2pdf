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
	WebhookUrl      string
	WebhookErrorUrl string
}

func NewGotenbergClient(origin string, wUrl string, wErrUrl string) *GotenbergClient {
	return &GotenbergClient{Origin: origin, WebhookUrl: wUrl, WebhookErrorUrl: wErrUrl}
}

func (c GotenbergClient) ConvertURLWithExtraHTTP(ctx context.Context, url string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormField("url")
	io.WriteString(part, url)
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
