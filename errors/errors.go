// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type WrappedError struct {
	message string
	wrapped error
	caller  xerrors.Frame
}

var _ error = &WrappedError{}
var _ xerrors.Wrapper = &WrappedError{}

func (err *WrappedError) Is(x error) bool {
	return xerrors.Is(err, x)
}

func (err *WrappedError) CausedBy(x error) bool {
	if err.wrapped == x {
		return true
	}

	if typedWrapped, ok := err.wrapped.(*WrappedError); ok {
		if typedWrapped.Is(x) || typedWrapped.CausedBy(x) {
			return true
		}
	}

	return false
}

func (err *WrappedError) Unwrap() error {
	return err.wrapped
}

func (err *WrappedError) Error() string {
	if err.wrapped != nil {
		return fmt.Sprintf("%s: %s", err.message, err.wrapped.Error())
	}
	return err.message
}

func (err *WrappedError) Format(f fmt.State, c rune) { // implements fmt.Formatter
	xerrors.FormatError(err, f, c)
}

func (err *WrappedError) FormatError(p xerrors.Printer) error { // implements xerrors.Formatter
	p.Print(err.message)
	if p.Detail() {
		err.caller.Format(p)
	}
	return err.wrapped
}

func New(msg string) error {
	return newWrappedError(msg, nil)
}

func newWrappedError(msg string, err error) *WrappedError {
	return &WrappedError{
		message: msg,
		wrapped: err,
		caller:  xerrors.Caller(2),
	}
}

func Wrap(msg string, err error) error {
	return newWrappedError(msg, err)
}

func Must(x ...interface{}) {
	lastValue := x[len(x)-1]
	if lastValue == nil {
		return
	}
	if wrappedErr, ok := lastValue.(WrappedError); ok {
		panic(fmt.Sprintf("%+v", wrappedErr))
	}
	if err, ok := lastValue.(error); ok {
		panic(err)
	}
}
