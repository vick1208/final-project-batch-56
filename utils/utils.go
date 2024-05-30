package utils

var Empty struct{}

func SumFee(mainFee, addFee uint) uint {
	return mainFee + addFee
}
