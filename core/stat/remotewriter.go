/*
 * @Author: yujiajie
 * @Date: 2024-03-18 17:43:23
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 17:46:58
 * @FilePath: /gateway/core/stat/remotewriter.go
 * @Description:
 */
package stat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const jsonContentType = "application/json; charset=utf-8"

var ErrWriteFailed = errors.New("submit failed")

type RemoteWriter struct {
	endpoint string
}

func NewRemoteWriter(endpoint string) Writer {
	return &RemoteWriter{
		endpoint: endpoint,
	}
}

func (rw *RemoteWriter) Write(report *StatReport) error {
	bs, err := json.Marshal(report)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post(rw.endpoint, jsonContentType, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("write report failed, code: %d, reason: %s\n", resp.StatusCode, resp.Status)
		return ErrWriteFailed
	}
	return nil
}
