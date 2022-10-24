/*
Copyright (c) 2016-2023 Uber Technologies, Inc.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package location

// LocationReferencePoint A Location Reference Point consists of coordinate, line attributes and path attributes
type LocationReferencePoint struct {

	// Lon WGS84 Longitude coordinates
	Lon float64

	// Lat WGS84 Latitude coordinates
	Lat float64

	// Frc Functional Road Class with 0 being the highest and 7 being of the lowest, in terms of importance
	Frc int

	// Fow Form of way
	Fow int

	// Bear Bearing angle in degrees
	Bear int

	// Lfrcnp Lowest FRC to the next point
	Lfrcnp int

	// Dnp Distance to next LR-point
	Dnp int
}
