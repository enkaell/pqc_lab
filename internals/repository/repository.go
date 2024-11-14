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

func DBGetGetAllAlgVersionByAlgName(algorithm_name string) ([]*server.Algorithm_Version, error) {
	var versions []*server.Algorithm_Version

	rows, err := db.Query(`SELECT v.id, v.name from algorithms a, algorithm_versions v where  a.id = v.algorithm_id and a.name=?;`, algorithm_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ver := &server.Algorithm_Version{}
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

func DBGetAllAlgRunsByVersion(version_name string) ([]*server.Algorithm_Version_Run, error) {
	var version_runs []*server.Algorithm_Version_Run

	rows, err := db.Query(`SELECT r.id, r.name, r.keygen, r.sign, r.verify, r.private_key_size, r.public_key_size, r.signature_size
	 from algorithm_version_runs r, algorithm_versions v where  v.id = r.algorithm_version_id and v.name=?;`, version_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		runs := &server.Algorithm_Version_Run{}
		if err := rows.Scan(&runs.Id, &runs.Name, &runs.Keygen, &runs.Sign,
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

func DBPrepareAndRunStatementTransaction(vers []*server.Algorithm_Version_Run) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() //  The rollback will be ignored if the tx has been committed later in the function

	stmt, err := tx.Prepare("INSERT INTO algorithm_version_runs(name, keygen, sign, verify, private_key_size, public_key_size, signature_size) VALUES( ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use

	for _, ver := range vers {
		if _, err := stmt.Exec(ver.Name, ver.Keygen, ver.Sign, ver.Verify, ver.PrivateKeySize, ver.PublicKeySize, ver.SignatureSize); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

// #TODO: Database transaction flow - ?
// func DBWriteRunResults(algs *server.AlgorithmList) error {

// 	_, err := db.Exec(`INSERT INTO algorithms (name) VALUES (?)`, algs.Algorithms...)
// 	if err != nil {
// 		return fmt.Errorf("failed to batch insert algorithms: %w", err)
// 	}

// 	// Batch insert algorithm_versions
// 	_, err = db.Exec(`INSERT INTO algorithm_versions (algorithm_id, name) VALUES (?, ?)`, algs.Algorithms...)
// 	if err != nil {
// 		return fmt.Errorf("failed to batch insert versions: %w", err)
// 	}

// 	// Batch insert algorithm_version_runs
// 	_, err = db.Exec(`INSERT INTO algorithm_version_runs
//         (algorithm_version_id, name, keygen, sign, verify, private_key_size, public_key_size, signature_size)
//         VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, runsValues...)
// 	if err != nil {
// 		return fmt.Errorf("failed to batch insert runs: %w", err)
// 	}

// 	return nil

// }
