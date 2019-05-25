package repository

import (
	"testing"
)

func TestBuildConnString_ShouldReturn_Success(t *testing.T) {
	property := ConnProperty{
		UserName: "test",
		Password: "password",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	}
	connString, err := BuildConnString(&property)
	expected := "test:password@tcp(127.0.0.1:3306)/test"
	if err != nil {
		t.Errorf(err.Error())
	}
	if *connString != expected {
		t.Errorf("\nExpected: %s\n Actual: %s\n", expected, *connString)
	}
}

func TestBuildConnString_ShouldReturn_Error(t *testing.T) {
	property := ConnProperty{
		UserName: "test",
		Password: "password",
		Protocol: "test",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	}
	connString, err := BuildConnString(&property)
	if err == nil || connString != nil {
		t.Errorf("Should have protocol error")
	}
}

func TestRepository_ShouldReturn_Success(t *testing.T) {
	connString := "tester:password@tcp(localhost:23306)/RatkerDev"
	testRepository := Repository("test", connString)
	betaRepository := Repository("beta", connString)

	defer testRepository.Conn.Close()
	defer betaRepository.Conn.Close()

	if testRepository.Name != "test" || testRepository.Conn == nil {
		t.Errorf("Repository setup failed")
	}

	if betaRepository.Name != "beta" || betaRepository.Conn == nil {
		t.Errorf("Repository setup failed")
	}
}
