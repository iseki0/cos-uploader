package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"regexp"
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
	pattern := regexp.MustCompile("([^\\s=]+)=(\\S*)")
	var m [][]string
	if b := os.Getenv("COS_BASE64"); b != "" {
		r, e := io.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(b))))
		if e != nil {
			panic("decode COS_BASE64 base64 failed")
		}
		m = pattern.FindAllStringSubmatch(string(r), -1)
	} else if b := os.Getenv("COS_STR"); b != "" {
		m = pattern.FindAllStringSubmatch(b, -1)
	}
	for _, it := range m {
		k := it[1]
		v := it[2]
		switch k {
		case "bucket":
			if CosBucketUrl == "" {
				CosBucketUrl = v
			}
		case "service":
			if CosServiceUrl == "" {
				CosServiceUrl = v
			}
		case "sid":
			if CosSecretId == "" {
				CosSecretId = v
			}
		case "sk":
			if CosSecretKey == "" {
				CosSecretKey = v
			}
		}
	}
}
