package treeword

import(
	"os"
	"log"
	"bufio"
	"fmt"
	"runtime"
	"testing"
)

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

func test_memory_with_string() []string {
	tmp := read_file()
	// PrintMemUsage()
	fmt.Println("Readed :", len(tmp), " names")
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
// 	PrintMemUsage()
// 	test_memory_with_tree_word(tmp)
// 	PrintMemUsage()

// 	fmt.Println(len(tmp))
// }


func TestMain(t *testing.T) {
	// fmt.Println(len(read_file()))
	tmp := test_memory_with_string()
	PrintMemUsage()
	test_memory_with_tree_word(tmp)
	PrintMemUsage()

	fmt.Println(len(tmp))
	t.Log("ok")
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
