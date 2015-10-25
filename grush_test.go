package main

import (
	"testing"
)

func TestStatus(t *testing.T) {
}

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}
