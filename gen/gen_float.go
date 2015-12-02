package gen

import (
  "math/rand"
  "fmt"
  "time"
)

type GenFloat struct {
}

func (gen *GenFloat) Generate(n int) ([]float32, error) {
	if n == 0 {
		return []float32{}, fmt.Errorf("Number of examples, must be > 0")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	result := []float32{}
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
func (gem *GenFloat) GenerateOne() (float32, error) {
	return rand.Float32(), nil
}

func (gem *GenFloat) DefaultRiles() []float32 {
	return []float32{0.0, -1.0}
}