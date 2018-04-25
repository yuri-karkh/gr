package main

import "errors"

type article struct {
	ID 		int `json:"id"`
	Title 	string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ ID : 1, Title : "Article 1", Content : "Content 1" },
	article{ ID : 2, Title : "Article 2", Content : "Content 2" },
}

func getArticleById( id int ) ( *article, error ) {
	for _, a := range articleList {
	  if a.ID == id {
				return &a, nil
	  }
	}
	return nil, errors.New( "Article not found" )
}

func getAllArticles() []article {
	return articleList
}