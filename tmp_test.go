package main

import (
	"fmt"
	"log"
	"testing"
)

func TestTmp(t *testing.T) {
	log.Println("test tmp")
	if []byte("aaa")[0] == []byte("ccc")[0] {
		fmt.Println([]byte("aaa")[0])
		fmt.Println([]byte("ccc")[0])
		fmt.Println("ok")
	}
}
