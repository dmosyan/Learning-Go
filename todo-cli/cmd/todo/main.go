package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmosyan/Learning-Go/todo-cli"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark as completed")
	del := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("sample todo")
		executeStore(todos, todoFile)
	case *complete > 0:
		err := todos.Complete(*complete)
		handleError(err)
		executeStore(todos, todoFile)
	case *del > 0:
		err := todos.Delete(*del)
		handleError(err)
		executeStore(todos, todoFile)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func executeStore(todos *todo.Todos, todoFile string) {
	err := todos.Store(todoFile)
	handleError(err)
}
