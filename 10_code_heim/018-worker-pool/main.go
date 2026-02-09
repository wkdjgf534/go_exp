package main

import "fmt"

func main() {
	// Create new tasks

	tasks := []Task{
		&EmailTask{Email: "test1@example.com", Subject: "test1", MessageBoddy: "For test1"},
		&EmailTask{Email: "test2@example.com", Subject: "test1", MessageBoddy: "For test2"},
		&ImageProcessingTask{ImageURL: "/images/sample1.jpg"},
		&EmailTask{Email: "test3@example.com", Subject: "test1", MessageBoddy: "For test3"},
		&EmailTask{Email: "test4@example.com", Subject: "test1", MessageBoddy: "For test4"},
		&ImageProcessingTask{ImageURL: "/images/sample2.jpg"},
		&EmailTask{Email: "test1@example.com", Subject: "test2", MessageBoddy: "For test1"},
		&EmailTask{Email: "test2@example.com", Subject: "test2", MessageBoddy: "For test2"},
		&ImageProcessingTask{ImageURL: "/images/sample3.jpg"},
		&EmailTask{Email: "test3@example.com", Subject: "test2", MessageBoddy: "For test3"},
		&EmailTask{Email: "test4@example.com", Subject: "test2", MessageBoddy: "For test4"},
		&ImageProcessingTask{ImageURL: "/images/sample4.jpg"},
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:      tasks,
		concurency: 5, // Number of worker that can run a time
	}

	// Runy the pool
	wp.Run()
	fmt.Println("All tasks have been processed!")
}
