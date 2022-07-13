package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber() int {

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 99

	res := rand.Intn(max-min+1) + min
	//fmt.Println(res)

	return res
}

func GenerateAccountNumber() string {

	res_s := ""
	for i := 0; i < 8; i++ {
		val := GenerateRandomNumber()
		res_s += strconv.Itoa(val)
	}

	return res_s
}
