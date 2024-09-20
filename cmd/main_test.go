package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/test-go/testify/assert"
)

func TestSend(t *testing.T) {
	t.SkipNow()
	urls := []string{
		"http://localhost:8060/dingtalk/webhook1/send",
		"http://localhost:8060/wechat/webhook2/send",
		"http://localhost:8060/lark/webhook3/send",
	}
	cli := http.DefaultClient
	reqData, err := os.ReadFile("./test_data.json")
	assert.NoError(t, err)
	fmt.Println("req data:", string(reqData))
	for _, url := range urls {
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqData))
		assert.NoError(t, err)
		resp, err := cli.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		dd, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		fmt.Printf("status: %d, resp:%s\n", resp.StatusCode, string(dd))
	}
}
