package sender

import (
	"encoding/json"
	"log"

	"github.com/garyburd/redigo/redis"

	"github.com/urlooker/alarm/g"
)

func PopAllSms(queue string) []*g.Sms {
	ret := []*g.Sms{}

	rc := g.RedisConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var sms g.Sms
		err = json.Unmarshal([]byte(reply), &sms)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &sms)
	}

	return ret
}

func PopAllIM(queue string) []*g.IM {
	ret := []*g.IM{}

	rc := g.RedisConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var im g.IM
		err = json.Unmarshal([]byte(reply), &im)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &im)
	}

	return ret
}

func PopAllMail(queue string) []*g.Mail {
	ret := []*g.Mail{}

	rc := g.RedisConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var mail g.Mail
		err = json.Unmarshal([]byte(reply), &mail)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &mail)
	}

	return ret
}
