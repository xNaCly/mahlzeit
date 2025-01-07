-- do not use ;\n for any other case other than terminating statements. 
-- validate all statements in this file with sqleibniz before contributing
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS meals (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    image TEXT DEFAULT '',
    recipe TEXT DEFAULT ''
) STRICT;
CREATE INDEX IF NOT EXISTS meal_id ON meals(id);

CREATE TABLE IF NOT EXISTS ingredients (
    id INTEGER PRIMARY KEY,
    meal_id INTEGER REFERENCES meals(id),
    name TEXT,
    unit TEXT DEFAULT 'x',
    amount REAL DEFAULT 1
) STRICT;
