package main

import (
	"testing"
)

func TestSliceRemoveIntList(t *testing.T) {
	s := []int64{5, 8, 19, 15, 7, 9, 14, 15, 6}
	k := []int64{6, 19, 15, 9, 48, 21}

	//5 8 14 7

	r := SliceRemoveIntList(s, k)

	if len(r) != 4 {
		t.Fatal(r)
	}
}

func TestSliceRemoveInt(t *testing.T) {
	s := []int64{5, 8, 19, 15, 7, 9, 14, 15, 6}
	k := int64(15)

	//5 8 19 6 7 9 14

	r := SliceRemoveInt(s, k)

	if len(r) != 7 {
		t.Fatal(r)
	}
}

func TestSliceRemoveStringList(t *testing.T) {
	s := []string{"5", "8", "19", "15", "7", "9", "14", "15", "6"}
	k := []string{"6", "19", "15", "9", "48", "21"}

	//5 8 14 7

	r := SliceRemoveStringList(s, k)

	if len(r) != 4 {
		t.Fatal(r)
	}
}

func TestSliceRemoveString(t *testing.T) {
	s := []string{"5", "8", "19", "15", "7", "9", "14", "15", "6"}
	k := string("15")

	//5 8 19 6 7 9 14

	r := SliceRemoveString(s, k)

	if len(r) != 7 {
		t.Fatal(r)
	}
}

func BenchmarkSliceRemoveIntList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int64{5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15, 5, 8, 19, 15, 7, 9, 14, 15}
		k := []int64{6, 19, 15, 9, 48, 21}

		SliceRemoveIntList(s, k)

	}
}
