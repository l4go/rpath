package rpath_test

import (
	"fmt"

	"github.com/l4go/rpath"
)

func ExampleIsDir() {
	fmt.Println(rpath.IsDir("foo/"))
	fmt.Println(rpath.IsDir("foo"))
	fmt.Println(rpath.IsDir("/foo/"))
	fmt.Println(rpath.IsDir("/foo"))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleSetFile() {
	fmt.Println(rpath.SetFile("/foo"))
	fmt.Println(rpath.SetFile("/foo/"))

	// Output:
	// /foo
	// /foo
}

func ExampleSetDir() {
	fmt.Println(rpath.SetDir("/foo"))
	fmt.Println(rpath.SetDir("/foo/"))

	// Output:
	// /foo/
	// /foo/
}

func ExampleSetType() {
	fmt.Println(rpath.SetType("/foo", true))
	fmt.Println(rpath.SetType("/foo", false))
	fmt.Println(rpath.SetType("/foo/", true))
	fmt.Println(rpath.SetType("/foo/", false))

	// Output:
	// /foo/
	// /foo
	// /foo/
	// /foo
}

func ExampleBase() {
	fmt.Println(rpath.Base("/foo"))
	fmt.Println(rpath.Base("/foo/"))
	fmt.Println(rpath.Base("/foo/bar"))
	fmt.Println(rpath.Base("/foo/bar/"))

	// Output:
	// foo
	// foo/
	// bar
	// bar/
}

func ExampleClean() {
	fmt.Println(rpath.Clean("foo/../bar"))
	fmt.Println(rpath.Clean("foo/../bar/"))
	fmt.Println(rpath.Clean("foo/../../bar"))
	fmt.Println(rpath.Clean("foo/../../bar/"))
	fmt.Println(rpath.Clean("/foo/../../bar"))
	fmt.Println(rpath.Clean("/foo/../../bar/"))

	// Output:
	// bar
	// bar/
	// ../bar
	// ../bar/
	// /bar
	// /bar/
}

func ExampleDir() {
	fmt.Println(rpath.Dir("/foo"))
	fmt.Println(rpath.Dir("/foo/"))
	fmt.Println(rpath.Dir("/foo/hoge.txt"))
	fmt.Println(rpath.Dir("/foo/bar"))
	fmt.Println(rpath.Dir("/foo/bar/"))
	fmt.Println(rpath.Dir("/foo/bar/hoge.txt"))

	// Output:
	// /
	// /foo/
	// /foo/
	// /foo/
	// /foo/bar/
	// /foo/bar/
}

func ExampleExt() {
	fmt.Println(rpath.Ext("foo/bar"))
	fmt.Println(rpath.Ext("foo/bar/"))
	fmt.Println(rpath.Ext("foo/bar/baz.txt"))
	fmt.Println(rpath.Ext("foo/bar.ext/"))

	// Output:
	//
	//
	// .txt
	//
}

func ExampleSplit() {
	fmt.Println(rpath.Split("/foo/bar"))
	fmt.Println(rpath.Split("/foo/bar/"))
	fmt.Println(rpath.Split("/foo/bar/baz.txt"))
	fmt.Println(rpath.Split("/foo/bar.ext/"))

	// Output:
	// /foo/ bar
	// /foo/bar/ ./
	// /foo/bar/ baz.txt
	// /foo/bar.ext/ ./
}

func ExampleMatch() {
	var m bool
	var err error
	m, err = rpath.Match("f*o", "foo")
	fmt.Println(m, err == nil)
	m, err = rpath.Match("f*o", "foo/")
	fmt.Println(m, err == nil)
	m, err = rpath.Match("f*o/", "foo/")
	fmt.Println(m, err == nil)
	m, err = rpath.Match("f*o", "/bar/foo")
	fmt.Println(m, err == nil)
	m, err = rpath.Match("f*o[", "/bar/foo")
	fmt.Println(m, err == nil)

	// Output:
	// true true
	// false true
	// true true
	// false true
	// false false
}

func ExampleJoin() {
	fmt.Println(rpath.Join("/foo", "bar"))
	fmt.Println(rpath.Join("/foo/", "bar/"))
	fmt.Println(rpath.Join("foo", "bar", "baz.txt"))
	fmt.Println(rpath.Join("foo", "bar", "baz/"))

	// Output:
	// /foo/bar
	// /foo/bar/
	// foo/bar/baz.txt
	// foo/bar/baz/
}
