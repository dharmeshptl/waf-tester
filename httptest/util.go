package httptest

import (
	"crypto/rand"
	"fmt"
	"log"
)

// https://yourbasic.org/golang/generate-uuid-guid/
func genUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func intInSlice(n int, slice []int) bool {
	for _, m := range slice {
		if n == m {
			return true
		}
	}
	return false
}

func stringInSlice(s string, slice []string) bool {
	for _, r := range slice {
		if s == r {
			return true
		}
	}
	return false
}

func (t *Test) addCustomHeader() {
	t.ID = genUUID()
	if t.Headers == nil {
		t.Headers = make(map[string]string)
	}
	t.Headers["waf-tester-id"] = t.ID
}

func (t *Test) fixHostHeader(host string) {
	if host == "localhost" || host == "127.0.0.1" {
		return
	}

	if t.Headers["Host"] == "localhost" {
		t.Headers["Host"] = host
	}
}
