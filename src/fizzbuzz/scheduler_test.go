package fizzbuzz

import (
	"fmt"
	"testing"
	_ "testing"

	"github.com/jasonlvhit/gocron"
)

func TestScheduler(t *testing.T) {
	runScheduler1()
}

func runScheduler1() {
	s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(task)
	<-s.Start()
}

func task() {
	fmt.Println("Task is being performed.")
}
