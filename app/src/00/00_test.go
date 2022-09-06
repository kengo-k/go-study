package study00

import (
  "testing"
  "strconv"
  "os"
  "io"
  "io/ioutil"
  "encoding/json"
)

func TestHello(t *testing.T) {  
  // varで変数宣言 変数名の後に型名が来る
  var hello string = "Hello"
  // ただし右辺で型が決まるなら型名は省略してよい
  var world = "World"
  // 宣言と同時に代入する場合はさらにvarも省略可能
  // ただし関数内部でしか使えない
  space := ` ` // 文字列リテラルはバッククオートも使用可能
  
  strcat := hello + space + world
  expected := "Hello World"
  if (strcat != expected) {
    t.Errorf("got `%v` , expected `%v`", strcat, expected)
  }
}

func TestNum(t *testing.T) {
  // 使える数値の型はint, uint, およびintN, uintN, flontN(Nはビット数。8,16,32,64が使える)
  // 型指定なしの場合はintもしくはfloat64になる
  num := 123
  // 型変換の際は型名を関数として使用する
  var num2 float64 = float64(num)
  // 数値を文字列に変換するにはstrconv.FormatFloatを使う
  // 第2引数は表現形式: 数値をそのまま文字列にしたい場合は'f'でよさそう
  // 第3引数は精度: 基本的には-1でよさそう
  // 第4引数はビット数: 基本的には64でよさそう
  numstr := strconv.FormatFloat(num2, 'f', -1, 64)
  expectedNumStr := "123"
  if (numstr != expectedNumStr) {
    t.Errorf("got `%v` , expected `%v`", numstr, expectedNumStr)
  }

  strnum := "123.456"
  // 数値文字列を数値に変換するにはstrconv.ParseFloatを使う
  // 戻り値として変換後の値とエラーを返す
  // 未使用の変数があるとビルドエラーになってしまうので必ずチェックする必要がある
  parsedFloat, err := strconv.ParseFloat(strnum, 64)
  // 成功時errはnil
  if (err != nil) {
    //num3str := strconv.FormatFloat(num3, 'f', -1, 64)
    t.Errorf("parse error, input=`%v`", strnum)
  }

  expectedParsedFloat := 123.456
  if (parsedFloat != expectedParsedFloat) {
    t.Errorf("got `%v` , expected `%v`", parsedFloat, expectedParsedFloat)
  }
}

func TestPointer(t *testing.T) {
  a := 100
  b := 200
  // 変数aのポインタをpaに格納
  pa := &a
  expectedPa := 100
  // *でポインタが指す値を取得する
  if (*pa != expectedPa) {
    t.Errorf("got `%v` , expected `%v`", *pa, expectedPa)
  }

  // ポインタ型にはnilを設定できる
  var pb *int = nil
  pb = &b

  sum := *pa + *pb
  expectedSum := 300
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)
  }
}

func TestArray(t *testing.T) {
  
  // 固定長配列
  nums := [3]int{10,20,30}
  expectedLen := 3
  if (len(nums) != expectedLen) {
    t.Errorf("got `%v` , expected `%v`", len(nums), expectedLen)
  }

    // インデックスアクセス
  expected1 := 20
  if (nums[1] != expected1) {
    t.Errorf("got `%v` , expected `%v`", nums[1], expected1)
  }

  // 可変長配列
  nums2 := []int{5,6,7}
  // appendで追加
  nums2 = append(nums2, 8)
  nums2 = append(nums2, 9)

  expectedLen = 5
  if (len(nums2) != expectedLen) {
    t.Errorf("got `%v` , expected `%v`", len(nums2), expectedLen)
  }

  // 切り出し
  nums3 := nums2[2:4]
  expectedLen = 2
  if (len(nums3) != expectedLen) {
    t.Errorf("got `%v` , expected `%v`", len(nums3), expectedLen)
  }

  expected0 := 7
  if (nums3[0] != expected0) {
    t.Errorf("got `%v` , expected `%v`", nums3[0], expected0)
  }
}

