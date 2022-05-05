package main

import (
	"fmt"
	"log"
	"time"

	"os"
	"os/signal"
	"strconv"
	"dev.com/system/messages"
	"github.com/asynkron/protoactor-go/actor"
)

type TopActor struct{
	Models int
}
var count int

var SelPids [] *actor.PID
var trainerPids [] *actor.PID

func spawnTrainers(sys *actor.ActorSystem){
	
	for i:=0; i<23; i++{
		trainerProps := actor.PropsFromProducer(func() actor.Actor{
			return &TrainerActor{
				SelectorPid: SelPids[i/6],
			}
		})
		tPid ,_ := sys.Root.SpawnNamed(trainerProps, "Trainer"+strconv.Itoa(i))
		trainerPids = append(trainerPids, tPid)
	}
	log.Println("Spawned 23 Trainers")

}

func NewAggregator(ctx actor.Context, t *TopActor){
	childProps := actor.PropsFromProducer(func() actor.Actor{
		return &MasterAggrActor{
			Region: "Asia",
			MinDevices: 10,
			MaxDevices: 25,
	
		}
	})

	childPid,_ := ctx.SpawnNamed(childProps, "Master1")
	ctx.Send(childPid, &messages.InitiateAgg{})

}

func (t *TopActor) Receive(ctx actor.Context){
	switch  ctx.Message().(type){
	case *actor.Started:
		//time.Sleep(10*time.Second)
		// fmt.Print("Attempting to start the round....\n")
		// for i:=0;i<4;i++{
		// 	ctx.Send(SelPids[i], &messages.StartAcceptingConnections{
		// 		Params: "Selection Criterias",
		// 	})
		// }
		// time.Sleep(2*time.Second)
		// NewAggregator(ctx, t)
		log.Println("Orchestrator Spawned!")

	case *messages.SaveNewModel:
		log.Print("New Model Received!!\n")
		fmt.Print(ctx.Message().(*messages.SaveNewModel).Model)
		ctx.Stop(ctx.Children()[0])
		log.Println("Stopping the server...")
		time.Sleep(3*time.Second)
		os.Exit(0)
	case *messages.StartRound:
		fmt.Print("Attempting to start the round....\n")
		for i:=0;i<4;i++{
			ctx.Send(SelPids[i], &messages.StartAcceptingConnections{
				Params: "Selection Criterias",
			})
		}
		for i:=0; i<23; i++{
			ctx.Send(trainerPids[i], &messages.PingSelector{})
		}
		NewAggregator(ctx, t)
		
	case *messages.Abort:
		log.Println("Aborting the round")
		ctx.Stop(ctx.Children()[0])
		log.Println("Stopping the server...")
		time.Sleep(3*time.Second)
		os.Exit(1)
	default:
		fmt.Print("Orchestrator: ")
		fmt.Println(ctx.Message())
	}
}


func InitializeTop(){
	system := actor.NewActorSystem()

	topProps := actor.PropsFromProducer(func() actor.Actor{
		return &TopActor{
			Models: 1,
		}
	})

	
	fmt.Print("Top actor spawned\n")
	//fmt.Print(topPid)
	// Create Selectors
	selProps := actor.PropsFromProducer(func() actor.Actor{
		return &SelectorActor{
			MinDevices: 2,
			MaxDevices: 5,
		}
	}) 
	
	selPid, _ := system.Root.SpawnNamed(selProps,"S1")
	SelPids = append(SelPids, selPid)
	selPid2, _ := system.Root.SpawnNamed(selProps,"S2")
	SelPids = append(SelPids, selPid2)
	selPid3, _ := system.Root.SpawnNamed(selProps,"S3")
	SelPids = append(SelPids, selPid3)
	selPid4, _ := system.Root.SpawnNamed(selProps,"S4")
	SelPids = append(SelPids, selPid4)

	spawnTrainers(system)
	topPid := system.Root.Spawn(topProps)
	fmt.Println(topPid)
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	ticker := time.NewTicker(1*time.Second)
	defer ticker.Stop()
	
	var count int
	count = 0
	// system.Root.Send(topPid, &messages.SaveNewModel{
	// 	Model: "New model",
	// })
	for {
		select {
		case <- ticker.C:
			if(count == 0){
				system.Root.Send(topPid, &messages.StartRound{
					Devices: 25,
				})
				count = 1
			}
		case <- finish:
			
			log.Print("Stopping the server...")
			return
		}
	}
}
