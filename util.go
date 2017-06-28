// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webdriver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

type pos struct {
	X int `json:"x" console:"x"`
	Y int `json:"y" console:"y"`
}

type geo struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
	Alt float64 `json:"altitude"`
}

type size struct {
	Width  int `json:"width" console:"width"`
	Height int `json:"height" console:"height"`
}

type response struct {
	SessionId json.RawMessage `json:"sessionId"`
	Status    int             `json:"status"`
	Value     json.RawMessage `json:"value"`
}

func oops(c int, obj *response) error {
	switch c {
	case 400:
		return errors.New("400: Missing Command Parameters")
	case 404:
		return errors.New("404: Unknown command/Resource Not Found")
	case 405:
		return errors.New("405: Invalid Command Method")
	case 500:
		return errors.New("500: Failed Command")
	case 501:
		return errors.New("501: Unimplemented Command")
	default:
		return errors.New("Unknown error")
	}
}

func wait(port int, timeout time.Duration) error {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	now := time.Now()
	for {
		if conn, err := net.Dial("tcp", address); err == nil {
			if err = conn.Close(); err != nil {
				return err
			}
			break
		}
		if time.Since(now) > timeout {
			return errors.New("start failed: timeout expired")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
