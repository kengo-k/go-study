package study00

// 一般的な関数定義
func add(a int, b int) int {
	return a + b
}

// a, bが同じ型なので型をまとめて定義している
func div(a, b int) (ret int, err bool) {

	// 戻り値の変数名を定義することができる
	// 既存の言語のようにreturnで値を返すのではなく変数に値を設定して単にreturnとだけすればよい

	if b == 0 {
		err = true
		return
	}
	// int同士で割り算した場合は結果はintとなる
	ret = a / b
	err = false
	return
}

// 関数を返す関数と無名関数
func addFactory(a int) func(int) int {
	f := func(b int) int {
		return a + b
	}
	return f
}

// 構造体
type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
