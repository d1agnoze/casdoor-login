package err

type InvalidInputError struct{}

func (*InvalidInputError) Error() string {
	return "input string is empty"
}
