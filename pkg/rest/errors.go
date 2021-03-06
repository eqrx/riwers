/*
Copyright 2019 Alexander Sowitzki.

GNU Affero General Public License version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/AGPL-3.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"errors"
	"fmt"
)

// HTTPError represents an HTTP error in combination with an HTTP status code.
type HTTPError struct {
	URL        string
	StatusCode int
	Text       string
}

// Error returns the error as string.
func (h HTTPError) Error() string {
	return fmt.Sprintf("%v -> %v %v", h.URL, h.StatusCode, h.Text)
}

// ErrRequest happens when a HTTP request is invalid.
var ErrRequest = errors.New("invalid request")
