package main

type User struct {
	display_name string
	login_name   string
	company      string
	location     string
	followers    int
	email        string
}

type Repo struct {
	visibility    string
	about         string
	tags          []string
	creation      string
	license       string
	stars         int
	watching      int
	forks         int
	open_issues   int
	pull_requests int
	language      string
}
