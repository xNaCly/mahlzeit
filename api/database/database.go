package database

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"database/sql"
	"log/slog"

	_ "github.com/glebarez/go-sqlite"
	"github.com/xnacly/mahlzeit/assert"
)

//go:embed init.sql
var init_sql string

//go:embed default.sql
var default_sql string

type Database struct {
	m    sync.Mutex
	conn *sql.DB
}

func New(withDefault bool) (*Database, error) {
	var path string
	if envPath := os.Getenv("MAHLZEIT_DB"); len(envPath) != 0 {
		slog.Info("got MAHLZEIT_DB, switching to custom db path", "MAHLZEIT_DB", envPath)
		absPath, err := filepath.Abs(envPath)
		path = absPath
		assert.NoError(err, "ctx", "tried to convert MAHLZEIT_DB to an absolute path")
	} else {
		config, err := os.UserConfigDir()
		assert.NoError(err, "ctx", "tried to get UserConfigDir")
		path = filepath.Join(config, "mahlzeit", "mahlzeit.db")
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return nil, err
		}
	}

	slog.Info("backing database path", "path", path)
	conn, err := sql.Open("sqlite", path)
	assert.NoError(err, "ctx", "tried to connect to database")
	assert.NoError(conn.Ping(), "ctx", "tried to ping the database")
	assert.True(len(init_sql) > 0, "ctx", "init.sql should not be empty")

	for _, stmt := range strings.Split(init_sql, ";\n") {
		slog.Info("executing init stmt", "stmt", stmt)
		_, err := conn.Exec(strings.TrimSpace(stmt))
		assert.NoError(err, "ctx", "attempted to execute stmt from init.sql", "stmt", stmt)
	}

	if withDefault {
		slog.Info("got withDefault option, inserting default meals")
		for _, stmt := range strings.Split(default_sql, ";\n") {
			_, err := conn.Exec(strings.TrimSpace(stmt))
			assert.NoError(err, "ctx", "attempted to execute stmt from default.sql", "stmt", stmt)
		}
	}

	return &Database{
		conn: conn,
		m:    sync.Mutex{},
	}, nil
}

func (d *Database) Destroy() error {
	d.m.Lock()
	err := d.conn.Close()
	d.m.Unlock()
	return err
}
