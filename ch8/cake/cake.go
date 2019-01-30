

// Package cake provides a simulation of
// a concurrent cake shop with numerous parameters.

// Use this command to run the benchmarks:
// $ go test -bench=. ./ch8/cake

package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // Number of cakes to bake
	BakeTime       time.Duration // time to bake one cake
	BakeStdDev     time.Duration // Standard deviation of one cake
	BakeBuf        int           // buffer slots between banking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration // time to ice one cake
	IceStdDev      time.Duration // standard deviation of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation of inscribing time
}

type cake int

func (s *Shop) baker(baked chan<- cake)  {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev) // baking time
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, backed <-chan cake) {
	for c := range backed {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev) // icing time
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev) // inscribing time
		if s.Verbose {
			fmt.Println("Finished", c)
		}
	}
}



// Work runs the simulation 'runs' times
func (s *Shop) Work(runs int)  {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}


// work blocks the calling goroutine for a period of time
// that is normally distributed around d
// with a standard deviation of stddev
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
