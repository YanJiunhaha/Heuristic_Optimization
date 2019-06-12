package main

import (
    "fmt"
    //"time"
    "math"
    "math/rand"
    "../../object_function"
)

const(
    MG int = 50           // Maximal Number of Generations
    N  int = 10           // Population Size
    CL int = 32           // Number of bits in each chromsome
    SF float64 = 2.0      // Selection Factor
    CR float64 = 0.5      // Crossover Rate
    MR float64 = 0.05     // Mutation Rate
)

// data declare
var(
    c [N][CL]int
    f [N]float64
    best_c [CL]int
    best_f float64
)

func RAND() float64{
    return rand.Float64()
}

// Decode Chromosome
func decoder(chromosome [CL]int, x *float64, y *float64){
    *x = 0

    for i := 0; i < CL / 2; i++ {
        *x = (*x) * 2 + float64(chromosome[i])
    }
    *x = (*x) / (1 << 16) * 10 - 5

    *y = 0
    for i := CL / 2; i < CL; i++ {
        *y = (*y) * 2 + float64(chromosome[i])
    }
    *y = (*y) / (1 << 16) * 10 - 5
}

func main(){
    // rand seed
    rand.Seed(4)
    var x,y float64

    // Initialize Population
    best_f := -1.0e99
    for i := 0; i < N; i++ {
        // Randomly set each gene to '0' or '1'
        for j := 0; j < CL; j++ {
            c[i][j] = rand.Intn(2)
        }
    }

    // Repeat Genetic Algorithm cycle for MG times
    for gen := 0; gen < MG; gen++ {
        // Evaluation
        for i := 0; i < N; i++ {
            decoder(c[i], &x, &y)
            f[i] = object.Result(x, y)

            // Update best solution
            if f[i] > best_f {
                best_f = f[i]
                for j := 0; j < CL; j++ {
                    best_c[j] = c[i][j]
                }
            }
            // Selection
            // Evaluate Selection Probability
            tmpf := 0.0
            var p [N]float64
            var tmpc [N][CL]int
            for i := 0; i < N; i++ {
                p[i] = math.Pow(f[i], SF)
                tmpf += p[i]
            }
            for i := 0; i < N; i++ {
                p[i] /= tmpf
            }
            // Retain the best Chromosome found so far
            for i := 0; i < CL; i++ {
                tmpc[0][i] = best_c[i]
            }
            // Roulette wheel selection with replacment
            for i := 0; i < N; i++ {
                tmpf = RAND()
                var k int
                for k = 0; tmpf > p[k]; k++ {
                    tmpf -= p[k]
                }
                // Chromosome j is selected
                for j := 0; j < CL; j++ {
                    tmpc[i][j] = c[k][j]
                }
            }
            // Copy temporary population to population
            for i := 0; i < N; i++ {
                for j := 0; j < CL; j++ {
                    c[i][j] = tmpc[i][j]
                }
            }
            // 1-site Crossover
            tmpi := 0
            var site int
            for i := 0; i < N; i += 2 {
                if RAND() < CR {
                    site = int(RAND() * float64(CL))
                    for j := 0; j < site; j++{
                        tmpi = c[i][j]
                        c[i][j] = c[i + 1][j]
                        c[i + 1][j] = tmpi
                    }
                }
            }
            // Mutation
            for i := 0; i < N; i++ {
                for j := 0; j < CL; j++ {
                    if RAND() < MR {
                        c[i][j] = 1 - c[i][j]
                    }
                }
            }
            fmt.Printf("%f\n", best_f)
        }
        decoder(best_c, &x, &y)
        fmt.Printf("F(%f,%f) = %f\n", x, y ,object.Result(x, y))
    }
}
