/*
Copyright (c) 2016-2023 Uber Technologies, Inc.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package location_test

import (
	b64 "encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/openlr-go/location"
)

func TestEncodings(t *testing.T) {
	var tests = []struct {
		llr      location.LineLocationReference
		expected string
	}{
		{
			llr: location.LineLocationReference{
				Poffs: 0.26757812,
				Noffs: 0,
				Points: []location.LocationReferencePoint{
					location.LocationReferencePoint{Lon: 6.1268198,
						Lat:    49.6085178,
						Frc:    3,
						Fow:    2,
						Bear:   141,
						Lfrcnp: 3,
						Dnp:    557},
					location.LocationReferencePoint{Lon: 6.1283698,
						Lat:    49.6039878,
						Frc:    3,
						Fow:    3,
						Bear:   231,
						Lfrcnp: 5,
						Dnp:    264},
					location.LocationReferencePoint{Lon: 6.1281598,
						Lat:    49.6030578,
						Frc:    5,
						Fow:    3,
						Bear:   287,
						Lfrcnp: 7,
						Dnp:    0},
				},
			},
			expected: "CwRbWyNG9RpsCQCb/jsbtAT/6/+jK1lE",
		},
		{
			llr: location.LineLocationReference{
				Poffs: 0,
				Noffs: 0.45898438,
				Points: []location.LocationReferencePoint{
					location.LocationReferencePoint{Lon: 0.6752192,
						Lat:    47.3651611,
						Frc:    3,
						Fow:    4,
						Bear:   28,
						Lfrcnp: 3,
						Dnp:    498},
					location.LocationReferencePoint{Lon: 0.6769992,
						Lat:    47.3696011,
						Frc:    3,
						Fow:    2,
						Bear:   197,
						Lfrcnp: 7,
						Dnp:    0},
				},
			},
			expected: "CwB67CGukRxiCACyAbwaMXU=",
		},
		{
			llr: location.LineLocationReference{
				Poffs: 0,
				Noffs: 0,
				Points: []location.LocationReferencePoint{
					location.LocationReferencePoint{Lon: 9.9750602,
						Lat:    48.0632865,
						Frc:    1,
						Fow:    3,
						Bear:   298,
						Lfrcnp: 1,
						Dnp:    88},
					location.LocationReferencePoint{Lon: 9.9750602,
						Lat:    48.0632865,
						Frc:    1,
						Fow:    3,
						Bear:   298,
						Lfrcnp: 7,
						Dnp:    0},
				},
			},
			expected: "CwcX6CItqAs6AQAAAAALGg==",
		},
		{
			llr: location.LineLocationReference{
				Poffs: 0,
				Noffs: 0,
				Points: []location.LocationReferencePoint{
					location.LocationReferencePoint{Lon: 6.1268198,
						Lat:    49.6084964,
						Frc:    3,
						Fow:    2,
						Bear:   6,
						Lfrcnp: 3,
						Dnp:    29},
					location.LocationReferencePoint{Lon: 6.1283598,
						Lat:    49.6039664,
						Frc:    3,
						Fow:    3,
						Bear:   6,
						Lfrcnp: 5,
						Dnp:    29},
					location.LocationReferencePoint{Lon: 6.1281498,
						Lat:    49.6030464,
						Frc:    5,
						Fow:    3,
						Bear:   6,
						Lfrcnp: 7,
						Dnp:    0},
				},
			},
			expected: "CwRbWyNG9BpgAACa/jsboAD/6/+kKwA=",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%+v == %s ?", tt.llr, tt.expected)
		t.Run(testname, func(t *testing.T) {
			encodedLocation, _ := tt.llr.Encode()
			assert.Equal(t, tt.expected, b64.StdEncoding.EncodeToString(encodedLocation))
		})
	}
}
