package main

import (
	"os"
	"strings"
	"testing"
)

func TestParsePuppetReport(t *testing.T) {
	f, err := os.Open("testdata/report.txt")
	if err != nil {
		t.Errorf("Failed to open testdata file: %v", err)
	}
	report, err := parsePuppetReport(f)
	if err != nil {
		t.Errorf("Parsing report failed: %v", err)
	}
	if strings.Contains(string(report), "gemstash1-bleep") != true {
		t.Errorf("Report does not contain expected data")
	}
}
