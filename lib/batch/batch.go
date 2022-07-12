package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	// buffered channels
	toDo := make(chan int, int(n))
	received := make(chan user, int(n))

	for i := 0; i < int(pool); i++ {
		go worker(toDo, received)
	}
	// don't understand why this part goes after go worker :(
	for i := 0; i < int(n); i++ {
		toDo <- i
	}
	close(toDo)
	for j := 0; j < int(n); j++ {
		res = append(res, <-received)
	}
	close(received)
	return res
}

// worker function with two channels as arguments.
func worker(toDo <-chan int, received chan<- user) {
	for pin := range toDo {
		received <- getOne(int64(pin))
	}
}
