package test

import (
	"os"
	"testing"
)

func TestDate(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(wd)
}