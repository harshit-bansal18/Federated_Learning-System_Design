package main

import (
	"fmt"
	"log"
	"math/rand"

	//"math/rand"
	"time"

	"dev.com/system/messages"
	"github.com/asynkron/protoactor-go/actor"
)

type TrainerActor struct{
	SelectorPid *actor.PID

}

func (t *TrainerActor) Receive(ctx actor.Context){
	switch ctx.Message().(type){
	
	case *actor.Started:
		log.Println("I am trainer\n",ctx.Actor())
		
	case *messages.PingSelector:
		log.Println("Trainer:: Pinging Selector...")
		time.Sleep(5*time.Second)
		n := rand.Intn(2)
		//time.Sleep(time.Duration(rand.Int31n(30)* int32(time.Second)))
		if n==1 {
			ctx.Request(t.SelectorPid, &messages.AttemptConnection{
				Properties: "My Properties",
			})
		}
		
	case *messages.TrainModel:
		log.Println("Trainer:: Training Model...")
		time.Sleep(10*time.Second)
		ctx.Respond(&messages.LocalTrainedModel{
			Model: "This is the new model!",
		})
	case *messages.ConnectionAccepted:
		log.Println("Trainer:: Connected to FL Server! Preparing for the round...")
	
	case *messages.ConnectLater:
		log.Println("Trainer", ctx.Message().(*messages.ConnectLater).Time)

	case *messages.ConnectionRefused:
		log.Println("Trainer", ctx.Message().(*messages.ConnectionRefused).Reason)
	
	case *messages.StopTraining:
		log.Println("Trainer::", ctx.Message().(*messages.StopTraining).Message)
	
	default:
		fmt.Print("Trainer: ")
		fmt.Println(ctx.Message())
	}
}