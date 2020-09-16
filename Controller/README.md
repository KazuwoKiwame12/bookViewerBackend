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
  "Book": [
    {
      "ChapterID": int,
      "Pages": [
        {
          "PageID": int,
          "Sentences": [
            {
              "SentenceID": int,
              "Content": text,
              "HasQuestion": boolean(あり:0, なし:1)
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
  user_id : int
  sentence_id : int
  page_num : int
  row_num : int
  title : string
  content : string
}
```
- **返すデータ**
```
{
  has_success : bool
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
question_id
title(質問のタイトル)
name(ユーザー)
created_at(投稿日時)

{
  questions: [{
    question_id: int,
    user_name: string
    title:string,
    created_at: datetime
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
  "user_name": string,
  "title": string,
  "content": string,
  "page_num": int,
  "row_num": int,
  "created_at": timestamp,
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
      "created_at": timestamp,
      "content": text,
      "like_num": int(default: 40)
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
      "user_name": string
      "created_at": timestamp,
      "content": text,
      "like_num": int(default: random)←高い順でsortする
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
  "content_list": [{
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
  user_id : int
  question_id : int
  content : string
}
```
- **返すデータ**
```
{
  has_success : bool
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
  Questions: [
    {
      question_id: int,
      user_name: string
      title:string,
      created_at: datetime
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
  Questions: [{
    question_id: int,
    user_name: string
    title:string,
    created_at: datetime
  }]
}
```
