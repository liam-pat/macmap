package macmap_test

import (
	"fmt"
	"github.com/YaoMiss/macmap"
	"testing"
)

func TestSearch(t *testing.T) {
	mac1 := "48:e2:44:45:0b:04"
	d1 := macmap.Search(mac1)
	fmt.Printf("%s ==> %v\n", mac1, d1)

}
