package config

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Tutorials []int  `json:"tutorials"`
}

type Tutorial struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Author   Author   `json:"author"`
	Comments []string `json:"comments"`
	Likes    int      `json:"likes"`
}
