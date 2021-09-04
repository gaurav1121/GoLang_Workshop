package main

// s "github.com/inancgumus/prettyslice"

func main() {
	var todo []string

	todo = append(todo, "sing")

	todo = append(todo, "run")

	todo = append(todo, "code", "play")
	tomorrow := []string{"see mom", "learn go"}
	todo = append(todo, tomorrow...)
	// s.Show("todo", todo)
}
