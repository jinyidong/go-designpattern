package Strategy

import (
	"testing"
)

func TestAddOperator_Operate(t *testing.T) {
	calculator:=NewCalculater(AddOperator{})

	expectResult:=calculator.Strategy.Operate(1,2)

	if expectResult!=3 {
		t.Error("加法计算错误")
	}
}

func TestSubOperator_Operate(t *testing.T) {
	calculator:=NewCalculater(SubOperator{})

	expectResult:=calculator.Strategy.Operate(1,2)

	if expectResult!=-1 {
		t.Error("减法计算错误")
	}
}

type DivisionOperator struct {
}

//整数除法，demo
func (DivisionOperator) Operate(i, j int) int {
	return i/j
}

func TestNewOperator(t *testing.T)  {
	calculator:=NewCalculater(DivisionOperator{})

	expectResult:=calculator.Strategy.Operate(2,1)

	if expectResult!=2 {
		t.Error("除法计算错误")
	}
}