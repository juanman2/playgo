package playgo

import (
	"fmt"
	"testing"
	"time"
)

func TestBBQSimple(t *testing.T) {
	tests := []struct {
		producer int
		consumer int
		ops      []string
		args     []int
		out      [][]int
	}{
		{
			producer: 1,
			consumer: 1,
			ops:      []string{"BoundedBlockingQueue", "enqueue", "dequeue", "dequeue", "enqueue", "enqueue", "enqueue", "enqueue", "dequeue"},
			args:     []int{2, 1, -1, -1, 0, 2, 3, 4, -1},
			out: [][]int{
				{1, 0, 2, 2},
			},
		},
		{
			producer: 3,
			consumer: 4,
			ops:      []string{"BoundedBlockingQueue", "enqueue", "enqueue", "enqueue", "dequeue", "dequeue", "dequeue"},
			args:     []int{3, 1, 0, 2, -1, -1, -1},
			out: [][]int{
				{1, 0, 2, 0},
				{1, 2, 0, 0},
				{0, 1, 2, 0},
				{0, 2, 1, 0},
				{2, 0, 1, 0},
				{2, 1, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		fmt.Printf("Test %v\n", tt)
		if tt.ops[0] != "BoundedBlockingQueue" {
			t.Errorf("Test %v does not start with Queueue allocation", tt)
		}
		// make the queue
		bbq := NewBoundedBlockingQueue(tt.args[0])

		// make a channel to get the output back from the threads
		ci := make(chan int, tt.consumer)
		cc := make(chan int, tt.consumer)
		cp := make(chan int, tt.producer)
		ct := make(chan int, tt.producer)

		// start the producer threads
		for i := 0; i < tt.producer; i = i + 1 {
			go func() {
				// receive from channel cp until closed
				for element := range cp {
					bbq.Enqueue(element)
				}
				ct <- 1
			}()
		}

		for i := 0; i < tt.consumer; i = i + 1 {
			go func() {
				// receive from channel cp until closed
				for range cc {
					element := bbq.Dequeue()
					ci <- element // signal & send back
				}
				ct <- 1
			}()
		}

		for i, op := range tt.ops {
			switch op {
			case "enqueue":
				cp <- tt.args[i]
			case "dequeue":
				cc <- 1
			}
		}

		// receive the dequeued integers from the ci channel, or timeout
		got := make([]int, len(tt.out[0]))

		for i := 0; i < len(got)-1; i = i + 1 {
			select {
			case e := <-ci:
				got[i] = e
			case <-time.After(5 * time.Second):
				fmt.Printf("Timing out: len=%d got=%d\n",
					len(got)-1, i)
				panic("timed out")
			}
		}

		// close the channels
		close(cp)
		close(cc)

		// Make sure all go routines exited
		for i := 0; i < tt.consumer+tt.producer; i = i + 1 {
			select {
			case <-ct: // do nothing
			case <-time.After(5 * time.Second):
				fmt.Printf(
					"Timing out waiting for go routines: got=%d\n",
					i)
				panic("timed out")

			}
		}

		// now get the size
		got[len(got)-1] = bbq.Size()
		fmt.Printf("Size of %v is %d got:%v\n", bbq, bbq.Size(), got)

		// Verify what we got.
		// The function below matches against all possible outcomes.
		if !verifyMatch(tt.out, got) {
			t.Errorf(
				"BBQ failed to find a match, expected %v \n got %v \n",
				tt.out, got)
		}
	}
}

func verifyMatch(out [][]int, got []int) bool {
	match := false
	for i := 0; i < len(out) && !match; i = i + 1 {
		count := 0
		for j := 0; j < len(out[i]); j = j + 1 {
			if got[j] != out[i][j] {
				break
			}
			count = count + 1
		}
		if count == len(got) {
			match = true
			break
		}
	}
	return match
}
