package main

import (
	"sync"
)

var arbiter sync.Mutex

type philosopher struct {
	id                 int
	smallFork          *fork
	bigFork            *fork
	hasBigFork         bool
	hasSmallFork       bool
	isEating           int
	numberOfTimesEaten int
	inputChannel       chan int
	outputChannel      chan int
}

func newPhilosopher(id int, smallFork *fork, bigFork *fork) *philosopher {
	p := philosopher{id: id, smallFork: smallFork, bigFork: bigFork}
	p.hasBigFork = false
	p.hasSmallFork = false
	p.isEating = 0
	p.numberOfTimesEaten = 0
	p.inputChannel = make(chan int)
	p.outputChannel = make(chan int)
	return &p
}

func (self *philosopher) readInput() {
	var input = <-self.inputChannel
	if input == 0 {
		self.outputChannel <- self.numberOfTimesEaten
	} else if input == 1 {
		self.outputChannel <- self.isEating
	}
}

func (self *philosopher) think() {
	for {
		self.pickUpFork()
		self.readInput()
	}
}

func (self *philosopher) pickUpFork() {
	if !self.hasSmallFork {
		self.pickUpSmallFork()
	} else if !self.hasBigFork {
		self.pickUpBigFork()
	}
}

func (self *philosopher) eat() {
	self.isEating = 1
	self.numberOfTimesEaten++
	self.putDownForks()
	self.isEating = 0
}

func (self *philosopher) pickUpSmallFork() {
	arbiter.Lock()
	if self.smallFork.isTaken == 0 {
		self.smallFork.pickedUp()
		self.hasSmallFork = true
	}
	arbiter.Unlock()
}

func (self *philosopher) pickUpBigFork() {
	arbiter.Lock()
	if self.bigFork.isTaken == 0 {
		self.bigFork.pickedUp()
		self.hasBigFork = true
	}
	arbiter.Unlock()

	if self.hasBigFork {
		self.eat()
	}
}

func (self *philosopher) putDownForks() {
	self.hasBigFork = false
	self.hasSmallFork = false
	self.bigFork.putDown()
	self.smallFork.putDown()
}
