package treeword

import(
	"fmt"
	"time"
	"github.com/sihcgo/search"
)

type TreeWord struct {
	Char     byte
	Amount   int
	NextChar []byte
	NextTree []*TreeWord
}

func New() *TreeWord {
	tmp := new(TreeWord)
	tmp.Char = 0
	tmp.Amount = 0
	return tmp
}

func Construct(b byte, a int, nl []byte, nc []*TreeWord) *TreeWord {
	tmp := new(TreeWord)
	tmp.Char = b
	tmp.Amount = a
	tmp.NextChar = nl
	tmp.NextTree = nc
	return tmp
}

func (this *TreeWord) append(char byte, pos... int) {
	// fmt.Print(" [", string(this.NextChar), "]")
	if len(pos) > 0 {
		p := pos[0]+1
		// fmt.Print(" [", string(this.NextChar[:p]), "]")
		tree := Construct(char, 0, []byte{}, []*TreeWord{})

		tmpNextChar := append([]byte{char},      this.NextChar[p:]...)
		tmpNextTree := append([]*TreeWord{tree}, this.NextTree[p:]...)
		if p-1 < len(this.NextChar) {
			this.NextChar = append(this.NextChar[:p], tmpNextChar...)
			this.NextTree = append(this.NextTree[:p], tmpNextTree...)
		}
	} else {
		tree := Construct(char, 0, []byte{}, []*TreeWord{})

		this.NextChar = append([]byte{char},      this.NextChar...)
		this.NextTree = append([]*TreeWord{tree}, this.NextTree...)
	}
	// fmt.Print(" -> [", string(this.NextChar), "] ")
}

func (this *TreeWord) Insert(s string) int {
	if len(s) == 0 {
		this.Amount++
		return this.Amount
	}
	char := s[0]
	tmp := search.BinarySearchB(this.NextChar, char)
	if tmp == -1 {
		this.append(char)
		return this.NextTree[0].Insert(s[1:])
	} else if this.NextChar[tmp] == char {
		return this.NextTree[tmp].Insert(s[1:])
	} else {
		this.append(char, tmp)
		return this.NextTree[tmp+1].Insert(s[1:])
	}

	return 0
}

func (this *TreeWord) Inserts(words... string) {
	for _, word := range words {
		this.Insert(word)
	}
}

func (this *TreeWord) print(head, tmp string) {
	st := 0
	tmp2 := tmp + string(this.Char)
	if this.Amount > 0 {
		st = -1
		fmt.Print( " ~~>> ", tmp2, " [", this.Amount, "]")
		fmt.Println()
	}
	for i, c := range this.NextChar {
		if i == st {
			fmt.Print(" ~ ", string(c))
			if len(this.NextChar) == 1 {
				this.NextTree[i].print(head+"    ", tmp2)
			} else {
				this.NextTree[i].print(head+"|   ", tmp2)
			}
		} else {
			fmt.Print(head, "|  ")
			fmt.Println()
			fmt.Print(head, "|~~ ", string(c))
			if i < len(this.NextChar)-1 {
				this.NextTree[i].print(head+"|   ", tmp2)
			} else {
				this.NextTree[i].print(head+"    ", tmp2)
			}
		}
		// fmt.Println("2")
	}
}

func (this *TreeWord) all(arr *[]string, tmp string, distinct bool) {
	tmp2 := tmp+string(this.Char)
	// if this.Amount > 0 {
	// 	*arr = append(*arr, tmp2)
	// }
	if distinct && this.Amount > 0 {
		*arr = append(*arr, tmp2)
	} else {
		for i := 0; i < this.Amount; i++ {
			*arr = append(*arr, tmp2)
		}
	}
	for _, c := range this.NextTree {
		c.all( arr, tmp2, distinct)
	}
}

func (this *TreeWord) All() []string {
	arr := new([]string)
	this.all(arr, "", false)
	return *arr
}

func (this *TreeWord) AllDistinct() []string {
	arr := new([]string)
	this.all(arr, "", true)
	return *arr
}

func (this *TreeWord) AllChild() []string {
	arr := new([]string)
	for _, child := range this.NextTree {
		*arr = append(*arr, child.AllDistinct()...)
	}
	return *arr
}


func (this *TreeWord) Print() {
	fmt.Print(" ")
	this.print("", "")
	fmt.Println()
}

func (this *TreeWord) helper(s string) {
	this.Insert(s)
	// fmt.Print( " : ", s, " >> ", this.Insert(s))
}

func PrintString(arrString []string) {
	fmt.Print("[")
	for i, s := range arrString {
		// fmt.Print("\"",s,"\"")
		fmt.Print(s)
		if i < len(arrString)-1 {
			fmt.Print(";")
		}
	}
	fmt.Print("]\n")
}

