package util

import (
	"crypto/rand"
	"math"

	"math/big"
)

// 生成范围随机数随机数
func GenerateRangeNum(min, max int) int {
	// 随机种子
	randomNum := randomNum(int64(max - min))
	// 随机数
	return randomNum + min
}

// 生成随机浮点数
func GenerateFloatNum(max, min, validNum int) float32 {
	maxNum := min + randomNum(int64(max-min))
	//vailedNum := randomNum(math.Pow10(3))
	if maxNum == max {
		return float32(maxNum)
	} else {
		if validNum > 0 {
			valid := randomNum(int64(math.Pow10(validNum)))
			return float32(maxNum) + float32(valid)/float32(math.Pow10(validNum))
		} else {
			return float32(maxNum)
		}
	}
}

// 真正的随机数
func randomNum(maxInt int64) int {
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	return int(randomNum.Int64())
}
