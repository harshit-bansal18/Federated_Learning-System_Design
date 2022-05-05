package main

import (
	"fmt"
	"log"
	"time"

	"dev.com/system/messages"
	"github.com/asynkron/protoactor-go/actor"
)


type AggregateActor struct{
	Region string

	AssignedSelector *actor.PID
	connectedDevices [] *actor.PID
	receivedModels [] string 
}



func (a *AggregateActor) Receive(ctx actor.Context){
	switch ctx.Message().(type){

	case *actor.Started:
		log.Println("I am Aggregator ", ctx.Actor())
		a.connectedDevices = nil
		a.receivedModels = nil 

		time.Sleep(15*time.Second)

		future:= ctx.RequestFuture(a.AssignedSelector, &messages.ReturnDevices{}, 25*time.Second)
		res, err := future.Result()
		
		if err != nil{
			log.Println("Selector Failed To Return Devices")
			ctx.Request(ctx.Parent(), &messages.NotEnoughDevices{})
			time.Sleep(time.Minute)
		}
		fmt.Println("Received Devices")
		a.connectedDevices = res.(messages.AvailableDevices).Devcies
		time.Sleep(30*time.Second)
		// now ping devices for model
		//concurrency required here
		for i:=0; i < len(a.connectedDevices); i++{
			f := ctx.RequestFuture(a.connectedDevices[i], &messages.TrainModel{
				FlPlan: "Federated Learning Plan",
			}, time.Minute)
			r, e := f.Result()
			if e != nil{
				log.Panicln("Couldn't receive Model from the device")
				ctx.Send(a.connectedDevices[i], &messages.StopTraining{
					Message: "Failed to received Model",
				})
				continue		
			}
			log.Println("Aggregator::Device returned Model!")
			a.receivedModels = append(a.receivedModels, r.(*messages.LocalTrainedModel).Model)
			ctx.Send(a.connectedDevices[i], &messages.StopTraining{
				Message: "Model received Successfully!",
			})
		}

		ctx.Request(ctx.Parent(), &messages.LocalTrainedModels{
			Models: a.receivedModels,
		})
		log.Println("Aggregator:: Job Done!")

	case *messages.LocalTrainedModel:
		fmt.Println("Received Model from ", ctx.Sender())
		//a.receivedModels[ctx.Sender()] = ctx.Message().(*messages.LocalTrainedModel).Model
		a.receivedModels = append(a.receivedModels, ctx.Message().(*messages.LocalTrainedModel).Model)
		ctx.Respond(&messages.StopTraining{
			Message: "Successfully Received the model!",
		})

	case *actor.Stopping:
		log.Println("Shutting Down Aggregator....")

	default:
		log.Print("Default :")
		log.Print(ctx.Message())
	}
}





