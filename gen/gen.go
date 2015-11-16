package gen

type Gen interface {
	//Generate i examples
	Generate(i int)

	GenerateOne()

	String() string
}