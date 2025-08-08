// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Directory struct {
}

func NewDirectory() *Directory {
	dir := &Directory{}
	return dir
}

func (dir *Directory) IsDirectory() bool {
	return true
}
