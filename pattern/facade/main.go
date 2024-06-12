package main

import "fmt"

type Get struct{}

func (g *Get) getData() {
	fmt.Println("Get data")
}

type Post struct{}

func (p *Post) postData() {
	fmt.Println("Post data")
}

type RequestFacade struct {
	get  *Get
	post *Post
}

func newRequest() *RequestFacade {
	return &RequestFacade{
		get:  &Get{},
		post: &Post{},
	}
}

func (r *RequestFacade) actions() {
	fmt.Println("Accesible actions:")
	r.get.getData()
	r.post.postData()
}

func main() {
	facade := newRequest()
	facade.actions()
}
