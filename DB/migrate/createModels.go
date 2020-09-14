package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name	 string			`json:id`
	Email  string			`json:email` 
	Password string   `json:password`
	Status int        `json:status`
}

type Reply struct {
	Content string
  CreatedAt timestamp
}

type Genre struct {
	Name string
}

type Book struct {
	Title string
	Image string
	Price int
	Author string
	Bio string
}

type Chapter struct {
	Number int
}	

type Page struct {
	Number int	
}

type Like struct {
}

type Sentence struct {
	Content string
}

type Question struct {
	Title string
	Content string
	PageNum int
	RowNum int
	CreatedAt time.Time
}

type UserBook struct {
	Status int
	CreatedAt time.Time
}






