package sender

import (
	"github.com/urlooker/alarm/g"
)

var (
	SmsWorkerChan  chan int
	MailWorkerChan chan int
	IMWorkerChan chan int
)

func Init() {
	workerConfig := g.Config.Worker
	SmsWorkerChan = make(chan int, workerConfig.Sms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
	IMWorkerChan = make(chan int, workerConfig.IM)

	Consume()
}

func Consume() {
	go ConsumeMail()
	go ConsumeSms()
	go ConsumeIM()
}
