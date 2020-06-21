package config

func PopulateAuthors() []Author {

	var authors []Author

	author1 := Author{ID: 1, Name: "Nick", Tutorials: []int{1, 2, 3, 4}}
	authors = append(authors, author1)

	author2 := Author{ID: 2, Name: "John", Tutorials: []int{5, 6}}
	authors = append(authors, author2)

	return authors
}

func PopulateTutorials() []Tutorial {

	var tutorials []Tutorial

	tutorial1 := Tutorial{ID: 1, Title: "Start with leetcode", Author: Author{ID: 1, Name: "Nick", Tutorials: []int{1, 2, 3, 4}}, Comments: []string{}, Likes: 10}
	tutorials = append(tutorials, tutorial1)

	tutorial2 := Tutorial{ID: 2, Title: "LeetCode Two Sum", Author: Author{ID: 1, Name: "Nick", Tutorials: []int{1, 2, 3, 4}}, Comments: []string{}, Likes: 20}
	tutorials = append(tutorials, tutorial2)

	tutorial3 := Tutorial{ID: 3, Title: "LeetCode Three Sum", Author: Author{ID: 1, Name: "Nick", Tutorials: []int{1, 2, 3, 4}}, Comments: []string{}, Likes: 30}
	tutorials = append(tutorials, tutorial3)

	tutorial4 := Tutorial{ID: 4, Title: "LeetCode longest palindrome", Author: Author{ID: 1, Name: "Nick", Tutorials: []int{1, 2, 3, 4}}, Comments: []string{}, Likes: 40}
	tutorials = append(tutorials, tutorial4)

	tutorial5 := Tutorial{ID: 5, Title: "Intro to System Design", Author: Author{ID: 2, Name: "John", Tutorials: []int{5, 6}}, Comments: []string{}, Likes: 1}
	tutorials = append(tutorials, tutorial5)

	tutorial6 := Tutorial{ID: 6, Title: "Horizontal vs Vertical Scaling", Author: Author{ID: 2, Name: "John", Tutorials: []int{5, 6}}, Comments: []string{}, Likes: 2}
	tutorials = append(tutorials, tutorial6)

	return tutorials

}
