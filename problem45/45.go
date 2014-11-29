package problem45

import (
	"fmt"
	"log"

	"github.com/samertm/euler/math"
)

func Solve() string {
	triangleNumsCh := make(chan math.NVal)
	go math.GenNVals(nil, triangleNumsCh, math.TriangleN, nil)
	pentagonalNumsCh := make(chan math.NVal)
	go math.GenNVals(nil, pentagonalNumsCh, math.PentagonalN, nil)
	hexagonalNumsCh := make(chan math.NVal)
	go math.GenNVals(nil, hexagonalNumsCh, math.HexagonalN, nil)
	var tri math.NVal
	tri = <-triangleNumsCh
	for tri.Val <= 40755 {
		tri = <-triangleNumsCh
	}
	var pent math.NVal
	var hex math.NVal
	for {
		for tri.Val < pent.Val || tri.Val < hex.Val {
			tri = <-triangleNumsCh
		}
		if tri.Val == pent.Val && tri.Val == hex.Val {
			return fmt.Sprintf("%d", tri.Val)
		} else {
			log.Println(tri, pent, hex)
		}
		for pent.Val < tri.Val {
			pent = <-pentagonalNumsCh
		}
		for hex.Val < tri.Val {
			hex = <-hexagonalNumsCh
		}

	}
	panic("unreachable")
	return ""
}
