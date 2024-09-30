package service

import (
	"fmt"
	"time"
)

func MovieWorker() {
	fmt.Println("Mandando pelis")
	time.Sleep(time.Second * 30)
	fmt.Println("tERMINÓ")
}
