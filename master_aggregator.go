package main

import (
	"fmt"
	"log"
	"time"

	"dev.com/system/messages"
	"github.com/asynkron/protoactor-go/actor"
)

type MasterAggrActor struct{
	MinDevices int
	MaxDevices int
	Region string
	allModels [] string
	jobdone bool
}

var models [] string


func CreateChildAggregator(name string, context actor.Context, i int){
	props := actor.PropsFromProducer(func() actor.Actor{
		return &AggregateActor{
			Region: "India",
			AssignedSelector: SelPids[i],
		}
	})
	context.SpawnNamed(props, name)
	
}

func AggregateModels(ctx actor.Context, m *MasterAggrActor){
	log.Println("Aggregating the models...")
	time.Sleep(10*time.Second)
	ctx.Request(ctx.Parent(), &messages.SaveNewModel{
		Model: "I am the latest global model!",
	})
}
func (m *MasterAggrActor) Receive(ctx actor.Context){
	switch ctx.Message().(type){
	case *messages.LocalTrainedModels:
		log.Println("Received Models from ", ctx.Sender())
		msg := ctx.Message().(*messages.LocalTrainedModels)
		// for key, element := range msg.Models{
		// 	m.allModels[key] = element
		// }
		fmt.Println("No of models received:", len(msg.Models))
		m.allModels = append(m.allModels, msg.Models...)
		fmt.Println("Current Models in Master Aggregator:: " ,len(m.allModels))

		if(len(m.allModels) > m.MinDevices && !m.jobdone){
			m.jobdone = true
			AggregateModels(ctx, m)
			
		}
	case *messages.NotEnoughDevices:
		log.Println("Master Aggregator:: Devices Insufficient with aggregator")
		ctx.Stop(ctx.Sender())
		
	case *actor.Started:
		log.Println("Master Aggregator:: Spawned")
		m.jobdone = false
		CreateChildAggregator("A1", ctx, 0)
		CreateChildAggregator("A2", ctx, 1)
		CreateChildAggregator("A3", ctx, 2)
		CreateChildAggregator("A4", ctx, 3)
		fmt.Println("Created 4 child aggregators")
		//time.Sleep(30*time.Second)
		// var t time.Duration
		//for len(m.allModels) < m.MinDevices {
		// 	t += 2*time.Microsecond
		//}
		// log.Println("Time elapsed: ", t)
		// if len(m.allModels) < m.MinDevices{
		//  	log.Println("Couldn't Get Enough Models. Aborting...")
		//  	ctx.Request(ctx.Parent(), &messages.Abort{})
		// } else{
		//  	AggregateModels(ctx, m)
		// }


	case *actor.Stopping:
		log.Println("Shutting Down Master Aggregator...")
		ctx.Stop(ctx.Children()[0])
		ctx.Stop(ctx.Children()[1])
		ctx.Stop(ctx.Children()[2])
		ctx.Stop(ctx.Children()[3])

	}
}


