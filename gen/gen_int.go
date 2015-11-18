package gen

import (
"fmt"
"math/rand"
)

type GenInt struct {

}

func (gen* GenInt) Generate(i int) ([]int, error) {
	if i == 0 {
		return []int{}, fmt.Errorf("Number of examples, must be > 0")
	}

	return []int{}, nil

}

//GenerateOne provodes generation of one examples
func (gem *GenInt) GenerateOne() (int, error) {
	return rand.Intn(100000000), nil
}

func (gem *GenInt) DefaultRiles()[]int {
	return []int{0,-1}
}