package math

type NVal struct {
	Val uint64
	N   int
}

var triangleNumsCh chan NVal
var highestTriangleNum NVal
var triangleNums map[uint64]NVal

func TriangleN(n int) NVal {
	return NVal{Val: uint64(n) * uint64(n+1) / uint64(2), N: n}
}

func GenNVals(store map[uint64]NVal, ch chan NVal, nextFn func(n int) NVal, highest *NVal) {
	if highest == nil {
		highest = &NVal{}
	}
	for {
		next := nextFn(highest.N + 1)
		if store != nil {
			store[next.Val] = next
		}
		*highest = next
		ch <- next
	}
}

var pentagonalNumsCh chan NVal
var highestPentagonalNum NVal
var pentagonalNums map[uint64]NVal

func PentagonalN(n int) NVal {
	return NVal{Val: uint64(n) * uint64(3*n-1) / uint64(2), N: n}
}

var hexagonalNumsCh chan NVal
var highestHexagonalNum NVal
var hexagonalNums map[uint64]NVal

func HexagonalN(n int) NVal {
	return NVal{Val: uint64(n) * uint64(2*n-1), N: n}
}

func init() {
	triangleNums = make(map[uint64]NVal)
	triangleNumsCh = make(chan NVal)
	go GenNVals(triangleNums, triangleNumsCh, TriangleN, &highestTriangleNum)
	pentagonalNums = make(map[uint64]NVal)
	pentagonalNumsCh = make(chan NVal)
	go GenNVals(pentagonalNums, pentagonalNumsCh, PentagonalN, &highestPentagonalNum)
	hexagonalNums = make(map[uint64]NVal)
	hexagonalNumsCh = make(chan NVal)
	go GenNVals(hexagonalNums, hexagonalNumsCh, HexagonalN, &highestHexagonalNum)
}

func IsTriangleNum(val uint64) (NVal, bool) {
	for val > highestTriangleNum.Val {
		<-triangleNumsCh
	}
	v, ok := triangleNums[val]
	return v, ok
}
func IsPentagonalNum(val uint64) (NVal, bool) {
	for val > highestPentagonalNum.Val {
		<-pentagonalNumsCh
	}
	v, ok := pentagonalNums[val]
	return v, ok
}
func IsHexagonalNum(val uint64) (NVal, bool) {
	for val > highestHexagonalNum.Val {
		<-hexagonalNumsCh
	}
	v, ok := hexagonalNums[val]
	return v, ok
}
