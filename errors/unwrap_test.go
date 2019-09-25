package unwrap_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var err = errors.New("new errors")

func TestUnwrap(t *testing.T) {
	if innerErr := errors.Unwrap(fmt.Errorf("%w", err)); innerErr != err {
		t.Error(`errors.Unwrap(fmt.Errorf("%w", err))!=err`)
	} else {
		t.Log(`fmt.Errorf("%w").Unwrap()`)
	}
}

func TestIs(t *testing.T) {
	if errors.Is(fmt.Errorf("%w", err), err) {
		t.Log(`errors.Is(fmt.Errorf("%w", err), err) == true`)
	} else {
		t.Error(`errors.Is(fmt.Errorf("%w", err), err) == false`)
	}
}
func TestIsInMultiWrap(t *testing.T) {
	if errors.Is(fmt.Errorf("%w", fmt.Errorf("%w", err)), err) {
		t.Log(`errors.Is(fmt.Errorf("%w", fmt.Errorf("%w", err)), err)`)
	} else {
		t.Error(`errors.Is(fmt.Errorf("%w", fmt.Errorf("%w", err)), err) == false`)
	}
}

func TestAsAsign(t *testing.T) {
	var err error
	if errors.As(fmt.Errorf("%w", err), &err) {
		t.Log("pass")
	} else {
		t.Error("error")
	}
}

func TestAsWithNilPointer(t *testing.T) {
	var pathError **os.PathError
	if _, err := os.Open("not-exist"); errors.As(err, pathError) {
		t.Logf("%s", *pathError)
	} else {
		t.Error("")
	}
}

func TestAsWithErrorInterface(t *testing.T) {
	var plainError error
	if _, err := os.Open("not-exist"); errors.As(err, &plainError) {
		t.Logf("%s", plainError)
	} else {
		t.Fatal("error")
	}
}

func TestAsWithInterface(t *testing.T) {
	var notError interface {
		NotError()
		Error() string
	}
	if _, err := os.Open("not-exist"); errors.As(err, &notError) {
		t.Logf("%s", err)
	} else {
		t.Fatal("error")
	}
}
