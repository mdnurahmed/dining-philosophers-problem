# dining-philosophers-problem

## Problem Description 

It was originally formulated in 1965 by Edsger Dijkstra as a student exam exercise. The Dining Philosopher Problem states that 5 philosophers seated around a circular table with one chopstick between each pair of philosophers. There is one chopstick between each philosopher. A philosopher may eat if he can pick up the two chopsticks adjacent to him. One chopstick may be picked up by any one of its adjacent followers but not both. Each philosopher must alternately think and eat. 

![dining-philosophers-problem](https://github.com/mdnurahmed/dining-philosophers-problem/blob/main/dining_philosopher_problem.png)

## Conditions 

- Only one philosopher can hold a chopstick at a time.
- It must be impossible for a deadlock to occur.
- It must be impossible for a philosopher to starve waiting for a chopstick.
- It must be possible for more than one philosopher to eat at the same time.

## My Apporach

Each philosopher can keep doing three things in this serial :

- Eat 
- Drop chopstick
- Think

We can simulate chopstick using channels.Restaurent object serves the chopsticks at first . Each philosopher waits for the required chopsticks . When he is done eating he returns the chopsticks (by sending true signal). 

## Why there will be no deadlock 
Cause in the begining at least 1 phiosophers gets to eat meaning he has both the chopsticks .Other chopsticks are acquire by others by that time . So there will never be a time when each philosopher just has one chopstick and waiting for the other one . 

## Why there will be no starvation 
Because whenever a goroutine blocks the GO scheduler runs an awaiting runnable goroutine putting the blocked goroutine in a queue. There will never be a time when all goroutines are asleep in this solution . 
