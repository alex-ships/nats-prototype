package main
import (
	"fmt"
	"github.com/nats-io/nats.go"
)
func main(){
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Publish("Orders", []byte("Hello World"))

	fmt.Println("mesage sent")
}