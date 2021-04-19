package log

import (
	"bytes"
	"fmt"
	"github.com/sonereker/distributed-go-sandbox/basic/registry"
	stlog "log"
	"net/http"
)

type clientLogger struct {
	url string
}

//SetClientLogger set logger options for the provided service
func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer(data)
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with status %v", res.StatusCode)
	}
	return len(data), nil
}
