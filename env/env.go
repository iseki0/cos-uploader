package env

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"regexp"
)

var (
	CosBucketUrl  *url.URL
	CosServiceUrl *url.URL
	CosSecretId   string
	CosSecretKey  string
	RelayListen   = os.Getenv("RELAY_LISTEN")
	RelayKey      = os.Getenv("RELAY_KEY")
	RelayBase     = os.Getenv("RELAY_BASE")
)

//go:generate stringer -type envError -linecomment
type envError int

const (
	_                     envError = iota
	ErrCosBase64EnvNotSet          // env: COS_BASE64 not set
	ErrCosBase64ParseFail          // env: COS_BASE64 parse fail
	ErrBucketURLNotSet             // env: bucket URL not set
	ErrServiceURLNotSet            // env: service URL not set
)

func (i envError) Error() string {
	return i.String()
}

func InitEnv() error {
	ev := os.Getenv("COS_BASE64")
	if ev == "" {
		return ErrCosBase64EnvNotSet
	}
	raw, e := base64.StdEncoding.DecodeString(ev)
	if e != nil {
		return ErrCosBase64ParseFail
	}
	ev = string(raw)
	pattern := regexp.MustCompile(`([^\s=]+)=(\S*)`)
	m := pattern.FindAllStringSubmatch(ev, -1)
	if len(m) == 0 {
		return ErrCosBase64ParseFail
	}
	for _, it := range m {
		k := it[1]
		v := it[2]
		var e error
		switch k {
		case "bucket":
			CosBucketUrl, e = url.Parse(v)
			if e != nil {
				return fmt.Errorf("parse bucket url: %w", e)
			}
		case "service":
			CosServiceUrl, e = url.Parse(v)
			if e != nil {
				return fmt.Errorf("parse service url: %w", e)
			}
		case "sid":
			CosSecretId = v
		case "sk":
			CosSecretKey = v
		}
	}
	if CosServiceUrl == nil {
		return ErrServiceURLNotSet
	}
	if CosBucketUrl == nil {
		return ErrBucketURLNotSet
	}
	return nil
}
