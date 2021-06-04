package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var chopstick [5]chan bool

type Restaurant struct {
}

func (r *Restaurant) Run() {
	for i := 0; i <= 4; i++ {
		log.Printf("chopstick %d is served \n", i)
		chopstick[i] = make(chan bool, 5)
		chopstick[i] <- true
	}
}
func NewRestaurant() *Restaurant {
	return &Restaurant{}
}

type Philosopher struct {
	//ith philosopher needs ith and (i+1)%5 th chopstick to eat
	Id int
}

func (p *Philosopher) Eat() {
	log.Printf("philosopher %d is hungry now & waiting for chopsticks \n", p.Id)
	chopsticksGot := 0
waiting:
	for {
		select {
		case <-chopstick[(p.Id+1)%5]:
			chopsticksGot++
			log.Printf("philosopher %d is holding the chopstick %d\n", p.Id, (p.Id+1)%5)
			if chopsticksGot == 2 {
				log.Printf("philosopher %d got both chopsticks needed \n", p.Id)
				break waiting
			}

		case <-chopstick[p.Id]:
			chopsticksGot++
			log.Printf("philosopher %d is holding the chopstick %d \n", p.Id, p.Id)
			if chopsticksGot == 2 {
				log.Printf("philosopher %d got both chopsticks needed \n", p.Id)
				break waiting
			}
		}
	}
	log.Printf("philosopher %d is eating now \n", p.Id)
	//each philosopher eats for 1 to 3 second
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
}
func (p *Philosopher) Drop_chopstick() {
	log.Printf("philosopher %d is dropping both chopsticks \n", p.Id)
	chopstick[(p.Id+1)%5] <- true
	chopstick[p.Id] <- true

}
func (p *Philosopher) Think() {
	log.Printf("philosopher %d is thinkin now \n", p.Id)
	//each philosopher thinks for 1 to 2 second
	time.Sleep(time.Duration(rand.Intn(1)+1) * time.Second)
}
func (p *Philosopher) Run() {
	for {
		p.Eat()
		p.Drop_chopstick()
		p.Think()
	}
}

func NewPhilosopher(n int) *Philosopher {
	log.Printf("philosopher %d is seated \n", n)
	return &Philosopher{Id: n}
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Generate 5 Philosophers
	go func() {
		for n := 0; n <= 4; n++ {
			p := NewPhilosopher(n)
			go p.Run()
		}
	}()
	// cntrl + C to stop the program

	r := NewRestaurant()
	go r.Run()

	<-c
}
