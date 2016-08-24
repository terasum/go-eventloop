package event

import "fmt"

//EventManager 事件管理器
type EventManager struct {
	MsgHandler map[string]EventHandler
	Mchan      chan Massage
}

//NewEventManager 得到一个新的eventManager
func NewEventManager() EventManager {
	var eventManager EventManager
	eventManager.Mchan = make(chan Massage)
	eventManager.MsgHandler = make(map[string]EventHandler)
	return eventManager
}

//Regisiter 注册事件,通过预先注册事件可以对事件进行相应处理
func (em *EventManager) Regisiter(eventName string, handler EventHandler) {
	_, ok := em.MsgHandler[eventName]
	if ok {
		fmt.Printf("事件已经绑定")
	} else {
		em.MsgHandler[eventName] = handler
	}
}

//DeRegisiter 注销事件,注销之后该事件将不被监听
func (em *EventManager) DeRegisiter(eventName string) {
	_, ok := em.MsgHandler[eventName]
	if ok {
		delete(em.MsgHandler, eventName)
	} else {
		fmt.Printf("事件未注册")

	}
}

//Start 开启事件循环
func (em *EventManager) Start() {
	go em.eventLoop()
}

//事件循环
func (em *EventManager) eventLoop() {
	for {
		fmt.Println("hahah ")
		for msg := range em.Mchan {
			handler, ok := em.MsgHandler[msg.MsgType]
			if ok {
				handler.ProcessEvent(&msg)
			} else {
				fmt.Println("该事件未注册")
			}
		}
	}
}
