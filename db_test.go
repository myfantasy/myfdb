package main

import (
	"testing"
	"time"
)

func TestCreateDB(t *testing.T) {

	db, err := CreateDB("test/db1/", time.Second*2, func(e error) { t.Fatal(e) })
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateTableInt(TableMapIntName, "tst_tbl1", 10)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Flush()
	if err != nil {
		t.Fatal(err)
	}

	_, err = CreateDB("test/db1/", time.Second*3, func(e error) { t.Fatal(e) })

}
