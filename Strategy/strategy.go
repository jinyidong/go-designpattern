package Strategy

//策略模式，定义算法族
type Operator interface {
	Operate(i,j int) int
}

type AddOperator struct {
}

func (AddOperator) Operate(i, j int) int {
	return i+j
}

type SubOperator struct {
}

func (SubOperator) Operate(i, j int) int {
	return i-j
}

type Calculator struct {
	Strategy Operator
}

func NewCalculater(operator Operator) Calculator {
	calculator:=&Calculator{}

	calculator.Strategy=operator

	return *calculator
}





