/*
Copyright (c) 2016-2023 Uber Technologies, Inc.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package location

import (
	"github.com/uber/openlr-go/binary"
)

// LineLocationReference describes a path within a map and consists of location reference point(s),
// including a last location reference point and positive and negative offset data.
type LineLocationReference struct {

	// Points
	Points []LocationReferencePoint

	// Poffs Positive offset is the difference between the start point of the location and start point of the desired location
	Poffs float64

	// Noffs Negative offset is the difference between the endpoint of the desired location and the end point of the location
	Noffs float64
}

// Encode
//
//	@receiver llr Line Location Reference
//	@return []byte binary representation in bytes of the Location Reference
//	@return error
func (llr LineLocationReference) Encode() ([]byte, error) {
	bs := []byte{}
	bs = append(bs, binary.Status2Bytes(3, 1)...)
	firstPoint := llr.Points[0]
	bs = append(bs, binary.Coords2Bytes(firstPoint.Lon, firstPoint.Lat)...)
	bs = append(bs, binary.Attributes2Bytes(firstPoint.Fow, firstPoint.Frc, firstPoint.Bear, firstPoint.Lfrcnp, 0)...)
	bs = append(bs, binary.Dnp2Bytes(firstPoint.Dnp)...)
	prevPoint := firstPoint
	for i := 1; i < len(llr.Points)-1; i++ {
		point := llr.Points[i]
		bs = append(bs, binary.RelativeCoords2Bytes(point.Lon, point.Lat, prevPoint.Lon, prevPoint.Lat)...)
		bs = append(bs, binary.Attributes2Bytes(point.Fow, point.Frc, point.Bear, point.Lfrcnp, 0)...)
		bs = append(bs, binary.Dnp2Bytes(point.Dnp)...)
		prevPoint = point
	}
	lastPoint := llr.Points[len(llr.Points)-1]
	bs = append(bs, binary.RelativeCoords2Bytes(lastPoint.Lon, lastPoint.Lat, prevPoint.Lon, prevPoint.Lat)...)
	pOffsetFlag := 0
	nOffsetFlag := 0
	pOffsetBytes := []byte{}
	nOffsetBytes := []byte{}
	if llr.Poffs > 0 {
		pOffsetFlag = 2
		pOffsetBytes = binary.Offset2Bytes(llr.Poffs)
	}
	if llr.Noffs > 0 {
		nOffsetFlag = 1
		nOffsetBytes = binary.Offset2Bytes(llr.Noffs)
	}
	offsetFlags := nOffsetFlag + pOffsetFlag
	bs = append(bs, binary.Attributes2Bytes(lastPoint.Fow, lastPoint.Frc, lastPoint.Bear, offsetFlags, 0)...)
	bs = append(bs, pOffsetBytes...)
	bs = append(bs, nOffsetBytes...)

	return bs, nil
}
