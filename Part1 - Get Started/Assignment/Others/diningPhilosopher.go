package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number          int
	host            *Host
	leftCS, rightCS *ChopS
	myChannel       chan int
}

func (p *Philo) eat(host *Host) {
	for i := 0; i < 3; i++ {
		go host.philoWantToEat(p)

		// wait for approval from the host on my channel
		<-p.myChannel

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d on the %d time \n", p.number, i+1)

		fmt.Printf("finishing eating %d on the %d time \n", p.number, i+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		host.philoFinishedToEat(p)
	}
}

type Host struct {
	waitingList   []*Philo
	currentEating []int
	lock          sync.Mutex
}

func (host *Host) philoWantToEat(p *Philo) {
	host.lock.Lock()

	// check if we can allow the Philo to eat right away
	if host.canEat(p) {
		// signal to the Philo that he can start eating
		host.currentEating = append(host.currentEating, p.number)
		p.myChannel <- 1
	} else {
		// add Philo to the waiting list
		host.waitingList = append(host.waitingList, p)
	}

	host.lock.Unlock()
}

func (host *Host) canEat(p *Philo) bool {

	return len(host.currentEating) == 0 ||
		(len(host.currentEating) == 1 &&
			!contains(host.currentEating, (p.number+4)%5) && // prev neighbor
			!contains(host.currentEating, (p.number+1)%5)) // next neighbor
}

func (host *Host) philoFinishedToEat(p *Philo) {
	host.lock.Lock()

	host.currentEating = removeFromSlice(host.currentEating, p.number)

	// look for next one that can eat
	for _, element := range host.waitingList {
		if host.canEat(element) {
			host.waitingList = removeFromSliceP(host.waitingList, *element)

			// signal to the Philo that he can start eating
			host.currentEating = append(host.currentEating, element.number)
			element.myChannel <- 1
		}
	}
	host.lock.Unlock()
}

func removeFromSliceP(slice []*Philo, value Philo) []*Philo {
	i := indexOfP(slice, value)
	if i == -1 {
		return slice
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func indexOfP(slice []*Philo, value Philo) int {
	for p, v := range slice {
		if v.number == value.number {
			return p
		}
	}
	return -1
}

func removeFromSlice(slice []int, value int) []int {
	i := indexOf(slice, value)
	if i == -1 {
		return slice
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func indexOf(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	host := new(Host)

	chopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, host, chopSticks[i], chopSticks[(i+1)%5], make(chan int)}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat(host)
	}

	time.Sleep(3 * time.Second)
}

// ------------------------------- Another one ----------------------------
package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var ch_began, ch_finish chan bool

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	number int
}

func get_permission(){
	 ch_began <- true
}

func (p *Philo) eat() {
	get_permission()
  p.leftCS.Lock()
  p.rightCS.Lock()
  fmt.Println("Starting to eat", p.number)
  fmt.Println("Finishing to eat", p.number)
  p.rightCS.Unlock()
  p.leftCS.Unlock()
  wg.Done()
  ch_finish <- true
}

func host(){
	philos_eating := 0;
	for {
		// Max philos eating at the same time (2)
		if philos_eating == 2 {
			<- ch_finish
		 	philos_eating --
		}
 		select {
      case <- ch_finish:
          philos_eating--
      case <- ch_began:
          philos_eating++
    }
	}
}


func main() {
	ch_began = make(chan bool)
	ch_finish = make(chan bool)
	// number of philos
	max:= 5
	CSticks := make([]*ChopS, max)
	wg.Add(max*3)
	for i:=0; i<max;i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, max)
	for i:=0; i<max;i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		if rand.Intn(2) == 0 {
			philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i+1}
		} else {
			philos[i] = &Philo{CSticks[(i+1)%5], CSticks[i], i+1}
		}
	}
	go host()
	for j:=0; j<3;j++{
		for i:=0; i<max; i++ {
			go philos[i].eat()
		}
	}
	wg.Wait()
}