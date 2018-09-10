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

func (this *TreeWord) insert(char byte, pos... int) {
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
	// fmt.Println()
	// fmt.Print("-> ")
	if len(s) == 0 {
		this.Amount++
		return this.Amount
	}
	char := s[0]
	// fmt.Print(" ", string(char))
	// if this.Char == byte(0) {
		tmp := search.BinarySearchB(this.NextChar, char)
		// fmt.Print(" ", tmp, " vs ", len(this.NextChar))
// 
		if tmp == -1 {
			// this.NextChar = append([]byte{char}, this.NextChar...)
			// var newWord *TreeWord
			// if len(s) == 1 {
			// 	// newWord = Construct(char, 1, []byte{}, []*TreeWord{})
			// 	// this.NextTree = append([]*TreeWord{newWord}, this.NextTree...)
			// 	this.insert(char, 1)
			// 	return 1
			// } else {
				// newWord = Construct(char, 0, []byte{}, []*TreeWord{})
				// this.NextTree = append([]*TreeWord{newWord}, this.NextTree...)
				this.insert(char)
				return this.NextTree[0].Insert(s[1:])
			// }
		} else if this.NextChar[tmp] == char {
			// this.NextTree[tmp].Amount++
			return this.NextTree[tmp].Insert(s[1:])
		} else {
			// tmpNextChar := append(this.NextChar[:tmp], char)
			// if tmp < len(this.NextChar) {
			// 	this.NextChar = append(tmpNextChar, this.NextChar[tmp:]...)
			// }
			// if len(s) == 1 {
			// 	this.insert(char, 1, tmp)
			// 	return 1
			// } else {
				this.insert(char, tmp)
				return this.NextTree[tmp+1].Insert(s[1:])
			// }
		}

		// this.NextTree[]
	// } else {
	// 	fmt.Print(" Not yet! ")
	// }
	return 0
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
				this.NextTree[i].print(head+"   |", tmp2)
			}
		} else {
			fmt.Print(head, "   |")
			fmt.Println()
			fmt.Print(head, "   ", string(c))
			if i < len(this.NextChar)-1 {
				this.NextTree[i].print(head+"   |", tmp2)
			} else {
				this.NextTree[i].print(head+"    ", tmp2)
			}
		}
		// fmt.Println("2")
	}
}

func (this *TreeWord) Print() {
	this.print("", "")
}

func (this *TreeWord) helper(s string) {
	this.Insert(s)
	// fmt.Print( " : ", s, " >> ", this.Insert(s))
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

