package main

import "testing"

func TestGetAllArticles( t *testing.T ){
	alist := getAllArticles()

	if ( len( alist ) != len( articleList ) ){
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content || v.Title != articleList[i].Title || v.ID != articleList[i].ID {
			t.Fail()
			break
		}
	}
}