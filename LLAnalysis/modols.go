package main

import "strings"

type Edge struct {
	left     string
	right    string
	rlen     int
	first    string
	follow   string
	myselect string
}

func (e *Edge) getlf() string {
	return e.left
}

func (e *Edge) getrg() string {
	return e.right
}

func (e *Edge) getfirst() string {
	return e.first
}

func (e *Edge) getfollow() string {
	return e.follow
}

func (e *Edge) getro() string {
	str := ""
	str += string(e.right[0])
	return str
}

func (e *Edge) getselect() string {
	return e.myselect
}

func (e *Edge) getrlen() int {
	return e.rlen
}
func (e *Edge) newfirst(str string) {
	var i int
	for i = 0; i < len(str); i++ {
		strtemp :=string(str[i])
		if !strings.Contains(e.first, strtemp) {
			e.first += strtemp
		}
	}
}
func (e *Edge) newselect(str string) {
	var i int
	for i = 0; i < len(str); i++ {
		strtemp :=string(str[i])
		if !strings.Contains(e.myselect, strtemp)&&strtemp!="@" {
			e.myselect += strtemp
		}
	}
}
func (e *Edge) newfollow(str string) {
	var i int
	for i = 0; i < len(str); i++ {
		strtemp :=string(str[i])
		if !strings.Contains(e.follow, strtemp)&&strtemp!="@" {
			e.follow += strtemp
		}
	}
}
func (e *Edge) delfirst() {
	index := strings.Index(e.first, "@")
	tem := []rune(e.first)
	for i := 0; i < len(tem); i++ {
		if i==index {
			tem = append(tem[:i], tem[i+1:len(tem)]...)
		}
	}
	e.first = string(tem)
}
