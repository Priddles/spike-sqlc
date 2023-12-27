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
	_ sql.Scanner   = new(Geometry[orb.Point])
	_ driver.Valuer = new(Geometry[orb.Point])
)

type Geometry[T orb.Geometry] struct {
	Geom  T
	SRID  int
	Valid bool
}

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
		return ewkb.Value(g.Geom, g.SRID).Value()
	}

	return nil, nil
}

// Aliases for use with the SQLC type overrides.
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
