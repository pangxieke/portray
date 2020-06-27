package middle

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type response struct {
	http.ResponseWriter
	Status int
	Body   []byte
}

func (r *response) Write(data []byte) (int, error) {
	r.Body = data
	return r.ResponseWriter.Write(data)
}

func (r *response) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// Log middleware
// record request and response info in std error
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now().UnixNano()

		m := map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"rawQuery": r.URL.RawQuery,
		}

		headers := []string{}
		for k, v := range r.Header {
			headers = append(headers, fmt.Sprintf("%v: %v", k, v[0]))
		}
		m["req_header"] = strings.Join(headers, ", ")

		body, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		m["req_body"] = body

		// Call the next handler in the chain.
		res := response{ResponseWriter: w}
		next.ServeHTTP(&res, r)

		elapse := float64(time.Now().UnixNano()-begin) / 1000000.0
		headers = []string{}
		for k, v := range res.Header() {
			headers = append(headers, fmt.Sprintf("%v: %v", k, v[0]))
		}
		m["resp_body"] = string(res.Body)
		m["resp_headers"] = strings.Join(headers, ", ")
		m["resp_status"] = res.Status
		m["elapse"] = elapse

		logrus.Info(m)
	})
}
