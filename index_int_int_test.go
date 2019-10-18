package main

import "testing"

func TestIndexIntInt(t *testing.T) {

	mi := CreateIndexIntInt()

	mi.Set(234, []int64{7, 23, 465, 7, 9, 5, 3423, 2}, []int64{2, 7, 8})
	mi.Set(234, []int64{12}, []int64{2, 7, 8})
	mi.Set(2, []int64{7, 23, 465, 7, 9, 5, 3423, 2}, []int64{2, 7, 8})
	mi.Set(4, []int64{7, 23, 465, 7, 9, 5, 3423, 2}, []int64{2, 7, 8})

	err := mi.Flush("test/tst.json")
	if err != nil {
		t.Fatal(err)
	}

	mi2 := CreateIndexIntInt()

	err = mi2.Load("test/tst.json")
	if err != nil {
		t.Fatal(err)
	}
	err = mi2.Flush("test/tst2.json")
	if err != nil {
		t.Fatal(err)
	}

}
