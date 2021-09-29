package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Random struct {
	r *rand.Rand
}

var (
	_instance *Random
	once sync.Once
)

func InstanceRandom() *Random {
	once.Do(func() {
		_instance = newRandom()
	})
	return _instance
}

func newRandom() *Random{
	return &Random{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
}


func (rd *Random) AssignNInt(n int) int {
	var result = fmt.Sprintf("%d", rd.r.Int() + 100000)
	n, _ = strconv.Atoi(result)
	return n
}
