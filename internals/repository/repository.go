package repository

import (
	"database/sql"
	"fmt"

	"github.com/enkaell/pqc_lab/internals/server"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	initDatabaseConn() error
}

var db *sql.DB

func DBInitDatabaseConn() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./algs.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS algorithms (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS algorithm_versions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		algorithm_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		FOREIGN KEY (algorithm_id) REFERENCES algorithms(id)
	)`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS algorithm_version_runs (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		algorithm_version_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		keygen TEXT NOT NULL,
		sign TEXT NOT NULL,
		verify TEXT NOT NULL,
		private_key_size INTEGER NOT NULL,
		public_key_size INTEGER NOT NULL,
		signature_size INTEGER NOT NULL,
		FOREIGN KEY (algorithm_version_id) REFERENCES algorithm_versions(id)
	)`)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")
	return db, nil
}

func DBGetAllAlgorithms() ([]*server.Algorithm, error) {
	var algorithms []*server.Algorithm

	rows, err := db.Query(`SELECT * FROM algorithms;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		alg := &server.Algorithm{}
		if err := rows.Scan(&alg.Id, &alg.Name); err != nil {
			return nil, err
		}
		algorithms = append(algorithms, alg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return algorithms, nil
}

func DBGetGetAllAlgVersionByAlgName(algorithm_name string) ([]*server.Version, error) {
	var versions []*server.Version

	rows, err := db.Query(`SELECT v.id, v.name from algorithms a, algorithm_versions v where  a.id = v.algorithm_id and a.name=?;`, algorithm_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ver := &server.Version{}
		if err := rows.Scan(&ver.Id, &ver.Name); err != nil {
			return nil, err
		}
		versions = append(versions, ver)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return versions, nil
}

func DBGetAllAlgRunsByVersion(version_name string) ([]*server.Version_Run, error) {
	var version_runs []*server.Version_Run

	rows, err := db.Query(`SELECT r.* from algorithm_version_runs r, algorithm_versions v where  v.id = r.algorithm_version_id and v.name=?;`, version_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		runs := &server.Version_Run{}
		if err := rows.Scan(&runs.Id, &runs.AlgorithmVersionId, &runs.Keygen, &runs.Keygen, &runs.Sign,
			&runs.Verify, &runs.PrivateKeySize, &runs.PublicKeySize, &runs.SignatureSize); err != nil {
			return nil, err
		}
		version_runs = append(version_runs, runs)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return version_runs, nil
}
