package zed

import (
	"fmt"
	"log"
	"os"
)

const (
	ZED_HOST     = `E2E_ZED_HOST`
	SERVICE_PORT = 8082
	LOCAL_HOST   = `localhost`
)

var (
	ZedHost string
)

func init() {
	if host, ok := os.LookupEnv(ZED_HOST); ok {
		ZedHost = fmt.Sprintf("%s:%d", host, SERVICE_PORT)
	} else {
		ZedHost = fmt.Sprintf("%s:%d", LOCAL_HOST, SERVICE_PORT)
	}
	log.Printf("zed hostname: %s", ZedHost)
}
