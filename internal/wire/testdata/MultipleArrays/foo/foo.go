// Copyright 2025 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(inject())
	fmt.Println(injectWithGiven([]*Bar{
		{A: "gb.0"},
		{A: "gb.1"},
	}))
	fmt.Println(injectSingle())
}

type Foo struct {
	B    []*Bar
	baaz Baaz
}

func (f Foo) String() string {
	return fmt.Sprintf("{B %s baaz %s}", f.B, f.baaz)
}

type Bar struct {
	A string
}

func (b Bar) String() string {
	return b.A
}

type Baaz struct {
	bars []*Bar
}

func (b Baaz) String() string {
	return fmt.Sprintf("{bars %s}", b.bars)
}

func provideFoo(baaz *Baaz, b ...*Bar) *Foo {
	return &Foo{B: b, baaz: *baaz}
}

func provideBars() []*Bar {
	return []*Bar{
		{A: "pb.0.0"},
		{A: "pb.1.0 pb.1.1"},
	}
}

func provideBarsAgain() []*Bar {
	return []*Bar{
		{A: "pba.0.0"},
		{A: "pba.1.0 pba.1.1"},
	}
}
