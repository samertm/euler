package math

type triangleNum struct {
	val int
	n   int
}

var triangleNumsCh chan triangleNum
var highestTriangleNum triangleNum
var triangleNums map[int]bool

func triangleN(n int) triangleNum {
	return triangleNum{val: n * (n + 1) / 2, n: n}
}

func init() {
	triangleNums = make(map[int]bool)
	triangleNumsCh = make(chan triangleNum)
	go func() {
		for {
			next := triangleN(highestTriangleNum.n + 1)
			triangleNums[next.val] = true
			highestTriangleNum = next
			triangleNumsCh <- next
		}
	}()
}

func IsTriangleNum(val int) bool {
	for val > highestTriangleNum.val {
		<-triangleNumsCh
	}
	return triangleNums[val]
}
