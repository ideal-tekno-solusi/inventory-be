package utils

import (
	"bytes"
	"io"
	"net/http"
)

func SendHttpRequest(method, url string, body []byte) (int, *io.ReadCloser, error) {
	if method == "" {
		method = "GET"
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	//TODO: keknya bakal perlu penjagaan lebih deh ini
	defer res.Body.Close()

	return res.StatusCode, &res.Body, nil
}
