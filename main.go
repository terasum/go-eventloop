package main

import "github.com/terasum/go-eventloop/event"
import "fmt"
import "time"

type tickEventHandler struct {
}

func (tickEH *tickEventHandler) ProcessEvent(msg *event.Massage) {
	fmt.Printf(" %s = %s \n", msg.Content, msg.Time)
}

type helloEventHandler struct {
}

func (heloEH *helloEventHandler) ProcessEvent(msg *event.Massage) {
	fmt.Println("Hello", time.Now())
}

func main() {
	eventManager := event.NewEventManager()
	//注册事件
	eventManager.Regisiter("tick", new(tickEventHandler))
	eventManager.Regisiter("hello", new(helloEventHandler))
	//开启监听线程
	eventManager.Start()

	func() {
		c := time.Tick(1 * time.Second)
		for now := range c {
			massage := event.Massage{
				MsgType: "tick",
				Content: "tick",
				Time:    now,
			}

			eventManager.Post(massage)
			massage2 := event.Massage{
				MsgType: "hello",
				Content: "tick",
				Time:    now,
			}
			eventManager.Post(massage2)

		}

	}()
}
