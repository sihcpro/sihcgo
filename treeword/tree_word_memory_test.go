package treeword

import(
	"os"
	"log"
	"bufio"
	"fmt"
	"time"
	"runtime"
	"testing"
)

func compareMemorySize(tmp []string, loop int) (int, int) {
	sizeArrayString := 0
	sizeTreeWord := 0
	for i:= 0; i < loop; i++ {
		tmp2 := make([]string, 0)
		for _, s := range tmp {
			tmp2 = append(tmp2, s)
		}
		sizeArrayString += getSizeAlloc()
		printMemUsage()
	}

	for i:= 0; i < loop; i++ {
		tmp2 := New()
		tmp2.Inserts(tmp...)
		sizeTreeWord += getSizeAlloc()
		printMemUsage()
	}

	fmt.Println("Average size of ", len(tmp), "names : ")
	fmt.Println("Array String :", sizeArrayString/loop/1000, "KB")
	fmt.Println("TreeWord     :",    sizeTreeWord/loop/1000, "KB")

	return sizeArrayString, sizeTreeWord
}

func Test_Memory_Size(t *testing.T) {
	loop := 3

	printMemUsage()
	tmp := read_file()
	fmt.Println("Readed :", len(tmp), " names")

	arrSize, treewordSize := compareMemorySize(tmp,loop)

	if arrSize <= treewordSize {
		t.Error("Not good!")
	}
}

func compareTimeInsert(tmp []string, loop int) (int, int) {
	timeArrayString := 0
	timeTreeWord := 0

	for i := 0; i < loop; i++ {
		tmpT := time.Now()
		tmp2 := make([]string, 0)
		for _, s := range tmp {
			tmp2 = append(tmp2, s)
		}
		timeArrayString += int(time.Since(tmpT).Nanoseconds())
	}

	for i := 0; i < loop; i++ {
		tmpT := time.Now()
		tmp2 := New()
		tmp2.Inserts(tmp...)
		timeTreeWord += int(time.Since(tmpT).Nanoseconds())
	}

	fmt.Println("Array String :", timeArrayString/1000, " js")
	fmt.Println("TreeWord     :",    timeTreeWord/1000, " js")

	return 0,0
}

func Test_Time_Insert(t *testing.T) {
	loop := 3

	tmp := read_file()
	compareTimeInsert(tmp, loop)

	t.Log("ok")
}



func Test_Time_Get_All(t *testing.T) {

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
