package bar

import (
	"fmt"
	"log"
	"os"
)

const (
	BAR_HOST     = `E2E_BAR_HOST`
	LOCAL_HOST   = `localhost`
	SERVICE_PORT = 8083
)

var (
	BarHost string
)

func init() {
	if host, ok := os.LookupEnv(BAR_HOST); ok {
		BarHost = fmt.Sprintf("%s:%d", host, SERVICE_PORT)
	} else {
		BarHost = fmt.Sprintf("%s:%d", LOCAL_HOST, SERVICE_PORT)
	}
	log.Printf("bar hostname: %s", BarHost)
}
