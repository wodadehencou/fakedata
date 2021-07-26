package fakedata_test

import (
	"regexp"
	"testing"

	"github.com/wodadehencou/fakedata"
)

func TestMobileNumberGenerator(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := fakedata.MobileNumber()
		re := regexp.MustCompile(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
		t.Log(n)
		if !re.MatchString(n) {
			t.Errorf("number %s is not valid", n)
		}
	}

}
