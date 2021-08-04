# weater-data
気象庁の[地上気象観測](https://www.jma.go.jp/jma/kishou/know/chijyou/surf.html), [アメダス](https://www.jma.go.jp/jma/kishou/know/amedas/kaisetsu.html)の観測所一覧の番号をファイルから読み込む

## ファイル名
地上気象観測: sfc_master_2021.index <br>
アメダス: ame_master_20210707.csv

## 実行方法
地上気象観測: `go run getNum.go sfc sfc_master_2021.index` <br>
アメダス: `go run getNum.go amd ame_master_20210707.csv`
