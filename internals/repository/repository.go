package repository

import (
	"database/sql"
)

type Database interface {
	initDatabaseConn() error
}

type DB struct {
}

func (databaseStruct DB) initDatabaseConn() error {
	db, err := sql.Open("sqlite3", "./algs.db:?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()
	db.Exec(`create table algorithm (id integer not null primary key autoincrement, name text not null)`)
	db.Exec(`create table algorithm_version (id integer not null primary key autoicnrement, algorithm_id references algorithm(id), name text not null)`)
	db.Exec(`create table algorithm_version_run (id integer not null primary key autoicnrement, algorithm_version_id references algorithm_version(id), name text not null,
	keygen text not null, sign text not null, verify text not null, private_key_size ingeger not null, public_key_size integer not null, signature_size integer not null)`)
	db.Exec(`update players set name=?, score=?, active=? where jerseyNum=?`,
		"Smith, Steve", 42, true, "99")

	return nil
}
