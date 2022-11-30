package uploader

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iseki0/cos-uploader/env"
	"io"
	"os"
	"path/filepath"
)

//go:generate stringer -type relayError -linecomment
type relayError int

const (
	_                   relayError = iota
	ErrRelayKeyNotSet              // relay: token not set
	ErrRelayTokenWrong             // relay: token wrong
	ErrRelayBadFilename            // relay: bad filename
)

func (i relayError) Error() string {
	return i.String()
}

func Relay() error {
	if env.RelayKey == "" {
		return ErrRelayKeyNotSet
	}
	g := gin.New()

	g.POST("/upload/*name", func(c *gin.Context) {
		if token := c.GetHeader("token"); token != env.RelayKey {
			panic(ErrRelayTokenWrong)
		}
		fn := c.Param("name")
		if fn == "" {
			panic(ErrRelayBadFilename)
			return
		}
		f := m1(os.CreateTemp("", "relay-*"))
		tn := filepath.Join(os.TempDir(), m1(f.Stat()).Name())

		defer func() {
			os.Remove(tn)
		}()
		_ = m1(io.Copy(f, c.Request.Body))
		fmt.Println(tn)
		_ = f.Close()
		m0(UploadFile(tn, env.RelayBase))
		c.Status(201)
	})
	return g.Run(env.RelayListen)
}
func m1[T any](t T, e error) T {
	if e != nil {
		panic(e)
	}
	return t
}

func m0(e error) {
	if e != nil {
		panic(e)
	}
}
