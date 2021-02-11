package wraperr

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapErr(t *testing.T) {
	e1 := errors.New("Error 1")
	e2 := Wrapf(e1, "Error 2")
	e3 := Wrapf(e2, "Error 3")
	if Contains(e3, "Error 1") == false {
		t.Fatalf("Not wrapped propperly")
	}
	if Contains(e2, "Error 1") == false {
		t.Fatalf("Not wrapped propperly")
	}
	fmt.Println(e3)
	fmt.Println(e2)
	fmt.Println(e1)
}
