package helpers_test

import (
	"os"
	"testing"

	"github.com/themccallister/go-https-demo/helpers"
)

func TestEnvBoolOr(t *testing.T) {
	os.Setenv("TEST_ENV", "true")

	e := helpers.EnvBoolOr("TEST_ENV", false)

	if e != true {
		t.Fatalf("expected the env var to be true got %v instead", e)
	}

	os.Setenv("TEST_ENV", "")

	ne := helpers.EnvBoolOr("TEST_ENV", false)

	if ne != false {
		t.Fatalf("expected the env var to be false, got %v instead", ne)
	}

	// string must be true or false
	os.Setenv("TEST_ENV", "this is not true")

	nt := helpers.EnvBoolOr("TEST_ENV", false)
	if nt != false {
		t.Fatalf("expected the env var to be false, got %v instead", nt)
	}
}

func TestEnvStringOr(t *testing.T) {
	os.Setenv("TEST_ENV", "testing")

	e := helpers.EnvStringOr("TEST_ENV", "not testing")

	if e != "testing" {
		t.Fatalf("expected the string to be testing, got %v instead", e)
	}

	os.Setenv("TEST_ENV", "")

	ne := helpers.EnvStringOr("TEST_ENV", "not testing")

	if ne != "not testing" {
		t.Fatalf("expected the string to be not testing, got %v instead", ne)
	}
}
