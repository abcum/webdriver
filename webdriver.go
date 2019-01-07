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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Driver represents a WebDriver instance
type Driver struct {
	url string
	exe string
	cmd *exec.Cmd
}

// NewDriver creates a new driver instance.
func NewDriver(url, exe string) *Driver {
	return &Driver{url: url, exe: exe}
}

func (w *Driver) Start() error {

	if w.cmd != nil {
		return nil
	}

	w.cmd = exec.Command(w.exe)

	return w.cmd.Start()

}

func (w *Driver) Stop() error {

	if w.cmd == nil {
		return nil
	}

	w.cmd.Process.Signal(os.Interrupt)

	w.cmd = nil

	return nil

}

// Session creates a new WebDriver session, launching a new remote browser instance.
func (w *Driver) Session(desired, required map[string]interface{}) (*Session, error) {

	if desired == nil {
		desired = make(map[string]interface{})
	}

	if required == nil {
		required = make(map[string]interface{})
	}

	opt := map[string]interface{}{
		"desiredCapabilities":  desired,
		"requiredCapabilities": required,
	}

	id, res, err := w.post("/session", opt)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}

	return &Session{wd: w, ID: id, CB: out}, nil

}

// Sessions returns all of the currently active browser sessions.
func (w *Driver) Sessions() ([]*Session, error) {

	_, res, err := w.get("/sessions")
	if err != nil {
		return nil, err
	}

	var out []*Session
	err = json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}

	for i := range out {
		out[i].wd = w
	}

	return out, nil

}

func (w *Driver) del(url string, pms ...interface{}) (id string, out []byte, err error) {

	var obj response

	uri := w.url + fmt.Sprintf(url, pms...)

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return "", nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-charset", "utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(buf, &obj)

	if res.StatusCode == 200 && err != nil {
		return "", nil, errors.New("error: response must be a JSON object")
	}

	if res.StatusCode >= 400 || obj.Status != 0 {
		return "", nil, oops(res.StatusCode, &obj)
	}

	sid := string(bytes.Trim(obj.SessionId, "{}\""))

	return sid, []byte(obj.Value), nil

}

func (w *Driver) get(url string, pms ...interface{}) (id string, out []byte, err error) {

	var obj response

	uri := w.url + fmt.Sprintf(url, pms...)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-charset", "utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(buf, &obj)

	if res.StatusCode == 200 && err != nil {
		return "", nil, errors.New("error: response must be a JSON object")
	}

	if res.StatusCode >= 400 || obj.Status != 0 {
		return "", nil, oops(res.StatusCode, &obj)
	}

	sid := string(bytes.Trim(obj.SessionId, "{}\""))

	return sid, []byte(obj.Value), nil

}

func (w *Driver) post(url string, opt map[string]interface{}, pms ...interface{}) (id string, out []byte, err error) {

	var obj response

	if opt == nil {
		opt = make(map[string]interface{})
	}

	jsn, err := json.Marshal(opt)
	if err != nil {
		return "", nil, err
	}

	uri := w.url + fmt.Sprintf(url, pms...)

	req, err := http.NewRequest("POST", uri, bytes.NewReader(jsn))
	if err != nil {
		return "", nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-charset", "utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(buf, &obj)

	if res.StatusCode == 200 && err != nil {
		return "", nil, errors.New("error: response must be a JSON object")
	}

	if res.StatusCode >= 400 || obj.Status != 0 {
		return "", nil, oops(res.StatusCode, &obj)
	}

	sid := string(bytes.Trim(obj.SessionId, "{}\""))

	return sid, []byte(obj.Value), nil

}
