CREATE extension IF NOT EXISTS "postgis";

CREATE TABLE
    IF NOT EXISTS foo (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        area geom_any
    );

ALTER TABLE foo
ADD COLUMN IF NOT EXISTS location geom_point;
