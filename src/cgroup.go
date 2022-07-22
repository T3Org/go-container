package main

import (
	"fmt"
	"os"
)

/** 
cgroupは"Control Group"の略。
プロセスIDをグループに追加すると、共通の設定を適用できる。
ホストOSが持つCPUやメモリなどのリソースの制限をかけることができる。
**/

// プロセスIDを取得する(ランダムな文字列でも良い)


// メモリの制限を設定する
func main() {
	currenProcessInfo()
}

func currenProcessInfo() {
	// プロセスIDを取得する
	pid := os.Getpid()
	
	fmt.Printf("%v\n", pid)

}