func TestMap(t *testing.T) {
  // mapを作成するには`map`か`make`を使用する
  status := map[int]string {
    200: "OK",
    404: "Not Found",
  }
  expected200 := "OK"
  if (status[200] != expected200) {
    t.Errorf("got `%v` , expected `%v`", status[200], expected200)
  }

  // 二番目の戻り値を取って存在チェック
  persons := make(map[string]int)
  persons["Yamada"] = 42
  yamada, exists := persons["Yamada"]
  expectedExists := true
  expectedYamada := 42
  if (yamada != expectedYamada) {
    t.Errorf("got `%v` , expected `%v`", yamada, expectedYamada)
  }
  if (exists != expectedExists) {
    t.Errorf("got `%v` , expected `%v`", exists, expectedExists)
  }

  // 存在しないキーを使用した場合
  _, exists = persons["Suzuki"]
  expectedExists = false
  if (exists != expectedExists) {
    t.Errorf("got `%v` , expected `%v`", exists, expectedExists)
  }
}

func TestIf(t *testing.T) {
  num := 1
  b := true
  
  // 他の言語にあるようないい感じにbool値として判断してくれる仕様はないためbool値を明示する必要がある
  if (num == 0) {
    t.Errorf("Error, num=%v", num)
    // ifの後のかっこはなくてもいい
  } else if b {
    // nop
  } else {
    t.Errorf("Error, num=%v", num)
  }

  status := map[int]string {
    200: "OK",
    404: "Not Found",
  }

  // 結果の取得と存在チェックをひとつのifで記述できる
  if _, exists := status[404]; !exists {
    t.Errorf("Error, 404 exists")
  }
}

func TestFor(t *testing.T) {
  names := []string{"Yamada", "Tanaka", "Suzuki"}
  
  // rangeは
  // 配列の場合index, valueを返す
  // mapの場合はkey, valueを返す
  for i, name := range names {
    if (i == 0) {
      expected0 := "Yamada"
      if (name != expected0) {
        t.Errorf("got `%v` , expected `%v`", name, expected0)
      }
    }
  }

  // 二番目がいらない場合は省略してもよい
  lastIndex := 0
  for i := range names {
    lastIndex = i
  }
  expectedLastIndex := 2
  if (lastIndex != expectedLastIndex) {
    t.Errorf("got `%v` , expected `%v`", lastIndex, expectedLastIndex)    
  }

  // 二番目だけがほしい場合はアンダースコアを使って使用しないことを明示する
  lastValue := ""
  for _, name := range names {
    lastValue = name
  }
  expectedLastValue := "Suzuki"
  if (lastValue != expectedLastValue) {
    t.Errorf("got `%v` , expected `%v`", lastValue, expectedLastValue)    
  }

  // 他の言語と同様にbreakやcontinueが使える
  nums := []int{1,2,3,4,5}
  breakedValue := 0
  for _, n := range nums {
    if n > 3 {
      breakedValue = n
      break
    }
  }
  expectedBreakValue := 4
  if (breakedValue != expectedBreakValue) {
    t.Errorf("got `%v` , expected `%v`", breakedValue, expectedBreakValue)    
  }

  // いわゆるwhile
  i := 0
  sum := 0
  for i < len(nums) {
    sum += nums[i]
    i++
  }
  expectedSum := 15
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)    
  }

  i = 0
  sum = 0
  // whileの無限ループ版
  for {
    if i >= len(nums) {
      break
    }
    sum += nums[i]
    i++
  }
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)    
  }

  // 伝統的なfor
  sum = 0
  for i := 0; i < len(nums); i++ {
    sum += nums[i]
  }
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)    
  }  
}

func TestSwitch(t *testing.T) {
  case200 := false
  case404 := false
  case500 := false
  caseDefault := false
  status := 404
  switch status {
  case 200:
    // 他の言語のswitchのようにデフォルトではfallthroughしないため
    // breakを書く必要がない
    case200 = true
  case 404:
    case404 = true
    // fallthroughしたい場合は明示する
    fallthrough
  case 500:
    case500 = true
    fallthrough
  default:
    caseDefault = true
  }
  if (case200) {
    t.Errorf("case200 must be `%v`", false)    
  }
  if (!case404) {
    t.Errorf("case404 must be `%v`", true)    
  }
  if (!case500) {
    t.Errorf("case500 must be `%v`", true)    
  }
  if (!caseDefault) {
    t.Errorf("caseDefault must be `%v`", true)    
  }      
}