func (this *TreeWord) find(findPrev *TreeWord, arrResult *[]string, tmpStr string) {
	findAll, ok := findPrev.GetTree('*')
	// find := Copy(findPrev)
	var allNextFind []string
	if ok {
		allNextFind = findAll.AllDistinct()
		reg2 := New()
		reg2.Inserts(allNextFind...)
		reg2.Print()
		reg2.Inserts(findAll.AllChild()...)
		reg2.Print()
		allNextFind = reg2.AllDistinct()
	}

	// tmp := find.AllDistinct()
	// for _, findString := range tmp {
	// 	if len(findString) > 1 {
	// 		nextFind := findString[1:]
	// 		num := find.Insert(nextFind)
	// 		if num == 1 {
	// 			// firstInsert
	// 			// this.GetTree
	// 		}
	// 	}
	// }

	fmt.Println("len = ", len(this.NextChar))

	tT, tF := &(this.NextChar), &(findPrev.NextChar)
	lT, lF := len(*tT), len(*tF)
	fmt.Print("Find")
	PrintString(findPrev.AllChild())
	fmt.Println("In  ")
	PrintString(this.AllChild())

	for cT, cF := 0, 0; cT < lT && cF < lF; {
		fmt.Print(cT, "/", lT, " & ", cF, "/", lF, " ", (*tT)[cT], (*tF)[cF],
			" '", string((*tT)[cT]), "' '",string((*tF)[cF]), "'\n")
		tmpNextStr := tmpStr + string(this.Char)
		if (*tF)[cF] == '*' {
			if findPrev.NextTree[cF].Amount > 0 && this.NextTree[cT].Amount > 0 {
				*arrResult = append(*arrResult, tmpNextStr + string(this.NextTree[cT].Char))
			}
			cF++
			continue
		}

		if (*tT)[cT] == (*tF)[cF] {

			if this.NextTree[cT].Amount > 0 && findPrev.NextTree[cF].Amount > 0 {
				*arrResult = append(*arrResult, tmpNextStr + string(this.NextTree[cT].Char))
			}

			nextFindTree := New()
			nextFindTree.Inserts(allNextFind...)
			nextFindTree.Inserts((findPrev.NextTree[cF]).AllChild()...)
			nextFindTree.Print()
			this.NextTree[cT].Print()
			this.NextTree[cT].find(nextFindTree, arrResult, tmpNextStr)
			cT++
			cF++
		} else if (*tT)[cT] > (*tF)[cF] {
			cF++
		} else {
			// this.NextTree[cT].find()
			cT++
		}
		fmt.Print(cT, "/", lT, " & ", cF, "/", lF, "\n")
	}
}

func (this *TreeWord) Find(regex string) []string {
	findTree := New()
	arrResult := new([]string)
	findTree.Insert(regex)
	// findTree.Insert("abx")
	this.find(findTree, arrResult, "")
	return *arrResult
}

func (this *TreeWord) GetTree(char byte) (*TreeWord, bool) {
	pos := search.BinarySearchB(this.NextChar, char)
	if pos  == -1 {
		return New(), false
	} else {
		return this.NextTree[pos], true
	}
}

func Copy(tree *TreeWord) *TreeWord {
	newTree := New()
	newTree.Inserts(tree.All()...)
	return newTree
}

func main2() {
	startTime := time.Now()
	a := map[byte]int {
		0 : 0,
		// 1 : 1,
		// 3 : 3,
		// 2 : 2,
	}
	b := []int {
		1, 2, 3, 4,
	}
	var s = 0
	// for i:= 0; i < 100000; i++ {
	// 	// for _, k := range a {
	// 	// 	s += k
	// 	// }
	// 	// s+= b[0]
	// 	b = append(b, i)
	// 	b = append(b[0:len(b)/2], b[(len(b)/2):len(b)]...)
	// 	// fmt.Println(len(b))
	// }

	fmt.Println(s, a[0], b[0])

	fmt.Println("---------------------")

	c := New()
	c.helper("angular")
	c.helper("aws")
	c.helper("androi")
	c.helper("atal bihari vajpayee")
	c.helper("aretha franklin")
	c.helper("angularjs")
	c.helper("ariana grande")
	c.helper("amazon")
	c.helper("anime")
	c.helper("apple")
	c.helper("angry bird")
	c.helper("animals")
	c.helper("angel")
	c.helper("anna kendrick")
	c.helper("angelina jolie")
	c.helper("angela lang")
	c.helper("and here we go")
	c.helper("amber rose")
	c.helper("golang")
	c.helper("i love you")
	c.helper("anh Phuc")
	c.helper("anh Phuc")
	c.helper("anh Phuc")
	c.helper("anh Phuc")
	c.helper("anh Phuc")
	c.helper("anh Phuc")
	c.Print()

	fmt.Println("\n~~~ END ~~~", time.Since(startTime))
}

func main() {
	for i:= 0; i < 1; i++ {
		fmt.Println("Test ", i+1)
		main2()
	}
}

