# Enablement Bootcamp for Gopherizing (2023.9)

## 本試験は、Enablement Bootcamp for Gopherizingのエントリー課題です

- ブートキャンプの詳細については、connpassページをご確認ください。

conpass：https://knowledgework.connpass.com/event/290443

## お題　ー　splitコマンドのクローンGoで実装

- splitコマンド　：　csvなどのファイル行数で分割するコマンド

- 実装するオプション
  
  ***必須***
  
  - "-n" :  分割するファイル数の指定
  - "-l" :  分割する１ファイルの行数のして
  - "-b" :  分割する１ファイルのバイト数の指定
  - "-a" :  分割ファイルのsuffixの文字数の指定
  
  ***追加オプション*** 
  - "-col"  : csvにおける列による分割

- 実装するエラー対処
- 実行速度の改善
- 大量データを扱えるようにする
- Unit Testを書く

## 細則事項

- 使用できるパッケージは標準と自作パッケージのみ
- ファイル、パッケージの分け方は自由
- ファイルを分けた場合、解答欄にはtxtar形式で記述する
    - txtar形式でソースコードを出力するには以下のコマンドを使用する
    
    ~~~bash

    find . type -f | xargs go run golang.org/x/exp/cmd/txtar@latest

    ~~~

## 実装

## 作成したコマンドについて工夫した点

- ***観点***
  - 対応したオプション
  - 対応したイレギュラーな入力
  - パフォーマンスに関する工夫
  - テストコードのポイント
