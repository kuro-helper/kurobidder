package letao

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sendGetRequest(url string) (io.Reader, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36 Edg/140.0.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

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
