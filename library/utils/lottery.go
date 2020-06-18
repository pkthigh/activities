package utils

import (
	"activities/common/errs"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Lottery lottery
func Lottery(f float64) (bool, error) {
	if f < 0 {
		return false, errs.IllegalParameter.Error()
	}

	value, err := strconv.ParseFloat(fmt.Sprintf("%.4f", f), 64)
	if err != nil {
		return false, err
	}

	var pool []bool
	for i := 0; i < int(value*10000); i++ {
		pool = append(pool, true)
	}
	for i := 0; i < int(10000-value*10000); i++ {
		pool = append(pool, false)
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(pool)
	for i := 0; i < length; i++ {
		number := random.Intn(length)
		pool[length-1], pool[number] = pool[number], pool[length-1]
	}

	return pool[random.Intn(length)-1], nil
}
