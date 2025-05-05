package storage

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type KelcaDB struct {
	db *sql.DB
}

func NewKelcaDB(basePath string) (*KelcaDB, error) {
	dbPath := filepath.Join(basePath, "db", "kelca.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := initializeSchema(db); err != nil {
		db.Close()
		return nil, err
	}

	return &KelcaDB{db: db}, nil
}

func initializeSchema(db *sql.DB) error {
	schemaSQL := `
		CREATE TABLE IF NOT EXISTS certificates (
			id INTEGER PRIMARY KEY,
			serial_number TEXT NOT NULL UNIQUE,
			subject TEXT NOT NULL,
			issuer_id INTEGER,
			not_before TEXT NOT NULL,
			not_after TEXT NOT NULL,
			key_path TEXT,
			cert_path TEXT NOT NULL,
			status TEXT NOT NULL,
			created_at TEXT NOT NULL,
			FOREIGN KEY (issuer_id) REFERENCES certificates (id)
		);

		CREATE TABLE IF NOT EXISTS keys (
			id INTEGER PRIMARY KEY,
			key_id TEXT NOT NULL UNIQUE,
			algorithm TEXT NOT NULL,
			key_size INTEGER NOT NULL,
			usage TEXT NOT NULL,
			key_path TEXT NOT NULL,
			created_at TEXT NOT NULL,
			last_accessed TEXT,
			access_count INTEGER DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS revocations (
			cert_id INTEGER NOT NULL,
			revocation_date TEXT NOT NULL,
			reason TEXT,
			FOREIGN KEY (cert_id) REFERENCES certificates (id)
		);
	`

	_, err := db.Exec(schemaSQL)
	return err
}
