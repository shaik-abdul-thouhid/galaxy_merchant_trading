package setup

type TestInput[T comparable] struct {
	Name           string
	Input          string
	ExpectedResult T
	ExpectedError  string
}
