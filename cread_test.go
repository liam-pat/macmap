package macmap_test

import (
	"fmt"
	"github.com/YaoMiss/macmap"
	"os"
	"strings"
	"testing"
)

func TestMALTable(t *testing.T) {
	fName := "./db/MAL.csv"
	f, _ := os.Open(fName)
	defer f.Close()
	reader := macmap.NewReader(f)

	records, _ := reader.ReadAll()
	records2 := records[:5]
	fmt.Println("len :", len(records))
	for _, value := range records2 {
		fmt.Println(value)
	}
	fmt.Println(strings.Repeat("%", 30))
}
func TestMAMTable(t *testing.T) {
	fName := "./db/MAM.csv"
	f2, _ := os.Open(fName)
	defer f2.Close()

	reader2 := macmap.NewReader(f2)
	reader2.SetSkip(1)
	blockRecords, _ := reader2.ReadBlock()

	fmt.Println("len :", len(blockRecords))
	for _, value := range blockRecords[:5] {
		fmt.Println(value)
	}
	fmt.Println(blockRecords[len(blockRecords)-1])
	fmt.Println(strings.Repeat("%", 30))
}

func TestMASTable(t *testing.T) {
	fName := "./db/MAS.csv"
	f, _ := os.Open(fName)
	defer f.Close()

	reader := macmap.NewReader(f)
	_, _ = reader.GetFieldNames()
	records, _ := reader.ReadAll2Map("Assignment")

	fmt.Println("len :", len(records))
	var i = 0
	for key, value := range records {
		if i == 5 {
			break
		}
		fmt.Printf("key = %s => %v \n", key, value)
		i++
	}
	fmt.Println(strings.Repeat("%", 30))
}
