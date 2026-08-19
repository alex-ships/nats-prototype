package main
import (
	"fmt"
	"time"
	"github.com/nats-io/nats.go"
)

func main(){
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("Orders", func(m *nats.Msg) {
    fmt.Printf("Received a message: %s\n", string(m.Data))
	})

time.Sleep(time.Hour)
}