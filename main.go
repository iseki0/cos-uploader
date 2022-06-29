package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

var ErrBadBucketUrl = errors.New("COS_BUCKET_URL invalid")
var ErrBadServiceUrl = errors.New("COS_SERVICE_URL invalid")

var bucketFileName string
var localFileName string

func main() {
	c := &cobra.Command{
		RunE:               rootRunE,
		FParseErrWhitelist: cobra.FParseErrWhitelist{},
		CompletionOptions:  cobra.CompletionOptions{},
		SilenceUsage:       true,
	}

	c.Flags().StringVar(&bucketFileName, "remote", "", "file path in bucket")
	must(c.MarkFlagRequired("remote"))
	c.Flags().StringVar(&localFileName, "local", "", "local file")
	must(c.MarkFlagFilename("local"))
	must(c.MarkFlagRequired("local"))
	if e := c.Execute(); e != nil {
		os.Exit(1)
	}
}

func rootRunE(cmd *cobra.Command, args []string) error {
	if CosSecretId == "" {
		PrintWarn("COS_SECRET_ID is empty")
	}
	if CosSecretKey == "" {
		PrintWarn("COS_SECRET_KEY is empty")
	}
	var bucketUrl, serviceUrl *url.URL
	var e error
	bucketUrl, e = url.Parse(CosBucketUrl)
	if e != nil {
		return ErrBadBucketUrl
	}
	if CosServiceUrl != "" {
		serviceUrl, e = url.Parse(CosServiceUrl)
		if e != nil {
			return ErrBadServiceUrl
		}
	}
	baseUrl := &cos.BaseURL{
		BucketURL:  bucketUrl,
		ServiceURL: serviceUrl,
	}
	client := cos.NewClient(baseUrl, &http.Client{Transport: &cos.AuthorizationTransport{
		SecretID:  CosSecretId,
		SecretKey: CosSecretKey,
	}})
	PrintInfo("Uploading...", localFileName, "->", bucketFileName)
	resp, e := client.Object.PutFromFile(context.Background(), bucketFileName, localFileName, nil)
	if e != nil {
		PrintError("Upload fail", e.Error())
		return e
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		PrintError(fmt.Sprintf("HTTP status: %d - %s", resp.StatusCode, resp.Status))
	} else {
		PrintInfo(fmt.Sprintf("HTTP status: %d - %s", resp.StatusCode, resp.Status))
	}
	return nil
}
