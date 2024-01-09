package ewkb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/hex"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
)

// Compile-time checks.
var (
	_ sql.Scanner   = new(Geometry[orb.Geometry])
	_ driver.Valuer = new(Geometry[orb.Geometry])
)

type Geometry[T orb.Geometry] struct {
	Geom  T
	SRID  int
	Valid bool
}

func New[T orb.Geometry](geom T, srid int) Geometry[T] {
	return Geometry[T]{
		SRID:  srid,
		Geom:  geom,
		Valid: true,
	}
}

// The scanner/valuer implementations of this generic geometry wrapper rely on the fact that that a
// string can be cast into a PostGIS geometry/geography type as long as it is a hex-encoded string
// of Extended Well-Known Bytes (HEXEWKB).
//
// The purpose of this is two-fold:
//   - To sidestep the need to access spatial types via conversion functions (e.g. ST_AsEWKB and
//     ST_GeomFromEWBK).
//   - To enable the use type casts as way to inform sqlc (or more specifically, the PostgreSQL
//     language engine it uses) which type it should actually use for the generated Go code in a
//     type-safe manner.
//
// However, this does come with the downside of potentially doubling the amount of information
// passed over the wire as we are now sending/receiving a hex-string instead of the raw bytes it
// would represent.
func (g *Geometry[T]) Scan(x any) error {
	// If string, try decode to bytes.
	if str, ok := x.(string); ok {
		b, err := hex.DecodeString(str)
		if err == nil {
			// Was a valid hex string! Replace scan value.
			x = b
		}
	}

	scanner := ewkb.Scanner(&g.Geom)

	err := scanner.Scan(x)
	if err != nil {
		var emptyGeom T
		g.Geom = emptyGeom
		g.SRID = 0
		g.Valid = false
		return err
	}

	g.SRID = scanner.SRID
	g.Valid = scanner.Valid
	return nil
}

func (g Geometry[T]) Value() (driver.Value, error) {
	if g.Valid {
		return ewkb.MarshalToHex(g.Geom, g.SRID)
	}

	return nil, nil
}

// Aliases for use with the sqlc type overrides.
type (
	Any             = Geometry[orb.Geometry]
	Bound           = Geometry[orb.Bound]
	Collection      = Geometry[orb.Collection]
	LineString      = Geometry[orb.LineString]
	MultiLineString = Geometry[orb.MultiLineString]
	Point           = Geometry[orb.Point]
	MultiPoint      = Geometry[orb.MultiPoint]
	Polygon         = Geometry[orb.Polygon]
	MultiPolygon    = Geometry[orb.MultiPolygon]
)
