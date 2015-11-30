package gen

import (
	"fmt"
	"math/rand"
	"time"
)

type GenInt struct {
}

func (gen *GenInt) Generate(n int) ([]int, error) {
	if n == 0 {
		return []int{}, fmt.Errorf("Number of examples, must be > 0")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	result := []int{}
	for i := 0; i < n; i++ {
		num, err := gen.GenerateOne()
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil

}

//GenerateOne provodes generation of one examples
func (gem *GenInt) GenerateOne() (int, error) {
	return rand.Intn(100000000), nil
}

func (gem *GenInt) DefaultRiles() []int {
	return []int{0, -1}
}
