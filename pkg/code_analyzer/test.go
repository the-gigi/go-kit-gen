package code_analyzer

//func Increment(x int) int {
//	return x + 1
//}

type Foo interface {
	Op1(a int, b int) (int, error)
	Op2(x map[string]interface{}, y *[]bool) (res string, err error)
	//Op3()
}
