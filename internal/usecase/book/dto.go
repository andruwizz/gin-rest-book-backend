package book

type BookCreateParam struct {
	Title  string
	Author string
}

type BookListParam struct {
	Limit int
	Page  int
}

type BookGetParam struct {
	Id string
}

type BookUpdateParam struct {
	Id     string
	Title  string
	Author string
}

type BookDeleteParam struct {
	Id string
}
