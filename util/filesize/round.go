package filesize

import (
	"fmt"
	"math"
	"strconv"

	"github.com/sirupsen/logrus"

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
	getSize, err := math2.Round(afterBase, .5, 2)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func":   "HumanFileSize",
			"reason": err,
			"input":  fmt.Sprintf("Value ( %f ) round point ( %f ) precision %d", afterBase, 0.5, 2),
		}).Infof("Use default round...")
		getSize = math.Round(afterBase)
	}
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
