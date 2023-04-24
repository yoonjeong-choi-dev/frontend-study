package mocking

import "errors"

var ThrowError = func() error {
	return errors.New("always failed")
}

func DoSomeStuff(d DoStuffer) error {
	if err := d.DoStuff("test"); err != nil {
		return err
	}

	// 항상 실패
	if err := ThrowError(); err != nil {
		return err
	}

	return nil
}
