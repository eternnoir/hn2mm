package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDao struct {
	DbPath string
}

func (dao *SqliteDao) IsNewSotry(id int64) (bool, error) {
	db, err := sql.Open("sqlite3", dao.DbPath)
	if err != nil {
		return false, err
	}
	defer db.Close()
	rows, err := db.Query("select * from AL_POST_STORY where ID=?", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return false, nil
	} else {
		return true, nil
	}
}

func (dao *SqliteDao) AddPostedStory(id int64) error {
	db, err := sql.Open("sqlite3", dao.DbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into AL_POST_STORY (ID) values (?)", id)
	if err != nil {
		return err
	}
	return nil
}
