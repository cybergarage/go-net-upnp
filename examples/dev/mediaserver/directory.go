// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

type Directory struct {
}

func NewDirectory() *Directory {
	dir := &Directory{}
	return dir
}

func (self *Directory) IsDirectory() bool {
	return true
}
