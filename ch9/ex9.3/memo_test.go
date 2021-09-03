// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package ex9_3_test

import (
	ex9_3 "gopl.io/ch9/ex9.3"
	"testing"
	"time"
)

var httpGetBody = ex9_3.HTTPGetBody

func Test(t *testing.T) {
	var done = make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	m := ex9_3.New(httpGetBody)
	ex9_3.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	var done = make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	m := ex9_3.New(httpGetBody)
	ex9_3.Concurrent(t, m, done)
}
