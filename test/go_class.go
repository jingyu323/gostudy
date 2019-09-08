package main

import "fmt"

/**
golang用另外一种做法来实现属性的访问权限：
属性的开头字母是大写的则在其它包中可以被访问，否则只能在本包中访问。类的声明和方法亦是如此
*/
type Poem struct {
	Title  string
	Author string
	intro  string
}

func (poem *Poem) publish() {
	fmt.Println("poem publish")
}
