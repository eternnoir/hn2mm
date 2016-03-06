package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteChecker struct {
	DbPath             string
	postedSotriesCache []int64
}

func NewSqliteChecker(dbpath string) (*SqliteChecker, error) {
	checker := &SqliteChecker{DbPath: dbpath, postedSotriesCache: []int64{}}
	return checker, nil
}

func (dao *SqliteChecker) IsNewSotry(id int64) (bool, error) {
	if dao.isInCache(id) {
		return false, nil
	}
	result, err := dao.isInDb(id)
	if err != nil {
		return false, err
	}
	if result {
		dao.AddIdToCache(id)
		return false, nil
	} else {
		return true, nil
	}
}

func (dao *SqliteChecker) isInCache(id int64) bool {
	for _, cacheid := range dao.postedSotriesCache {
		if cacheid == id {
			return true
		}
	}
	return false
}

func (dao *SqliteChecker) isInDb(id int64) (bool, error) {
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
		return true, nil
	} else {
		return false, nil
	}
}

func (dao *SqliteChecker) AddIdToCache(id int64) {
	dao.postedSotriesCache = append(dao.postedSotriesCache, id)
}

func (dao *SqliteChecker) AddPostedStory(id int64) error {
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
