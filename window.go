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

// Window represents a browser window.
type Window struct {
	ID string `json:"WINDOW"`
	ws *Session
}

// Close closes the window.
func (w *Window) Close() error {
	_, _, err := w.ws.wd.del("/session/%s/window", w.ws.ID)
	return err
}

// Resize resizes the window to the specified size.
func (w *Window) Resize(width, height int) error {
	opt := map[string]interface{}{"width": width, "height": height}
	_, _, err := w.ws.wd.post("/session/%s/window/%s/size", opt, w.ws.ID, w.ID)
	return err
}

// Minimize minimizes the browser window.
func (w *Window) Minimize() error {
	_, _, err := w.ws.wd.post("/session/%s/window/%s/minimize", nil, w.ws.ID, w.ID)
	return err
}

// Maximize maximizes the browser window.
func (w *Window) Maximize() error {
	_, _, err := w.ws.wd.post("/session/%s/window/%s/maximize", nil, w.ws.ID, w.ID)
	return err
}
