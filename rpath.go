package rpath

import (
	"path"
)

func IsDir(p string) bool {
	if p == "" {
		return false
	}
	return p[len(p)-1] == '/'
}

func SetFile(p string) string {
	p = path.Clean(p)
	if len(p) > 0 && p[len(p)-1] == '/' {
		return p[:len(p)-1]
	}
	return p
}

func SetDir(p string) string {
	p = path.Clean(p)
	if len(p) > 0 && p[len(p)-1] != '/' {
		return p + "/"
	}
	return p
}

func SetType(p string, is_dir bool) string {
	if is_dir {
		return SetDir(p)
	}
	return SetFile(p)
}

func Base(p string) string {
	is_dir := IsDir(p)
	return SetType(path.Base(p), is_dir)
}

func Clean(p string) string {
	is_dir := IsDir(p)
	return SetType(path.Clean(p), is_dir)
}

func Dir(p string) string {
	return SetDir(path.Dir(p))
}

func Ext(p string) string {
	if IsDir(p) {
		return ""
	}
	return path.Ext(p)
}

var IsAbs = path.IsAbs

func Split(p string) (string, string) {
	is_dir := IsDir(p)
	d, n := path.Split(p)
	return SetDir(d), SetType(path.Clean(n), is_dir)
}

func Match(pat, n string) (bool, error) {
	n = Clean(n)
	return path.Match(pat, n)
}

func Join(pe ...string) string {
	if len(pe) == 0 {
		return ""
	}
	return SetType(path.Join(pe...), IsDir(pe[len(pe)-1]))
}
