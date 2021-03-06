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

// Package common contains the measurement struct that is used by all implementatons.
package common

import (
	"time"
)

// Measurement contains the compensated measurements of a BME680 and its timestamp and tags.
type Measurement struct {
	GasResistance float64           `json:"gas_resistance"`
	Humidity      float64           `json:"humidity"`
	Pressure      float64           `json:"pressure"`
	Temperature   float64           `json:"temperature"`
	Timestamp     time.Time         `json:"timestamp"`
	Tags          map[string]string `json:"tags"`
}
