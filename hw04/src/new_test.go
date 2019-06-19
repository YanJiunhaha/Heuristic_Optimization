package main

import (
    "testing"
    "math/rand"
    "time"
    "../../object_function"
)


func TestDecoder(t *testing.T) {
    rand.Seed(time.Now().Unix())
    var v uint64 = rand.Uint64()
    x, y := Decoder(v)
    t.Log(x,y)
    t.Log("success")
}

func TestStruct(t *testing.T){
    rand.Seed(time.Now().Unix())
    v := Vertex{2,2,0,rand.Uint64()}
    v.X, v.Y = Decoder(v.Chromosome)
    v.Value = object.Result(v.X,v.Y)
    t.Logf("Vertex:\nx:%f\ny:%f\nvalue:%f\nchromosome:%d",v.X,v.Y,v.Value,v.Chromosome);
}

func TestSelect(t *testing.T) {
    var test [N]Vertex
    for i := 0; i < N; i++ {
        test[i].Value = float64(i)
    }
    p := Select(test)
    var sum float64
    for _, i := range p {
        sum += i
    }
    t.Log(p)
    t.Log(sum)
}

func TestWheel(t *testing.T){
    var test,old [N]Vertex
    for i := 0; i < N; i++ {
        test[i].Value = float64(i)
    }
    p := Select(test)
    Wheel(test,&old,p)
    t.Log(test)
    t.Log(old)
}
