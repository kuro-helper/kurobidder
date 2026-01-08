package letao

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	kurohelperproxy "github.com/kuro-helper/kurohelper-proxy"
)

func sendGetRequest(url string) (io.Reader, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36 Edg/140.0.0.0")

	dialer, err := kurohelperproxy.GetProxyDialer(os.Getenv("PROXY_PRIVATE_IP"), nil, os.Getenv("PROXY_PRIVATE_PORT"))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("kurobidder: server return a error code %d", resp.StatusCode)
	}

	// Read all content into memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}
