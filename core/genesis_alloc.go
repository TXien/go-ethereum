// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8\xf5\xe2\x94/\xc2|{\x0f\x0f\x8e\xfby\x9a\x15\x14N\x89\x14\x8c\u0529\x99o\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\xe2\x94T\xba\xb0\xe2\x9eQ\x92\x9f\x93\xb5\xab\x0e\x87\x1b\nOI\x81o_\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\xe2\x94i\x9a\a\u0774\xe5\xef\xeb\u822c(\x10}b\xf3\xfc\xda#\x9c\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\u250a\xe3$&\x06^\x89\x01\x86\xec\x9b9\xf18/\xbc\xff\xf7\x19#\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\u2519\x136%\x05\xb7+\xed\xa6*\x9fE\xcd*j\xe9\xa9\xd5s\xb4\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\u252f\xb7\xcc\xfem\x91{\xdf\xf8\xafl\xc4f\xeb\x1ft\xfe^0?\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00\xe2\x94\xeb%3\x9dCHBW\x9b!\xff}W\xcd\a\x8b\x98\x8a=\r\x8c\x03;.<\x9f\u0400<\xe8\x00\x00\x00"
const testnetAllocData = "\xe3\xe2\x94\xe7\f.\xe3\f\xbfq\xf9-l;\xc6\x15?\xa5\x88\x04k\x84\u05cc\x16\x9eC\xa8^\xb3\x81\xaaX\x00\x00\x00"
const rinkebyAllocData = "\xe3\xe2\x94\xe7\f.\xe3\f\xbfq\xf9-l;\xc6\x15?\xa5\x88\x04k\x84\u05cc\x16\x9eC\xa8^\xb3\x81\xaaX\x00\x00\x00"
