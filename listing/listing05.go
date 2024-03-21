package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

/*
когда объявляется переменную err с помощью var err error,
она инициализируется значением по умолчанию для типа error,
которое является nil. Поэтому при присвоении err = test(),
переменная err будет содержать значение типа *customError,
которое не равно nil, что и отображается в выводе.
*/
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
