package event

import "fmt"

//EventManager 事件管理器
type EventManager struct {
	msgHandler map[string]EventHandler
	mchan      chan Massage
}

//NewEventManager 得到一个新的eventManager
func NewEventManager() EventManager {
	var eventManager EventManager
	eventManager.mchan = make(chan Massage)
	eventManager.msgHandler = make(map[string]EventHandler)
	return eventManager
}

//Regisiter 注册事件,通过预先注册事件可以对事件进行相应处理
func (em *EventManager) Regisiter(eventName string, handler EventHandler) {
	_, ok := em.msgHandler[eventName]
	if ok {
		fmt.Printf("事件已经绑定")
	} else {
		em.msgHandler[eventName] = handler
	}
}

//DeRegisiter 注销事件,注销之后该事件将不被监听
func (em *EventManager) DeRegisiter(eventName string) {
	_, ok := em.msgHandler[eventName]
	if ok {
		delete(em.msgHandler, eventName)
	} else {
		fmt.Printf("事件未注册")

	}
}

//Start 开启事件循环
func (em *EventManager) Start() {
	go em.eventLoop()
}

//Post 将事件post到该事件管理器
func (em *EventManager) Post(msg Massage) {
	em.mchan <- msg
}

//事件循环
func (em *EventManager) eventLoop() {
	for {
		fmt.Println("hahah ")
		for msg := range em.mchan {
			handler, ok := em.msgHandler[msg.MsgType]
			if ok {
				handler.ProcessEvent(&msg)
			} else {
				fmt.Println("该事件未注册")
			}
		}
	}
}
