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
		if rand.Intn(10) % 2 == 0 {
			result = append(result, num)
		} else {
			result = append(result, negative(num))
		}
	}
	return result, nil

}

//GenerateOne provodes generation of one examples
func (gem *GenInt) GenerateOne() (int, error) {
	value := rand.Intn(3)
	if value == 0 {
		return gem.generateOneSmall(), nil
	}
	if value == 1 {
		return gem.generateOneAvg(), nil
	}

	return gem.generateOneBig(), nil
}

func (gem *GenInt) generateOneBig() int {
	return rand.Intn(100000000)
}

func (gem *GenInt) generateOneAvg() int {
	return rand.Intn(1000000)
}

func (gem *GenInt) generateOneSmall() int {
	return rand.Intn(1000)
}

func (gem *GenInt) DefaultRiles() []int {
	return []int{0, -1}
}

func negative(num int) int {
	return -num
}
