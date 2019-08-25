package main

import (
    "fmt"
    "../../object_function"
    "math/rand"
    "time"
)

func Direct(x float64, y float64, step float64) (dx, dy, best float64) {
    best = 0
    for i := -1.; i <= 1.; i += 2. {
        for j := -1.; j <= 1.; j += 2. {
            r := object.Result(x + i * step, y + j * step)
            if r > best {
                best = r
                dx = i
                dy = j
            }
        }
    }
    return
}

func main() {
    // rand.Seed(4)
    rand.Seed(time.Now().Unix())
    var x, y, value float64
    x = rand.Float64() * 10 - 5
    y = rand.Float64() * 10 - 5
    value = object.Result(x,y)
    fmt.Printf("Init (x,y)=(%f,%f)\n",x,y)

    var stepSize float64 = 0.5

    for stepSize > 1e-5 {
        nx, ny, nv := Direct(x,y,stepSize)
        if nv > value {
            x += nx * stepSize
            y += ny * stepSize
            value = nv
            stepSize *= 2
            fmt.Println(value)
        } else {
          stepSize *= 0.5
        }
    }

    fmt.Printf("F(%f, %f)=%f\n",x,y,value)
}

