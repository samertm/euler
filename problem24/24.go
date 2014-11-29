package problem24

import (
	"fmt"
	"log"

	"github.com/fighterlyt/permutation"
)

func Solve() string {
	perms := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	p, err := permutation.NewPerm(perms, nil)
	if err != nil {
		log.Fatal(err)
	}
	var thingy interface{}
	for i := 0; i < 1000000; i++ {
		thingy, _ = p.Next()
	}
	return fmt.Sprintf("%v", thingy)
}