func TestFunc(t *testing.T) {
  a := 5
  b := 12
  sum := add(a, b)
  expectedSum := 17
  if (sum != expectedSum) {
    t.Errorf("got `%v` , expected `%v`", sum, expectedSum)    
  }

  // 0除算エラー
  _, err := div(7, 0)
  if (!err) {
    t.Errorf("Error, 0 divide must be Error")
  }
  // int同士の除算の結果はint
  divd, _ := div(7, 2)
  expectedDivd := 3
  if (divd != expectedDivd) {
    t.Errorf("got `%v` , expected `%v`", divd, expectedDivd)    
  }

  // 高階関数
  add100 := addFactory(100)
  add100Result := add100(50)
  expectedAdd100Result := 150
  if (add100Result != expectedAdd100Result) {
    t.Errorf("got `%v` , expected `%v`", add100Result, expectedAdd100Result)    
  }
}

func TestDefer(t *testing.T) {
  // Goのエラー処理にJavaのような例外処理はないのでerr変数を返す方式を使う
  // 関数の戻り値の最後をerror型で返すようにするのが慣習。errがnilであれば正常終了と判断する
  f, err := os.Create("/tmp/sample.txt")
  if err != nil {
    t.Errorf("Error, cannot create a file")
    return
  }
  // deferを使うと後処理を定義できる
  // この関数を抜けたあとにファイルをクローズする
  defer f.Close()
  
  // すでにCloseを実行しているのでエラーになってしまうように見えてしまうが
  // Closeはこの関数が終了した後に行われるためファイル出力は正常に行われる
  io.WriteString(f, "Hello, World")
  // 出力したファイルを読み取る
  bytes, _ := ioutil.ReadFile("/tmp/sample.txt")
  content := string(bytes)
  exptectedContent := "Hello, World"
  if (content != exptectedContent) {
    t.Errorf("got `%v` , expected `%v`", content, exptectedContent)    
  }
}

func TestStruct(t *testing.T) {
  // デフォルト値で初期化された構造体を生成
  var yamada Student
  
  // 一部フィールドを初期化して生成
  tanaka := Student {
    Name: "tanaka",
  }
  
  // 構造体を初期化してポインタで取得
  pSuzuki := &Student {
    Name: "suzuki",
  }

  // 未設定の構造体フィールドは初期値に設定されていることを検証
  expectedYamadaName := ""
  if (yamada.Name != expectedYamadaName) {
    t.Errorf("got `%v` , expected `%v`", yamada.Name, expectedYamadaName)    
  }
  expectedYamadaAge := 0
  if (yamada.Age != expectedYamadaAge) {
    t.Errorf("got `%v` , expected `%v`", yamada.Age, expectedYamadaAge)    
  }

  // 設定した構造体フィールドの検証
  expectedTanakaName := "tanaka"
  if (tanaka.Name != expectedTanakaName) {
    t.Errorf("got `%v` , expected `%v`", tanaka.Name, expectedTanakaName)    
  }

  suzuki := *pSuzuki
  expectedSuzukiName := "suzuki"
  if (suzuki.Name != expectedSuzukiName) {
    t.Errorf("got `%v` , expected `%v`", suzuki.Name, expectedSuzukiName)    
  }
  // *pSuzuki.Nameと書くとエラーになってしまう模様
  if (suzuki.Name != (*pSuzuki).Name) {
    t.Errorf("Error") 
  }
}

func TestJson(t *testing.T) {
  f, err := os.Open("student.json")
  if err != nil {
    t.Errorf("Error, cannot open json file")
    return
  }
  decoder := json.NewDecoder(f)
  // jsonから読み込んだ内容を構造体にマッピング
  var student Student
  decoder.Decode(&student)
  expectedName := "yamada"
  if (student.Name != expectedName) {
    t.Errorf("got `%v` , expected `%v`", student.Name, expectedName)    
  }
}
