package __ErrHandle

import (
	"database/sql"
	"errors"
	"log"
)

func queryDataFromDB(db sql.DB) (res []string, err error) {

	rows, err := db.Query("select id, name from users where id = ?", 1)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("result is empty")
		return res, err
	} else {
		log.Fatal(err)
		return res, err
	}

	for rows.Next() {
		//todo
	}
	return
}
