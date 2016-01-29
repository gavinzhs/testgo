package main
import (
    "runtime/pprof"
    "os"
    "fmt"
    "time"
    "runtime")

var s = "/Users/zhengsai/Documents/go/src/testgo/pprof_cpu.txt"
var cpuProfile = &s

func main(){
    print("start")
    startCPUProfile()

    time.Sleep(time.Second * 10)

    stopCPUProfile()
    print("end")
}


func startCPUProfile() {
    if *cpuProfile != "" {
        f, err := os.Create(*cpuProfile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s", err)
            return
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
            f.Close()
            return
        }
    }
}

func stopCPUProfile() {
    if *cpuProfile != "" {
        pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
    }
}

var m = "/Users/zhengsai/Documents/go/src/testgo/pprof_mem.txt"
var memProfile = &s
var mr = 512 * 1000
var memProfileRate = &mr

func startMemProfile() {
    if *memProfile != "" && *memProfileRate > 0 {
        runtime.MemProfileRate = *memProfileRate
    }
}

func stopMemProfile() {
    if *memProfile != "" {
        f, err := os.Create(*memProfile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
            return
        }
        if err = pprof.WriteHeapProfile(f); err != nil {
            fmt.Fprintf(os.Stderr, "Can not write %s: %s", *memProfile, err)
        }
        f.Close()
    }
}