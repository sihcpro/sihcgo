package treeword

import(
	"os"
	"log"
	"bufio"
	"fmt"
	"runtime"
	"testing"
)

func compareMemorySize(tmp []string, loop int) (int, int) {
	sizeArrayString := 0
	sizeTreeWord := 0
	for i:= 0; i < loop; i++ {
		test_memory_with_string(tmp)
		sizeArrayString += getSizeAlloc()
		printMemUsage()
	}
	for i:= 0; i < loop; i++ {
		test_memory_with_tree_word(tmp)
		sizeTreeWord += getSizeAlloc()
		printMemUsage()
	}

	fmt.Println("Average size of ", len(tmp), "names : ")
	fmt.Println("Array String :", sizeArrayString/loop/1000, "KB")
	fmt.Println("TreeWord     :", sizeTreeWord/loop/1000, "KB")

	return 0, 0
}

func TestTreeword(t *testing.T) {
	// fmt.Println(len(read_file()))

	loop := 3

	printMemUsage()
	tmp := read_file()
	fmt.Println("Readed :", len(tmp), " names")

	compareMemorySize(tmp,loop)

	// fmt.Println(len(tmp))
	t.Log("ok")
	if false {
		t.Error(
			"Not good!",
		)
	}
}

func read_file() []string {
	file, err := os.Open("name_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tmp := make([]string, 0)
	for scanner.Scan() {
		tmp = append(tmp, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tmp
}

func test_memory_with_string(arrStr []string) []string {
	tmp := make([]string, 0)
	for _, s := range arrStr {
		tmp = append(tmp, s)
	}
	return tmp
}

func test_memory_with_tree_word(arrStr []string) *TreeWord {
	tmp := New()
	for _, str := range arrStr {
		tmp.Insert(str)
	}
	return tmp
}

// func main() {
// 	// fmt.Println(len(read_file()))
// 	tmp := test_memory_with_string()
// 	printMemUsage()
// 	test_memory_with_tree_word(tmp)
// 	printMemUsage()

// 	fmt.Println(len(tmp))
// }

var (
	m runtime.MemStats
	oldTotalAlloc = 0
	tmpTotalAlloc = 0
)
func printMemUsage() {
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func getSizeAlloc() int {
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	tmpTotalAlloc = int(m.TotalAlloc) - oldTotalAlloc
	oldTotalAlloc = int(m.TotalAlloc)
	return tmpTotalAlloc
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
