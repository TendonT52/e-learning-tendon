package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestEndToEnd(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Db Suite")
}

func getCookie(res *http.Response) map[string]string {
	cookies := make(map[string]string)
	for _, cookie := range res.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	return cookies
}

func getJSON(res *http.Response) map[string]interface{} {
	var j map[string]interface{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body, &j)
	if err != nil {
		return nil
	}
	return j
}
