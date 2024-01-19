package macmap

import "C"
import (
	"fmt"
	"os"
	"strings"
)

var db24 map[string]interface{}
var db28 map[string]interface{}
var db36 map[string]interface{}

func init() {
	f24Name := "./MAS.csv"
	f28Name := "./MAM.csv"
	f36Name := "./MAL.csv"

	f24, _ := os.Open(f24Name)
	f28, _ := os.Open(f28Name)
	f36, _ := os.Open(f36Name)
	defer f24.Close()
	defer f28.Close()
	defer f36.Close()

	reader24 := NewReader(f24)
	reader28 := NewReader(f28)
	reader36 := NewReader(f36)

	_, _ = reader24.GetFieldNames()

	db24, _ = reader24.ReadAll2Map("Assignment")
	db28, _ = reader28.ReadAll2Map("Assignment")
	db36, _ = reader36.ReadAll2Map("Assignment")
}

func Search(mac string) (info string) {
	var bit24, bit28, bit36 int = 24, 28, 36
	strSlice := strings.Split(mac, ":")
	macStr := strings.Join(strSlice, "")

	index24 := strings.ToUpper(macStr[0 : bit24/4])
	index28 := strings.ToUpper(macStr[0 : bit28/4])
	index36 := strings.ToUpper(macStr[0 : bit36/4])

	var vendorInfo interface{} = nil
	if info1, ok := db24[index24]; ok {
		if info2, ok := db28[index28]; ok {
			if info3, ok := db36[index36]; ok {
				vendorInfo = info3
			}
			vendorInfo = info2
		}
		vendorInfo = info1
	}

	if vendorInfo == nil {
		return ""
	}
	vendorInfoSlice := make([]string, 0)
	vendorInfoSlice = vendorInfo.([]string)

	return fmt.Sprintf("%s = %s", vendorInfoSlice[0], vendorInfoSlice[2])
}
