package handlers

import "testing"

func Test_generateTokenKey(t *testing.T) {
	k := generateTokenKey()
	if len(k) != 9 {
		t.Fail()
	}
}

func Test_generateAPIKey(t *testing.T) {
	k := generateAPIKey()
	if len(k) != 35 {
		t.Fail()
	}
}
