// package main

// import (
//   "fmt"
//   "github.com/nats-io/nats.go"
//   "log"
//   "time"
// )

// func main() {

//   // Connect to NATS
//   nc, _ := nats.Connect(nats.DefaultURL)

//   // Create JetStream Context
//   js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

//   _, err := js.AddStream(&nats.StreamConfig{
//     Name:      "ORDERS",
//     Subjects:  []string{"ORDERS.*"},
//     Retention: nats.WorkQueuePolicy,
//   })
//   if err != nil {
//     log.Fatal(err)
//   }

//   // Simple Stream Publisher
//   //_, err = js.Publish("ORDERS.*", []byte("hello"))
//   //if err != nil {
//   //  log.Fatal(err)
//   //}

//   // Simple Async Stream Publisher
//   for i := 0; i < 500; i++ {
//     js.PublishAsync("ORDERS.scratch", []byte("hello"))
//   }

//   // Simple Async Ephemeral Consumer
//   sub, err := js.QueueSubscribeSync("ORDERS.*", "MONITOR")
//   if err != nil {
//     log.Fatal(err)
//   }

//   for {
//     var msg *nats.Msg
//     msg, err = sub.NextMsg(time.Second)
//     if err != nil {
//       if err == nats.ErrTimeout {
//         time.Sleep(time.Second)
//         continue
//       } else {
//         fmt.Println(err)
//       }
//     }
//     fmt.Println(string(msg.Data))
//   }

//   select {}
// }

package main

import (
	"fmt"
	"log"
	"service-1/kvstore"
	"sync/atomic"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)

	kv := kvstore.NewNATSKVStoreFromConn(nc)

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	_, err := js.AddStream(&nats.StreamConfig{
		Name:      "ORDERS",
		Subjects:  []string{"ORDERS.*"},
		Retention: nats.WorkQueuePolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Simple Stream Publisher
	//_, err = js.Publish("ORDERS.*", []byte("hello"))
	//if err != nil {
	//  log.Fatal(err)
	//}

	// Simple Async Stream Publisher
	for i := 0; i < 5; i++ {
		js.PublishAsync("ORDERS.scratch", []byte("hello"))
	}

	// Simple Async Ephemeral Consumer
	sub, err := js.QueueSubscribeSync("ORDERS.*", "MONITOR")
	if err != nil {
		log.Fatal(err)
	}

	// Create a pool goroutines
	var currentPoolSize atomic.Int64
	chs := make([]chan struct{}, 0)

	go func() {
		// this endless loop is responsible for pool size control
		for {
			var configPoolSize int64
			atomic.StoreInt64(&configPoolSize, kv.GetInt64("pool-size"))
			fmt.Println("Current pool size:", currentPoolSize.Load())

			if currentPoolSize.Load() < configPoolSize {
				currentPoolSize.Add(1)
				ch := make(chan struct{}, 1)
				chs = append(chs, ch)
				// increase here
				go func(ch chan struct{}) {
					// this endless loop is responsible for consuming messages
					for {
						select {
						case <-ch:
							return
						default:
							var msg *nats.Msg
							msg, err = sub.NextMsg(time.Second)
							if err != nil {
								if err == nats.ErrTimeout {
									time.Sleep(time.Second)
									continue
								} else {
									fmt.Println(err)
								}
							}
							fmt.Println(string(msg.Data))
						}
					}
				}(ch)
			} else if currentPoolSize.Load() > configPoolSize {
				if len(chs) > 0 {
					ch := chs[len(chs)-1]
					ch <- struct{}{}
					close(ch)
					currentPoolSize.Add(-1)
					chs = chs[:len(chs)-1]
				}
			} else {
				time.Sleep(time.Second)
			}
		}
	}()

	select {}
}
