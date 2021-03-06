// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors

// New returns an error that formats as the given text.
// 其实就是error就是string的固定
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
// errorString 是简单的简单实现
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
