# bookViewerBackend
## 本番環境
backend: yet  
frontend: yet

## 資料
Gorm: http://gorm.io/ja_JP/docs/  
echo:https://echo.labstack.com/guide  

## 他のREADME.md
API設計: https://github.com/KazuwoKiwame12/bookViewerBackend/tree/master/Controller  
DBディレクトリ: https://github.com/KazuwoKiwame12/bookViewerBackend/tree/master/DB  
Serviceディレクトリ: https://github.com/KazuwoKiwame12/bookViewerBackend/tree/master/Service  

## タスク管理
project:https://github.com/KazuwoKiwame12/bookViewerBackend/projects/1


## 導入
- ### 1: .envを.env.exampleより作成
```
$ cp .env.example .env
``` 
- ### 2: .envの中身を修正
```
※connect.goでは取得できていないので、直書きする必要性あり(絶対にcommitしていけない)
DATABASE_URL=自分のDATABASE_URL+"?sslmode=disable"
PORT=任意のポート番号
```
- ### 3: go.modにあるライブラリを自動インストール
```
$ go build
```

## 開発
### ファイル構成
- DB: Database関連の処理
  - Model: 各テーブルのCRUD機能
- Controller: API_URLと紐付ける機能
- Service: Controllerで用いる主要機能やDB処理を用いた機能などを提供
### Git・Githubについて
- Issue活用
  - Issueは様々な問題や疑問、課題を共有するための機能　
  - ラベル付けや、担当者へのassignなども行なう　
  - **終了**したissueは閉じること
- プロジェクト活用
  - 自身の進めるタスクなどのスケジュールや予定などを、まとめ共有するための機能
  - 一週間ごとにプロジェクトをまとめると進捗管理がしやすい
  - Issueとの関連付けが可能である→ [参考文献](https://help.github.com/ja/github/managing-your-work-on-github/adding-issues-and-pull-requests-to-a-project-board)
- Wiki活用
  - 自分が良いと思ったことや、開発の多末になるであろう知識を共有する機能
  - 参考記事があれば、リンクを貼っておきましょう
- **ブランチ**について
  - **masterに直接Pushは禁止**です（できないようにします）
  - **新たな作業(issue)を行う時は毎回masterをfetch, pullすること**
  - ブランチは作業ごとに切り替えること(最新のmasterの状態で作業を始めるため)
  - マージされたブランチ(リモート、ローカル)は削除すること
  - commitのメッセージは内容がわかるように
  - PullRequestには**Close #番号**とコメントに含めること
  - 作業が未完成の状態のPullRequestにはコメントに**WIP**と文字を入力すること

### 注意点
1. export(外部で使用する)する関数は関数名の最初の文字を大文字にしなければならない  
2. 自分でPullRequestをmergeしてはいけない
3. 使用しないファイルを作成していると、herokuのデプロイ時にエラーが発生する
