package main

import (
	"../../object_function"
	"fmt"
	"math"
	"math/rand"
    "time"
)

const (
	MG int     = 10000 // 代數
	N  int     = 50   // 群大小
	CL uint    = 64   // 染色體長度
	SF float64 = 2.0  // 選擇率
	CR float64 = 0.5  // 交換率
	MR float64 = 0.05 // 突變率
)

type Vertex struct {
	X          float64
	Y          float64
	Value      float64
	Chromosome uint64
}

func Decoder(chromosome uint64) (x, y float64) {
	left := chromosome >> 32
	right := chromosome & 0xffffffff
	x = float64(left)/(1<<32)*10 - 5
	y = float64(right)/(1<<32)*10 - 5
	return
}

func Decode(c *[N]Vertex) {
    for idx, i := range *c {
        (*c)[idx].X, (*c)[idx].Y = Decoder(i.Chromosome)
        (*c)[idx].Value = object.Result((*c)[idx].X, (*c)[idx].Y)
    }
}

func FindBest(c [N]Vertex, b *Vertex) {
    for _, i := range c {
        if (*b).Value < i.Value{
            (*b) = i
        }
    }
}

func Select(c [N]Vertex) (p [N]float64) {
    var sum float64
    for idx, i := range c {
        p[idx] = math.Pow(i.Value, SF)
        sum += p[idx]
    }
    for idx, i := range p {
        p[idx] = i / sum
    }
    return
}

func Wheel(c [N]Vertex, c_old *[N]Vertex, p [N]float64) {
    for i := 0; i < N; i++ {
        randf := rand.Float64()
        var s int
        for s = 0; randf > p[s]; s++ {
            randf -= p[s]
        }
        c_old[i] = c[s]
    }
}

func Crossover(c *[N]Vertex){
    for idx, i := range *c {
        if rand.Float64() < CR {
            site := uint(rand.Float64() * float64(CL))
            left := i.Chromosome >> site
            right := i.Chromosome << (CL - site)
            (*c)[idx].Chromosome = left + right
        }
    }
}

func Mutation(c *[N]Vertex){
    for idx := 0; idx < N; idx++ {
        if rand.Float64() < MR {
            (*c)[idx].Chromosome = rand.Uint64()
        }
    }
}

func main() {
	rand.Seed(time.Now().Unix())
    //rand.Seed(4)
	var best Vertex
	var c, c_old [N]Vertex
	// init
	for i := 0; i < N; i++ {
		c[i].Chromosome = rand.Uint64()
	}

	for gen := 0; gen < MG; gen++ {
        // decoder
        Decode(&c)
		// best
        FindBest(c, &best)
		// evaluate selection
        p := Select(c)
		// roulette wheel selection with replacment
        Wheel(c, &c_old, p)
		// copy
        c = c_old
		// crossover
        Crossover(&c)
		// Mutation
        Mutation(&c)
        fmt.Println(best.Value)
	}
}
