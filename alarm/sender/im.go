package sender

import (
	"log"
	"strings"
	"time"
	"github.com/toolkits/net/httplib"
	"github.com/urlooker/alarm/g"
)

func ConsumeIM() {
	queue := g.Config.Queue.Mail
	for {
		L := PopAllIM(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendIMList(L)
	}
}

func SendIMList(L []*g.IM) {
	for _, im := range L {
		if im.Tos == "" || im.Tos == "," || im.Tos == ";" || im.Content == "" {
			continue
		}

		toArr := strings.Split(im.Tos, ",")
		log.Println("IMCount", len(toArr))

		IMWorkerChan <- 1
		go SendIM(im)
	}
}

func SendIM(im *g.IM) {
	defer func() {
		<-IMWorkerChan
	}()

	url := "http://127.0.0.1:4567/send"
	r := httplib.Post(url).SetTimeout(5*time.Second, 2*time.Minute)
	tos := "DingZiAn"
	r.Param("tos", tos)
	r.Param("content", im.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	if g.Config.Debug {
		log.Println("==im==>>>>", im)
		log.Println("<<<<==im==", resp)
	}

}
