package structs

type Car struct {
	name    string
	Speed   float64 // 屬性同樣也要符合首字大寫 才能給外部檔案呼叫使用!
	Address Location
	Profile
}
