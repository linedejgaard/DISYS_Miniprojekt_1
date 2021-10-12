package main

type fork struct {
	isTaken       int
	id            int
	numberOfUse   int
	inputChannel  chan int
	outputChannel chan int
}

func newFork(id int) *fork {
	f := fork{id: id}
	f.isTaken = 0
	f.numberOfUse = 0
	f.inputChannel = make(chan int)
	f.outputChannel = make(chan int)
	return &f
}

func (self *fork) pickedUp() {
	self.isTaken = 1
}

func (self *fork) putDown() {
	self.isTaken = 0
	self.numberOfUse++
}

func (self *fork) exist() {
	for {
		self.readInput()
	}
}

func (self *fork) readInput() {
	var input = <-self.inputChannel
	if input == 0 {
		self.outputChannel <- self.numberOfUse
	} else if input == 1 {
		self.outputChannel <- self.isTaken
	}
}
