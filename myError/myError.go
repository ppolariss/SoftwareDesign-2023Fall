package myError

type MyError struct{
	str string
}

func NewMyError(str string) MyError {
	return MyError{str:str}
}

// func myError.Error() string{
// 	return str
// }

func (e MyError) Error() string{
	return e.str
}

// func (e *MyError) Error() string{
// 	return e.str
// }