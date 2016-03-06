package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

var dbpath = "/tmp/ts3db.db"

func TestCheckStoryId(t *testing.T) {
	asst := assert.New(t)
	CreateSqliteDb()
	CreateTable()
	dao := &SqliteDao{DbPath: dbpath}
	err := dao.AddPostedStory(int64(999))
	if err != nil {
		asst.Fail(fmt.Sprint(err))
	}
	result999, err := dao.IsNewSotry(int64(999))
	if err != nil {
		asst.Fail(fmt.Sprint(err))
	}
	asst.Equal(false, result999)
	resultxxx, err := dao.IsNewSotry(int64(11111))
	if err != nil {
		asst.Fail(fmt.Sprint(err))
	}
	asst.Equal(true, resultxxx)
}

func CreateSqliteDb() {
	DeleteSqliteDb()
	cmd := exec.Command("sqlite3", dbpath)
	cmd.Output()
	cmd.Process.Kill()
}

func DeleteSqliteDb() {
	cmd := exec.Command("rm", "-f", dbpath)
	cmd.Output()
}

func CreateTable() {
	sql_crate_table := `
CREATE TABLE AL_POST_STORY (
	ID	INTEGER,
	PRIMARY KEY(ID)
);
	`
	db, _ := sql.Open("sqlite3", dbpath)
	db.Exec(sql_crate_table)
}
