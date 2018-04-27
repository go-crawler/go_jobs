package page

import (
	"math"
)

func CalculateTotalPage(totalCount, pageSize float64) int {
	totalPage := float64(totalCount) / float64(pageSize)

	return int(math.Ceil(totalPage))
}
