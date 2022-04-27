package main

import "testing"

type TestTable struct {
	xarg int
	yarg int
}

func TestAdd(t *testing.T) {
	t.Log(Add(1, 2))
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

	sum = Add(2, 4)
	if sum == 6 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

	tables := []TestTable{
		{1, 2},
		{2, 4},
		{4, 8},
		{5, 10},
		{6, 12},
	}

	for _, table := range tables {
		result := Add(table.xarg, table.yarg)
		if result == (table.xarg + table.yarg) {
			t.Log("the result is ok")
		} else {
			t.Fatal("the result is wrong")
		}
	}
}
