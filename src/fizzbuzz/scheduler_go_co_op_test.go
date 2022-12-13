package fizzbuzz

import (
	"fmt"
	"testing"
	_ "testing"
	"time"

	"github.com/go-co-op/gocron"
)

func TestCron(t *testing.T) {
	RunScheduler1()
}
func RunScheduler1() {
	s := gocron.NewScheduler(time.UTC)

	// s.Every(2).Day().Tag("tag").At("10:00").Do(task2)
	s.Every(4).Second().Tag("tag").Do(task2)
	s.RunByTag("tag")
}

func task2() {
	fmt.Println("Task is being performed.")
}
