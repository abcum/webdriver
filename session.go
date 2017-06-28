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
	"encoding/base64"
	"encoding/json"
	"io"
)

// Session represents a web page session.
type Session struct {
	ID string                 `json:"id"`
	CB map[string]interface{} `json:"capabilities"`
	wd *Driver
}

// FindStrategy specifies which strategy to use when searching for elements.
type FindStrategy string

const (
	// FindByClass finds eelements whose class name contains the search value.
	FindByClass FindStrategy = "class name"
	// FindByCss finds elements matching a CSS selector.
	FindByCss = "css selector"
	// FindById finds elements whose ID attribute matches the search value.
	FindById = "id"
	// FindByName an element whose NAME attribute matches the search value.
	FindByName = "name"
	// FindByLinkText finds links whose visible text matches the search value.
	FindByLinkText = "link text"
	// FindByPartialLinkText finds links whose text partially matches the search value.
	FindByPartialLinkText = "partial link text"
	// FindByTagName finds elements whose tag name matches the search value.
	FindByTagName = "tag name"
	// XPath finds elements matching an XPath expression.
	XPath = "xpath"
)

// Window gets the current active window.
func (s *Session) Window() *Window {
	return &Window{ws: s, ID: "current"}
}

// Url gets the url of the current page.
func (s *Session) Url() (string, error) {
	_, res, err := s.wd.get("/session/%s/url", s.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Title gets the title of the current page.
func (s *Session) Title() (string, error) {
	_, res, err := s.wd.get("/session/%s/title", s.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Source gets the source code of the current page.
func (s *Session) Source() (string, error) {
	_, res, err := s.wd.get("/session/%s/source", s.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Delete deletes the current session, freeing resources.
func (s *Session) Delete() error {
	_, _, err := s.wd.del("/session/%s", s.ID)
	return err
}

// Timeouts enables specifying custom timeouts for the current session.
func (s *Session) Timeouts(what string, ms int) error {
	opt := map[string]interface{}{"type": what, "ms": ms}
	_, _, err := s.wd.post("/session/%s/timeouts", opt, s.ID)
	return err
}

// TimeoutsAsyncScript specifies a custom timeout when running asynchronous scripts.
func (s *Session) TimeoutsAsyncScript(ms int) error {
	opt := map[string]interface{}{"ms": ms}
	_, _, err := s.wd.post("/session/%s/timeouts/async_script", opt, s.ID)
	return err
}

// TimeoutsImplicitWait specifies a custom timeout when searching for elements on the page.
func (s *Session) TimeoutsImplicitWait(ms int) error {
	opt := map[string]interface{}{"ms": ms}
	_, _, err := s.wd.post("/session/%s/timeouts/implicit_wait", opt, s.ID)
	return err
}

// Load loads a new url in the current page.
func (s *Session) Load(url string) error {
	opt := map[string]interface{}{"url": url}
	_, _, err := s.wd.post("/session/%s/url", opt, s.ID)
	return err
}

// Back causes the browser to traverse one step backwards in the session history.
func (s *Session) Back() error {
	_, _, err := s.wd.post("/session/%s/back", nil, s.ID)
	return err
}

// Forward causes the browser to traverse one step forwards in the session history.
func (s *Session) Forward() error {
	_, _, err := s.wd.post("/session/%s/forward", nil, s.ID)
	return err
}

// Refresh causes the browser to reload the curretn page.
func (s *Session) Refresh() error {
	_, _, err := s.wd.post("/session/%s/refresh", nil, s.ID)
	return err
}

// ExecuteSync executes a JavaScript script synchronously in the current page.
func (s *Session) ExecuteSync(script string, args []interface{}) ([]byte, error) {
	opt := map[string]interface{}{"script": script, "args": args}
	_, res, err := s.wd.post("/session/%s/execute/sync", opt, s.ID)
	return res, err
}

// ExecuteAsync executes a JavaScript script asynchronously in the current page.
func (s *Session) ExecuteAsync(script string, args []interface{}) ([]byte, error) {
	opt := map[string]interface{}{"script": script, "args": args}
	_, res, err := s.wd.post("/session/%s/execute/async", opt, s.ID)
	return res, err
}

// Screenshot takes a screenshot of the full browser viewport.
func (s *Session) Screenshot() (io.Reader, error) {
	_, res, err := s.wd.get("/session/%s/screenshot", s.ID)
	if err != nil {
		return nil, err
	}
	return base64.NewDecoder(base64.StdEncoding, bytes.NewReader(res[1:len(res)-1])), nil
}

// Active returns the currently active element within the current page.
func (s *Session) Active() (*Element, error) {
	_, res, err := s.wd.post("/session/%s/element/active", nil, s.ID)
	if err != nil {
		return nil, err
	}
	var out Element
	err = json.Unmarshal(res, &out)
	out.ws = s
	return &out, err
}

// Element searches for a single element from within the current page.
func (s *Session) Element(using FindStrategy, value string) (*Element, error) {
	opt := map[string]interface{}{"using": using, "value": value}
	_, res, err := s.wd.post("/session/%s/element", opt, s.ID)
	if err != nil {
		return nil, err
	}
	var out Element
	err = json.Unmarshal(res, &out)
	out.ws = s
	return &out, err
}

// Elements searches for multiple elements from within the current page.
func (s *Session) Elements(using FindStrategy, value string) ([]*Element, error) {
	opt := map[string]interface{}{"using": using, "value": value}
	_, res, err := s.wd.post("/session/%s/elements", opt, s.ID)
	if err != nil {
		return nil, err
	}
	var out []*Element
	err = json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}
	for i := range out {
		out[i].ws = s
	}
	return out, err
}

//
//
//
//
//

// AlertText returns the text of the currently displayed dialog window.
func (s *Session) AlertText() (string, error) {
	_, res, err := s.wd.get("/session/%s/alert_text", s.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// RespondAlert sends keystrokes to the currently displayed dialog window.
func (s *Session) RespondAlert(text string) error {
	opt := map[string]interface{}{"text": text}
	_, _, err := s.wd.post("/session/%s/alert_text", opt, s.ID)
	return err
}

// AcceptAlert accepts the currently displayed alert dialog window.
func (s *Session) AcceptAlert() error {
	_, _, err := s.wd.post("/session/%s/accept_alert", nil, s.ID)
	return err
}

// DismissAlert cancels the currently displayed alert dialog window.
func (s *Session) DismissAlert() error {
	_, _, err := s.wd.post("/session/%s/dismiss_alert", nil, s.ID)
	return err
}

// Move moves the mouse by an offset from the specified element.
func (s *Session) Move(e *Element, x, y int) error {
	opt := map[string]interface{}{"element": e.ID, "xoffset": x, "yoffset": y}
	_, _, err := s.wd.post("/session/%s/moveto", opt, s.ID)
	return err
}

// Click clicks the specified mouse button at the current coordinates.
func (s *Session) Click(button int) error {
	opt := map[string]interface{}{"button": button}
	_, _, err := s.wd.post("/session/%s/click", opt, s.ID)
	return err
}

// ButtonDown clicks and holds the specified mouse button at the current coordinates.
func (s *Session) ButtonDown(button int) error {
	opt := map[string]interface{}{"button": button}
	_, _, err := s.wd.post("/session/%s/buttondown", opt, s.ID)
	return err
}

// ButtonUp releases the specified mouse button at the current coordinates.
func (s *Session) ButtonUp(button int) error {
	opt := map[string]interface{}{"button": button}
	_, _, err := s.wd.post("/session/%s/buttonup", opt, s.ID)
	return err
}

// DoubleClick double clicks the left mouse button at the current coordinates.
func (s *Session) DoubleClick() error {
	_, _, err := s.wd.post("/session/%s/doubleclick", nil, s.ID)
	return err
}

// TouchClick executes a single tap at the current coordinates.
func (s *Session) TouchClick(e *Element) error {
	opt := map[string]interface{}{"element": e.ID}
	_, _, err := s.wd.post("/session/%s/touch/click", opt, s.ID)
	return err
}

// TouchDown presses a finger down on the page at the current coordinates.
func (s *Session) TouchDown(x, y int) error {
	opt := map[string]interface{}{"x": x, "y": y}
	_, _, err := s.wd.post("/session/%s/touch/down", opt, s.ID)
	return err
}

// TouchUp lifts a finger up from the page at the current coordinates.
func (s *Session) TouchUp(x, y int) error {
	opt := map[string]interface{}{"x": x, "y": y}
	_, _, err := s.wd.post("/session/%s/touch/up", opt, s.ID)
	return err
}

// TouchMove moves the currently pressed finsed on the page.
func (s *Session) TouchMove(x, y int) error {
	opt := map[string]interface{}{"x": x, "y": y}
	_, _, err := s.wd.post("/session/%s/touch/move", opt, s.ID)
	return err
}

// TouchScroll scrolls on the page using finger based motion events.
func (s *Session) TouchScroll(e *Element, x, y int) error {
	opt := map[string]interface{}{"element": e.ID, "xoffset": x, "yoffset": y}
	_, _, err := s.wd.post("/session/%s/touch/scroll", opt, s.ID)
	return err
}

// TouchDoubleClick executes a double click on the specified element.
func (s *Session) TouchDoubleClick(e *Element) error {
	opt := map[string]interface{}{"element": e.ID}
	_, _, err := s.wd.post("/session/%s/touch/doubleclick", opt, s.ID)
	return err
}

// TouchLongClick executes a long tap on the specified element.
func (s *Session) TouchLongClick(e *Element) error {
	opt := map[string]interface{}{"element": e.ID}
	_, _, err := s.wd.post("/session/%s/touch/longclick", opt, s.ID)
	return err
}

// TouchFlick flicks a finger on the screen starting at the specified element.
func (s *Session) TouchFlick(e *Element, x, y, speed int) error {
	opt := map[string]interface{}{"element": e.ID, "xoffset": x, "yoffset": y, "speed": speed}
	_, _, err := s.wd.post("/session/%s/touch/flick", opt, s.ID)
	return err
}

// TouchFlickAnywhere flicks a finger on the screen starting anywhere.
func (s *Session) TouchFlickAnywhere(xspeed, yspeed int) error {
	opt := map[string]interface{}{"xspeed": xspeed, "yspeed": yspeed}
	_, _, err := s.wd.post("/session/%s/touch/flick", opt, s.ID)
	return err
}

// Cookie returns a new cookie which can be set within the page.
func (s *Session) Cookie(name string) *Cookie {
	return &Cookie{Name: name, ws: s}
}

// Cookies returns all of the cookies visible to the current page.
func (s *Session) Cookies() ([]Cookie, error) {
	_, res, err := s.wd.get("/session/%s/cookie", s.ID)
	if err != nil {
		return nil, err
	}
	var out []Cookie
	err = json.Unmarshal(res, &out)
	return out, err
}

// CookiesClear removes all cookies visible to the current page.
func (s *Session) CookiesClear() error {
	_, _, err := s.wd.del("/session/%s/cookie", s.ID)
	return err
}

// LocalStorageSize returns the current localStorage content size.
func (s *Session) LocalStorageSize() (int, error) {
	_, res, err := s.wd.get("/session/%s/local_storage/size", s.ID)
	if err != nil {
		return -1, err
	}
	var out int
	err = json.Unmarshal(res, &out)
	return out, err
}

// LocalStorageClear clears the localStorage of the current page.
func (s *Session) LocalStorageClear() error {
	_, _, err := s.wd.del("/session/%s/local_storage", s.ID)
	return err
}

// LocalStorageKeys returns all of the localStorage keys for the current page.
func (s *Session) LocalStorageKeys() ([]string, error) {
	_, res, err := s.wd.get("/session/%s/local_storage", s.ID)
	if err != nil {
		return nil, err
	}
	var out []string
	err = json.Unmarshal(res, &out)
	return out, err
}

// LocalStorageGetKey gets the specified localStorage key on the current page.
func (s *Session) LocalStorageGetKey(key string) (string, error) {
	_, res, err := s.wd.get("/session/%s/local_storage/key/%s", s.ID, key)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// LocalStorageDelKey deletes the specified key from localStorage on the current page.
func (s *Session) LocalStorageDelKey(key string) error {
	_, _, err := s.wd.del("/session/%s/local_storage/key/%s", s.ID, key)
	return err
}

// LocalStorageSetKey sets the specified key in localStorage on the current page.
func (s *Session) LocalStorageSetKey(key, value string) error {
	opt := map[string]interface{}{"key": key, "value": value}
	_, _, err := s.wd.post("/session/%s/local_storage", opt, s.ID)
	return err
}

// SessionStorageSize returns the current sessionStorage content size.
func (s *Session) SessionStorageSize() (int, error) {
	_, res, err := s.wd.get("/session/%s/session_storage/size", s.ID)
	if err != nil {
		return -1, err
	}
	var out int
	err = json.Unmarshal(res, &out)
	return out, err
}

// SessionStorageClear clears the sessionStorage of the current page.
func (s *Session) SessionStorageClear() error {
	_, _, err := s.wd.del("/session/%s/session_storage", s.ID)
	return err
}

// SessionStorageKeys returns all of the sessionStorage keys for the current page.
func (s *Session) SessionStorageKeys() ([]string, error) {
	_, res, err := s.wd.get("/session/%s/session_storage", s.ID)
	if err != nil {
		return nil, err
	}
	var out []string
	err = json.Unmarshal(res, &out)
	return out, err
}

// SessionStorageGetKey gets the specified sessionStorage key on the current page.
func (s *Session) SessionStorageGetKey(key string) (string, error) {
	_, res, err := s.wd.get("/session/%s/session_storage/key/%s", s.ID, key)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// SessionStorageDelKey deletes the specified key from sessionStorage on the current page.
func (s *Session) SessionStorageDelKey(key string) error {
	_, _, err := s.wd.del("/session/%s/session_storage/key/%s", s.ID, key)
	return err
}

// SessionStorageSetKey sets the specified key in sessionStorage on the current page.
func (s *Session) SessionStorageSetKey(key, value string) error {
	opt := map[string]interface{}{"key": key, "value": value}
	_, _, err := s.wd.post("/session/%s/session_storage", opt, s.ID)
	return err
}

// ApplicationCacheStatus returns the current application cache status for the current page.
func (s *Session) ApplicationCacheStatus() (int, error) {
	_, res, err := s.wd.get("/session/%s/application_cache/status", s.ID)
	if err != nil {
		return 0, err
	}
	var out int
	err = json.Unmarshal(res, &out)
	return out, err
}
