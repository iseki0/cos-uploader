package main

import (
	"os"
	"strings"
)

var (
	CosBucketUrl  string
	CosServiceUrl string
	CosSecretId   string
	CosSecretKey  string
)

func init() {
	CosBucketUrl = os.Getenv("COS_BUCKET_URL")
	CosServiceUrl = os.Getenv("COS_SERVICE_URL")
	CosSecretId = os.Getenv("COS_SECRET_ID")
	CosSecretKey = os.Getenv("COS_SECRET_KEY")
	s := strings.Split(os.Getenv("COS_STR"), ":")
	if CosBucketUrl == "" && len(s) > 0 {
		CosBucketUrl = s[0]
	}
	if CosServiceUrl == "" && len(s) > 1 {
		CosServiceUrl = s[1]
	}
	if CosSecretId == "" && len(s) > 2 {
		CosSecretId = s[2]
	}
	if CosSecretKey == "" && len(s) > 3 {
		CosSecretKey = s[3]
	}
}
