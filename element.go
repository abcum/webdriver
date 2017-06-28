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
)

// Element represents a web element within a page.
type Element struct {
	ID string `json:"ELEMENT"`
	ws *Session
}

// Size returns the size of the element.
func (e *Element) Size() (*size, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/size", e.ws.ID, e.ID)
	if err != nil {
		return nil, err
	}
	var out size
	err = json.Unmarshal(res, &out)
	return &out, err
}

// Name returns the node name of the element.
func (e *Element) Name() (string, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/name", e.ws.ID, e.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Text returns the visible text for the element.
func (e *Element) Text() (string, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/text", e.ws.ID, e.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Html returns the outer html of the element.
func (e *Element) Html() (string, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/attribute/outerHTML", e.ws.ID, e.ID)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Attr returns the specified attribute value for the element.
func (e *Element) Attr(name string) (string, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/attribute/%s", e.ws.ID, e.ID, name)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Css returns the specified computed css property for the element.
func (e *Element) Css(name string) (string, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/css/%s", e.ws.ID, e.ID, name)
	if err != nil {
		return "", err
	}
	var out string
	err = json.Unmarshal(res, &out)
	return out, err
}

// Location returns the current location coordinates of the element.
func (e *Element) Location() (*pos, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/location", e.ws.ID, e.ID)
	if err != nil {
		return nil, err
	}
	var out pos
	err = json.Unmarshal(res, &out)
	return &out, err
}

// Clear clears the value of the current element (must be a text input box).
func (e *Element) Clear() error {
	_, _, err := e.ws.wd.post("/session/%s/element/%s/clear", nil, e.ws.ID, e.ID)
	return err
}

// Click clicks on the current element.
func (e *Element) Click() error {
	_, _, err := e.ws.wd.post("/session/%s/element/%s/click", nil, e.ws.ID, e.ID)
	return err
}

// Submit submits the current element (must be a form).
func (e *Element) Submit() error {
	_, _, err := e.ws.wd.post("/session/%s/element/%s/submit", nil, e.ws.ID, e.ID)
	return err
}

// Value sends a sequence of key strokes to the element.
func (e *Element) Value(sequence string) error {
	keys := make([]string, len(sequence))
	for i, k := range sequence {
		keys[i] = string(k)
	}
	opt := map[string]interface{}{"value": keys}
	_, _, err := e.ws.wd.post("/session/%s/element/%s/value", opt, e.ws.ID, e.ID)
	return err
}

// Equals returns true if this element is the same as another element.
func (e *Element) Equals(o *Element) (bool, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/equal/%s", e.ws.ID, e.ID, o.ID)
	if err != nil {
		return false, err
	}
	var out bool
	err = json.Unmarshal(res, &out)
	return out, err
}

// Enabled returns whether the current element is enabled or not.
func (e *Element) Enabled() (bool, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/enabled", e.ws.ID, e.ID)
	if err != nil {
		return false, err
	}
	var out bool
	err = json.Unmarshal(res, &out)
	return out, err
}

// Displayed returns whether the current element is displayed or not.
func (e *Element) Displayed() (bool, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/displayed", e.ws.ID, e.ID)
	if err != nil {
		return false, err
	}
	var out bool
	err = json.Unmarshal(res, &out)
	return out, err
}

// Selected returns whether the current element is selected or not.
func (e *Element) Selected() (bool, error) {
	_, res, err := e.ws.wd.get("/session/%s/element/%s/selected", e.ws.ID, e.ID)
	if err != nil {
		return false, err
	}
	var out bool
	err = json.Unmarshal(res, &out)
	return out, err
}

// Element searches for a single element on the page, starting from this element.
func (e *Element) Element(using FindStrategy, value string) (*Element, error) {
	opt := map[string]interface{}{"using": using, "value": value}
	_, res, err := e.ws.wd.post("/session/%s/element/%s/element", opt, e.ws.ID, e.ID)
	if err != nil {
		return nil, err
	}
	var out Element
	err = json.Unmarshal(res, &out)
	out.ws = e.ws
	return &out, err
}

// Elements searches for multiple elements on the page, starting from this element.
func (e *Element) Elements(using FindStrategy, value string) ([]*Element, error) {
	opt := map[string]interface{}{"using": using, "value": value}
	_, res, err := e.ws.wd.post("/session/%s/element/%s/elements", opt, e.ws.ID, e.ID)
	if err != nil {
		return nil, err
	}
	var out []*Element
	err = json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}
	for i := range out {
		out[i].ws = e.ws
	}
	return out, err
}
