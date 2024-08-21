package models

type User struct {
	ID   int
	Name string
}

type Problem struct {
	ID          string
	Text        string
	Title       string
	Hint        string
	Constraints string
	Solution    string
	IsProject   bool
}
