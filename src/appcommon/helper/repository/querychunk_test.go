package repository

import "testing"

func TestFromQueryChunkShouldReturnCorrectString(t *testing.T) {
	actual := new(DbQuery)
	expected := new(DbQuery)
	*expected = "test from testTable"
	// set initial value
	*actual = "test"
	actual = actual.FromQueryChunk("testTable")

	if *actual != *expected {
		t.Errorf("\nExpected: %s,\nActual: %s\n", *expected, *actual)
	}
}

func MultiFromQueryCheckShouldReturnCorrectString(t *testing.T) {
	actual := new(DbQuery)
	expected := new(DbQuery)
	*expected = "test from a, b, c"
	// set initial value
	*actual = "test"
	actual = actual.MultiFromQueryCheck("a", "b", "c")

	if *actual != *expected {
		t.Errorf("\nExpected: %s,\nActual: %s\n", *expected, *actual)
	}
}

func TestFinishQueryShouldReturnCorrectString(t *testing.T) {
	actual := new(DbQuery)
	expected := new(DbQuery)
	*expected = "test;"
	// set initial value
	*actual = "test"
	actual = actual.FinishQuery()

	if *actual != *expected {
		t.Errorf("\nExpected: %s,\nActual: %s\n", *expected, *actual)
	}
}
