/*
 * @Author: yujiajie
 * @Date: 2024-03-18 17:43:23
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-21 14:38:21
 * @FilePath: /gateway/core/stat/remotewriter.go
 * @Description:
 */
package stat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gateway/server/rest/header"
	"net/http"
	"time"
)

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
	resp, err := client.Post(rw.endpoint, header.JsonContentType, bytes.NewReader(bs))
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
