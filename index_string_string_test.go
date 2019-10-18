package main

import "testing"

func TestIndexStringString(t *testing.T) {

	mi := CreateIndexStringString()

	mi.Set("234", []string{"7", "23", "465", "7", "9", "5", "3423", "2"}, []string{"2", "7", "8"})
	mi.Set("234", []string{"12"}, []string{"2", "7", "8"})
	mi.Set("2", []string{"7", "23", "465", "7", "9", "5", "3423", "2"}, []string{"2", "7", "8"})
	mi.Set("4", []string{"7", "23", "465", "7", "9", "5", "3423", "2"}, []string{"2", "7", "8"})

	err := mi.Flush("test/tst_s_ss.json")
	if err != nil {
		t.Fatal(err)
	}

	mi2 := CreateIndexStringString()

	err = mi2.Load("test/tst_s_ss.json")
	if err != nil {
		t.Fatal(err)
	}
	err = mi2.Flush("test/tst2_s_ss.json")
	if err != nil {
		t.Fatal(err)
	}

}
