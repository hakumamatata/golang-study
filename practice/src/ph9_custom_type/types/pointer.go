package types

import "fmt"

type MyTypeA string

// MethodA 通常定義方法時 會統一使用指標傳入 或者是 直接傳值 不會混用
func (m MyTypeA) MethodA() {
	fmt.Println(m)
	m = "Updated By A"
}

func (m *MyTypeA) MethodB() {
	fmt.Println(*m)
	*m = "Updated By B"
}
