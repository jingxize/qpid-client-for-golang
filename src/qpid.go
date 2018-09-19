package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"pack.ag/amqp"
)

//qpid服务端地址（service side address,of course ,it is known as broker）
const qpidservice string = "amqp://192.168.0.3"
//连接qpid服务端的用户名（the username for connect to service side address）
const uname string = "username"
//连接qpid服务的用户名对应密码（the password of the username above）
const password string = "password"
//发送消息时，需要连接qpid服务端对应exchange的地址(the exchange address of service side when send message)
const exchangeaddress string = "qpid.exchange"
//接收消息时，需要连接qpid服务端对应queue的地址(the queue address of service side when receive message)
const receiveaddress string = "qpid.queue"


func main() {
	// Create client
	client, err := amqp.Dial(qpidservice,amqp.ConnSASLPlain(uname, password))
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()
	// Send a message
	{
		// Create a sender
		sender, err := session.NewSender(
			amqp.LinkTargetAddress(exchangeaddress),
		)
		if err != nil {
			log.Fatal("Creating sender link:", err)
		}

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		// Send message
		err = sender.Send(ctx, amqp.NewMessage([]byte("Hello!")))
		if err != nil {
			log.Fatal("Sending message:", err)
		}

		sender.Close(ctx)
		cancel()
	}

	// Continuously read messages
	{
		// Create a receiver
		receiver, err := session.NewReceiver(
			amqp.LinkSourceAddress(receiveaddress),
			amqp.LinkCredit(10),
		)
		if err != nil {
			log.Fatal("Creating receiver link:", err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
			receiver.Close(ctx)
			cancel()
		}()

		for {
			// Receive next message
			msg, err := receiver.Receive(ctx)
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}

			// Accept message
			msg.Accept()

			fmt.Printf("Message received: %s\n", msg.GetData())
		}
	}


}

