// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type FileContent struct {
}

func NewFileContent() *FileContent {
	file := &FileContent{}
	return file
}

func (self *FileContent) IsDirectory() bool {
	return false
}
