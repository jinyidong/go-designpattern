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
