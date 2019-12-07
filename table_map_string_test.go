package main

// func TestTableMapString(t *testing.T) {

// 	tmi := CreateTableMapString("test/tbl_s/")

// 	err := tmi.Flush()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if tmi.needSave == true {
// 		t.Fatal("tmi.needSave == true")
// 	}

// 	err = tmi.IndexIntAdd("tif1", "tbl_ind_field1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = tmi.IndexStringAdd("tsf1", "tbl_s_ind_field1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = tmi.Flush()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_, _, _, err = tmi.Set(ItemString{Key: "82739",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 10},
// 		FieldsString: map[string]string{"b": "0",
// 			"c":    "5",
// 			"tsf1": "10"}})

// 	_, n, eq, err := tmi.Set(ItemString{Key: "82734",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 10},
// 		FieldsString: map[string]string{"b": "0",
// 			"c":    "5",
// 			"tsf1": "10"}})

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if n == false {
// 		t.Fatal("value must be new")
// 	}
// 	if eq == true {
// 		t.Fatal("value must be not equal")
// 	}

// 	err = tmi.Flush()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	out, n, eq, err := tmi.Set(ItemString{Key: "82734",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 5}})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if n == false {
// 		t.Fatal("value must be new 2")
// 	}
// 	if eq == true {
// 		t.Fatal("value must be not equal 2")
// 	}

// 	_, n, eq, err = tmi.Set(out)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if n == true {
// 		t.Fatal("value must be not new")
// 	}
// 	if eq == false {
// 		t.Fatal("value must be equal")
// 	}

// 	_, _, _, err = tmi.Set(ItemString{Key: "82734",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 5}})

// 	_, n, eq, err = tmi.Set(out)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if n == true {
// 		t.Fatal("value must be not new 3")
// 	}
// 	if eq == true {
// 		t.Fatal("value must be not equal 3")
// 	}

// 	_, _, _, err = tmi.Set(ItemString{Key: "82732",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 5}})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_, _, _, err = tmi.Set(ItemString{Key: "82733",
// 		FieldsInt: map[string]int64{"b": 0,
// 			"c":    5,
// 			"tif1": 15}})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = tmi.Flush()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tmi2 := CreateTableMapString("test/tbl_s/")
// 	err = tmi2.Load()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	//t.Fatal(tmi2.IntIndexes["tif1"])
// 	//t.Fatal(tmi2)

// 	keys, errPub, errInt := tmi2.GetKeysByIndexInt("tif1", 5, 5)
// 	if errPub != nil {
// 		t.Fatal(errPub)
// 	}
// 	if errInt != nil {
// 		t.Fatal(errInt)
// 	}
// 	if len(keys) != 2 {
// 		//[82734 82732]
// 		t.Fatal("Not correct load itms and search keys ", keys)
// 	}

// 	itms, errPub, errInt := tmi2.GetListByIndexInt("tif1", []int64{5}, 5)
// 	if errPub != nil {
// 		t.Fatal(errPub)
// 	}
// 	if errInt != nil {
// 		t.Fatal(errInt)
// 	}
// 	if len(itms) == 0 {
// 		t.Fatal("Not correct search itms itms ", itms)
// 	}
// }
