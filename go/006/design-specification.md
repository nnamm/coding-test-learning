# コーディング用の設計メモ

## ①地形データ作成

・行と列数の入力待ち（HとWの入力）
・行数に応じて入力待ち（Hの数分以下をループ）
　- [i][W]runeに入力文字列を設定
　- これが地形データとなる
・以下のデータを初期化する
　- キューのスライスを宣言（Point構造体を持つスライス）※1
　- 訪問済みスライスを宣言（booleanを持つスライス）※2
　- 上下左右、4方向の移動を定義 ※3
　- スタート地点であるstartRow, startColを求める（H行W列のスライスから"S"の文字がある位置）

```
// ※1
type Point struct {
    row, col, distance int
}
queue := []Point{} // 探索待ちの地点

// ※2
visited := make([][]bool, H) // 訪問済みフラグ

// ※3
directions := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}} // 上、下、左、右
```

## ②BFS（幅優先探索）ループ

・スタート地点をキューに登録。（キューはFIFOで管理）
・キューにデータがある間以下の処理を行う。
　- キューの位置をcurrentPointとして保持。取り出したキューは削除。
　- directions分、以下のループを行う。
　　- currentPointがゴールか判定する。ゴールならループを抜ける。（isGoal関数）
　　- currentPointに対して上下左右に移動できるかチェックする。（canMove関数）
　　- 訪問済みかチェックする。（visited）
　　- 移動可能で訪問済みでない場合、
　　　canMoveの返却値を確認（Gならループを抜ける）
　　　currentPointをキューに追加する。
　　　currentPoint.distance + 1を新しいPointのdistanceとして設定する。
　　　訪問済みの設定を行う。
・isGoal関数は以下の通り。
　- 引数：
　　currentPoint
　　地形データ（[H][W]rune）
　- 返却値：
　　bool（true＝Gである、false＝Gではない）
　- 引数からcurrentPointの位置にある文字列を調べる。
　　Gの場合、trueを返す
　　G以外の場合、falseを返す
・canMove関数は以下の通り。
　- 引数：
　　currentPointに対する上下左右のポイント（newRow, newCol）
　　地形データ（[H][W]rune）
　- 返却値：
　　bool（true＝移動可能、false＝移動不可）
　- 引数からnewRowとnewColの位置にある文字列を調べる。
　　0,Gの場合、trueを返す
　　1の場合、falseを返す

## 終了条件

・isGoalでゴール（”G"）に到達したら探索終了。
　- 当該Pointのdistanceが最短探索経路となるため、その値を出力する。
・キューにPointがなくなったら探索終了。
　- ゴールに到達できないと判断し、`-``を出力する。
