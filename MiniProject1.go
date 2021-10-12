package main

import (
	"fmt"
	"time"
)

func main() {
	f1 := newFork(1)
	f2 := newFork(2)
	f3 := newFork(3)
	f4 := newFork(4)
	f5 := newFork(5)

	p1 := newPhilosopher(1, f1, f2)
	p2 := newPhilosopher(2, f2, f3)
	p3 := newPhilosopher(3, f3, f4)
	p4 := newPhilosopher(4, f4, f5)
	p5 := newPhilosopher(5, f1, f5)

	go p5.think()
	go p1.think()
	go p2.think()
	go p3.think()

	go p4.think()
	go f1.exist()
	go f2.exist()
	go f3.exist()
	go f4.exist()
	go f5.exist()

	for {

		fmt.Println("----------------------------------------------------")
		time.Sleep(2 * time.Second)

		askForTimesEaten(p1)
		askForTimesEaten(p2)
		askForTimesEaten(p3)
		askForTimesEaten(p4)
		askForTimesEaten(p5)
		fmt.Println("")
		askForStatus(p1)
		askForStatus(p2)
		askForStatus(p3)
		askForStatus(p4)
		askForStatus(p5)
		fmt.Println("")
		askForTimesUsed(f1)
		askForTimesUsed(f2)
		askForTimesUsed(f3)
		askForTimesUsed(f4)
		askForTimesUsed(f5)
		fmt.Println("")
		askForIsTaken(f1)
		askForIsTaken(f2)
		askForIsTaken(f3)
		askForIsTaken(f4)
		askForIsTaken(f5)

	}
}

func askForTimesEaten(p *philosopher) {
	p.inputChannel <- 0

	var timesEaten = <-p.outputChannel
	fmt.Printf("Philosopher %d has eaten %d times", p.id, timesEaten)
	fmt.Println("")
}

func askForStatus(p *philosopher) {
	p.inputChannel <- 1

	var status = <-p.outputChannel
	if(status == 1){
		fmt.Printf("Philosopher %d is eating", p.id)
	} else if(status == 0){
		fmt.Printf("Philosopher %d is thinking", p.id)
	}
	
	fmt.Println("")
}

func askForIsTaken(f *fork) {
f.inputChannel <- 1
var status = <-f.outputChannel
	if(status == 1){
		fmt.Printf("Fork %d is taken", f.id)
	} else if(status == 0){
		fmt.Printf("Fork %d is free", f.id)
	}
	fmt.Println("")
}

func askForTimesUsed(f *fork) {
	f.inputChannel <- 0

	var timesUsed = <-f.outputChannel
	fmt.Printf("Fork %d has been used %d times", f.id, timesUsed)
	fmt.Println("")
}
