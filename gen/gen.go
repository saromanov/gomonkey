package gen

type Gen interface {
	//Generate i examples
	Generate(i int)

	GenerateOne()

	String() string

	// DefaultRiles provides some default items
	// for each type of generation
	DefaultRules()
}
