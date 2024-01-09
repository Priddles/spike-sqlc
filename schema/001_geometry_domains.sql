CREATE extension IF NOT EXISTS "postgis";

-- Use domains to effectively create type aliases with different geometry and SRID constraints
-- applied.
--
-- These aliases should be used instead of the underlying geometry/geography as sqlc has been
-- configured to use the corresponding Go types when generating code.
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_any') THEN
        CREATE DOMAIN geom_any AS geometry(Geometry, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_bound') THEN
        CREATE DOMAIN geom_bound AS geometry(Polygon, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_collection') THEN
        CREATE DOMAIN geom_collection AS geometry(GeometryCollection, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_linestring') THEN
        CREATE DOMAIN geom_linestring AS geometry(LineString, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_multilinestring') THEN
        CREATE DOMAIN geom_multilinestring AS geometry(MultiLineString, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_point') THEN
        CREATE DOMAIN geom_point AS geometry(Point, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_multipoint') THEN
        CREATE DOMAIN geom_multipoint AS geometry(MultiPoint, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_polygon') THEN
        CREATE DOMAIN geom_polygon AS geometry(Polygon, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geom_multipolygon') THEN
        CREATE DOMAIN geom_multipolygon AS geometry(MultiPolygon, 4326);
    END IF;
END;
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_any') THEN
        CREATE DOMAIN geog_any AS geography(Geometry, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_bound') THEN
        CREATE DOMAIN geog_bound AS geometry(Polygon, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_collection') THEN
        CREATE DOMAIN geog_collection AS geography(GeometryCollection, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_linestring') THEN
        CREATE DOMAIN geog_linestring AS geography(LineString, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_multilinestring') THEN
        CREATE DOMAIN geog_multilinestring AS geography(MultiLineString, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_point') THEN
        CREATE DOMAIN geog_point AS geography(Point, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_multipoint') THEN
        CREATE DOMAIN geog_multipoint AS geography(MultiPoint, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_polygon') THEN
        CREATE DOMAIN geog_polygon AS geography(Polygon, 4326);
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'geog_multipolygon') THEN
        CREATE DOMAIN geog_multipolygon AS geography(MultiPolygon, 4326);
    END IF;
END;
$$;
