package filesize

import (
	"fmt"
	"math"
	"strconv"

	math2 "github.com/tikivn/ops-delivery-kit/util/math"
)

var (
	suffixes [5]string
)

func HumanFileSize(size float64) string {
	fmt.Println(size)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	afterBase := math.Pow(1024, base-math.Floor(base))
	getSize := math2.RoundWithPrecision(afterBase, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
