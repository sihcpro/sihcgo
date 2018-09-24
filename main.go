package main

import (
	"fmt"

	"github.com/sihcgo/treeword"
)

func main() {
	a := treeword.New()

	a.Inserts("angular", "asdasdasds", "asddsa")
	a.Inserts("aws")
	a.Inserts("androi")
	a.Inserts("atal bihari vajpayee")
	a.Inserts("aretha franklin")
	a.Inserts("angularjs")
	a.Inserts("ariana grande")
	a.Inserts("amazon")
	a.Inserts("anime")
	a.Inserts("apple")
	a.Inserts("angry bird")
	a.Inserts("animals")
	a.Inserts("angel")
	a.Inserts("anna kendrick")
	a.Inserts("angelina jolie")
	a.Inserts("angela lang")
	a.Inserts("and here we go")
	a.Inserts("amber rose")
	a.Inserts("golang")
	a.Inserts("i love you")
	a.Inserts("anh Phuc")
	a.Inserts("anh Phuc")
	a.Inserts("anh Phuc")
	a.Inserts("anh Phuc")
	a.Inserts("anh Phuc")
	a.Inserts("anh Phuc")
	a.Inserts("annnn")

	// a.Print()
	fmt.Println(a.All())
}
