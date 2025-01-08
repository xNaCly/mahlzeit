-- do not use ;\n for any other case other than terminating statements. 
-- validate all statements in this file with sqleibniz before contributing
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS meals (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    minutes INTEGER DEFAULT 0,
    -- url to an image to use
    image TEXT DEFAULT '',
    -- url to a recipe
    recipe TEXT DEFAULT ''
) STRICT;
CREATE INDEX IF NOT EXISTS meal_id ON meals(id);

CREATE TABLE IF NOT EXISTS ingredients (
    id INTEGER PRIMARY KEY,
    meal_id INTEGER REFERENCES meals(id),
    name TEXT,
    -- used to group ingredients 
    kind TEXT DEFAULT 'ungrouped', 
    unit TEXT DEFAULT 'x',
    amount REAL DEFAULT 1
) STRICT;
