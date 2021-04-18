package main

import (
	"rabbit-mq/receive"
	"rabbit-mq/send"
)

func main()  {
	send.Run()
	receive.Run()
}
