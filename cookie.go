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

// Cookie represents a web cookie.
type Cookie struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Path   string `json:"path"`
	Domain string `json:"domain"`
	Secure bool   `json:"secure"`
	Expiry int    `json:"expiry"`
	ws     *Session
}

// Set saves the cookie to the current session.
func (c *Cookie) Set() error {
	opt := map[string]interface{}{"cookie": c}
	_, _, err := c.ws.wd.post("/session/%s/cookie", opt, c.ws.ID)
	return err
}

// Clear removes the cookie from the current session.
func (c *Cookie) Clear() error {
	_, _, err := c.ws.wd.del("/session/%s/cookie/%s", c.ws.ID, c.Name)
	return err
}
