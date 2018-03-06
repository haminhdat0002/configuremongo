package configuremongo

import (
	"fmt"
	"testing"
)

func test(received, expected interface{}, failFunc func()) {
	if received == expected {
		return
	}

	fmt.Printf("Failed test. Got '%[1]v' (type %[1]T), expected '%[2]v' (type %[2]T)\n", received, expected)

	failFunc()
}

func normalFailFunc(t *testing.T) func() {
	return func() {
		t.Fail()
	}
}

func TestSetupMongoConf(t *testing.T) {
	mongoConf := NewMongo("mongodb://127.0.0.1:27017/test_conf_db", "confs")
	mongoConf.Setup()
}

func TestMongoStrings(t *testing.T) {
	mongoConf := NewMongo("mongodb://127.0.0.1:27017/test_conf_db", "confs")
	mongoConf.Setup()

	ff := normalFailFunc(t)

	val, err := mongoConf.String("x")
	if err != nil {
		t.Errorf("Error while retrieving value : x")
	}
	test(val, "hello", ff)
	val, err = mongoConf.String("y")
	if err != nil {
		t.Errorf("Error while retrieving value : y")
	}
	test(val, "hello world", ff)
}
