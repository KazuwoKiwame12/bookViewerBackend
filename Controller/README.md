# APIの設計
## 本の内容を取得
- **Method**: GET
- **URL**: /api/mine/book/{book_id}
- **受け取るデータ**
```
id: 本のid、defaultで1
```
- **返すデータ**
```
{
  "book": [
    {
      "chapterId": int,
      "pages": [
        {
          "pageId": int,
          "sentences": [
            {
              "sentenceId": int,
              "content": text,
              "hasQuestion": boolean(あり:0, なし:1)
            },
            .......
          ]
        },
        .........
      ]
    },
    ..........
  ]
}
```

## 質問する
- **Method**: POST
- **URL**: /api/question/create
- **受け取るデータ**
```
{
  userId : int
  sentenceId : int
  pageNum : int
  rowNum : int
  title : string
  content : string
}
```
- **返すデータ**
```
{
  hasSuccess : bool
}

```

## 章の質問一覧を取得
- **Method**: GET
- **URL**: /api/chapter/{id}
- **受け取るデータ**
```
{
  id : 章のid
}
```
- **返すデータ**
```
questionId
title(質問のタイトル)
name(ユーザー)
createdAt(投稿日時)

{
  questions: [{
    questionId: int,
    userName: string
    title:string,
    createdAt: datetime
  }]
}

```


## 質問の内容を取得
- **Method**: GET
- **URL**: /api/question/{id}/content
- **受け取るデータ**
```
id: 質問id
```
- **返すデータ**
```
{
  "userName": string,
  "title": string,
  "content": string,
  "pageNum": int,
  "rowNum": int,
  "createdAt": timestamp,
}
```

## 著者の返信を取得
- **Method**: GET
- **URL**: /api/question/{id}/author/answer
- **受け取るデータ**
```
id: 質問id,
```
- **返すデータ**
```
{
  "answers": [
    {
      "createdAt": timestamp,
      "content": text,
      "likeNum": int(default: 40)
    },
    ........
  ]
}
```

## 読者の返信を取得
- **Method**: GET
- **URL**: /api/question/{id}/reader/answer
- **受け取るデータ**
```
id: 質問id,
```
- **返すデータ**
```
{
  "answers": [
    {
      "userName": string
      "createdAt": timestamp,
      "content": text,
      "likeNum": int(default: random)←高い順でsortする
    },
    ........
  ]
}
```

## 質問のidに対応するページ内容を取得
- **Method**: GET
- **URL**: /api/question/{id}/page
- **受け取るデータ**
```
id: 質問id,
```
- **返すデータ**
```
{
  "contents": [{
    "content": text
  }
  ...
  ]
}
```

## 質問に返信する
- **Method**: POST
- **URL**: /api/question/reply
- **受け取るデータ**
```
{
  userId : int
  questionId : int
  content : string
}
```
- **返すデータ**
```
{
  hasSuccess : bool
}
```

## 質問の検索
- **Method**: GET
- **URL**: /api/question/search/{title}
- **受け取るデータ**
```
title: 質問のタイトル,
```
- **返すデータ**
```
{
  questions: [
    {
      questionId: int,
      userNa,e: string
      title:string,
      createdAt: datetime
    },
    ..............
  ]
}
```

## 文書に対応した質問の検索
- **Method**: GET
- **URL**: /api/question/search/sentence/{id}
- **受け取るデータ**
```
id: 文書id,
```
- **返すデータ**
```
{
  question: [{
    questionId: int,
    userName: string
    title:string,
    createdAt: datetime
  }]
}
```
