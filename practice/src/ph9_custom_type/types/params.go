package types

type MyTypeB string
type MyTypeC string

// Add 輸入輸出目前都需要注重 型別 (TODO NOTES 可以研究是否可以向上兼容)
func (m MyTypeB) Add(s string) string {
	return string(m) + s
}
