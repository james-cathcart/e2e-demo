package foo

import (
	"fmt"
	"log"
	"os"
)

const (
	FOO_HOST     = `E2E_FOO_HOST`
	LOCAL_HOST   = `localhost`
	SERVICE_PORT = 8081
)

var (
	FooHost string
)

func init() {
	if host, ok := os.LookupEnv(FOO_HOST); ok {
		FooHost = fmt.Sprintf("%s:%d", host, SERVICE_PORT)
	} else {
		FooHost = fmt.Sprintf("%s:%d", LOCAL_HOST, SERVICE_PORT)
	}
	log.Printf("foo hostname: %s", FooHost)
}
