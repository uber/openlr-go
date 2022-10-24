/*
Copyright (c) 2016-2023 Uber Technologies, Inc.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package binary

import (
	"math"
)

const (
	DECA_MICRO_DEG_FACTOR = 100000.0
	DISTANCE_PER_INTERVAL = 58.6
	RESOLUTION            = 24
	BEAR_SECTOR           = 11.25
)

// Deg2Int
//
//	@param deg coordinate in degrees
//	@return int integer resolution of degrees as 3 bytes
func Deg2Int(deg float64) int {
	val := math.Copysign(1, deg)*0.5 + float64(deg*(1<<RESOLUTION))/360.0

	return int(math.Round(val))
}

// Int2Bytes
//
//	@param val integer value to convert
//	@param size
//	@param signed
//	@return []byte converted integer expressed as big endian
func Int2Bytes(val int, size int, signed bool) []byte {
	bs := make([]byte, size)
	max := 1 << (8 * size)
	if val < 0 {
		val += max
	}
	for i := size - 1; i >= 0; i-- {
		r := val % 256
		val = val / 256
		bs[i] = byte(r)
	}
	return bs
}

// Status2Bytes
//
//	@param version 		openlr version
//	@param locationType	location reference type
//	@return []byte 		status represented as a byte array of size 1
func Status2Bytes(version int, locationType int) []byte {
	version &= 0b111
	locationType &= 0b1111
	byteInteger := version + (locationType << 3)

	return Int2Bytes(byteInteger, 1, false)
}

// Offset2Bytes
//
//	@param offset 	offset rate
//	@return []byte 	offset represented as a byte array of size 1
func Offset2Bytes(offset float64) []byte {
	idx := 0
	if offset != 0.0 {
		idx = int(math.Round(offset*256 - 0.5))
	}

	return Int2Bytes(idx, 1, false)
}

// Coords2Bytes
//
//	@param lon 		longitude in degrees
//	@param lat 		latitude in degrees
//	@return []byte 	longitude and latitude represented as a byte array of size 6
func Coords2Bytes(lon float64, lat float64) []byte {
	lonInt := Deg2Int(lon)
	latInt := Deg2Int(lat)

	return append(Int2Bytes(lonInt, 3, true), Int2Bytes(latInt, 3, true)...)
}

// Attributes2Bytes
//
//	@param fow 		form of way
//	@param frc 		functional road class
//	@param bear 	bearing angle
//	@param lfrcnp 	lowest frc to next point
//	@param reserved reserved for future usage/side of road/orientation
//	@return []byte 	point attributes represented as a byte array of size 2
func Attributes2Bytes(fow int, frc int, bear int, lfrcnp int, reserved int) []byte {
	bearfloat := (float64(bear) - BEAR_SECTOR/2.0) / BEAR_SECTOR

	bearInt := int(math.Round(bearfloat))
	fow = fow & 0b111
	frc = frc & 0b111
	bearInt = bearInt & 0b11111
	lfrcnp = lfrcnp & 0b111
	reserved = reserved & 0b11
	bytesFowFrcReserved := Int2Bytes(fow+(frc<<3)+(reserved<<6), 1, false)
	bytesBearLFRCNP := Int2Bytes(bearInt+(lfrcnp<<5), 1, false)

	return append(bytesFowFrcReserved, bytesBearLFRCNP...)
}

// Dnp2Bytes
//
//	@param dnp 		distance to next point
//	@return []byte 	dnp represented as 1 byte array
func Dnp2Bytes(dnp int) []byte {
	interval := float64(dnp)/DISTANCE_PER_INTERVAL - 0.5
	intervalInteger := int(math.Round(interval))

	return Int2Bytes(intervalInteger, 1, false)
}

// RelativeCoords2Bytes
//
//	@param lon 		longitude
//	@param lat 		latitude
//	@param xLon 	relative longitude
//	@param xLat 	relative latitude
//	@return []byte 	relative coordinates represented as a byte array of size 4
func RelativeCoords2Bytes(lon float64, lat float64, xLon float64, xLat float64) []byte {
	relLon := int(math.Round(DECA_MICRO_DEG_FACTOR * (lon - xLon)))
	relLat := int(math.Round(DECA_MICRO_DEG_FACTOR * (lat - xLat)))

	return append(Int2Bytes(relLon, 2, true), Int2Bytes(relLat, 2, true)...)
}
