package main

import (
	"fmt"
	// "log"
	// "os"
	// "os/signal"
	// "time"
	// "github.com/asynkron/protoactor-go/actor"
)

// type ping struct{
// 	data string
// }

// type pong struct{
// 	data string
// }

// type pingActor struct{
// 	pongPid *actor.PID
// }


// func ( p *pingActor) Receive(ctx actor.Context){
// 	switch ctx.Message().(type){
// 	// case *actor.Started:
// 	// 	f := ctx.RequestFuture(p.pongPid, &ping{ data: "trials"}, 2*time.Second)
// 	// 	r, e:= f.Result()
// 	// 	fmt.Println("Result, error")
// 	// 	log.Println(r, e)

// 	case *ping:
// 		f := ctx.RequestFuture(p.pongPid, &ping{ data: "trials"}, 2*time.Second)
// 		r, e:= f.Result()
// 		fmt.Println("Result, error")
// 		log.Println(r, e)

// 		ctx.Request(p.pongPid, &ping{
// 			data: "Sent to pong",
// 		})
		
// 	case *pong:
// 		log.Print("Received Pong Message!")
// 		log.Print(ctx.Message().(*pong).data)
// 		time.Sleep(1*time.Second)
// 		ctx.Respond(&ping{
// 			data: "ping back",
// 		})
// 	}
// }


func main(){
	fmt.Printf("Basic Actor system\n")
	// system := actor.NewActorSystem()
	// defer system.Shutdown()
	// pongProps := actor.PropsFromFunc(func(ctx actor.Context){
	// 	switch ctx.Message().(type){
		
	// 	case *ping:
	// 		msg := ctx.Message().(*ping)
	// 		fmt.Println("trying to get values")
	// 		fmt.Println(msg.data)
	// 		time.Sleep(time.Second)
	// 		ctx.Respond(&pong{data: "reply"})
	// 	case *pong:
	// 		fmt.Println("Inside pong after future request")
	// 		msg:= ctx.Message().(*pong)
	// 		fmt.Println(msg)

	// 	default:
		
	// 	}
		
	// })

	// pongPid := system.Root.Spawn(pongProps)

	// pingProps := actor.PropsFromProducer(func() actor.Actor{
	// 	return &pingActor{
	// 		pongPid: pongPid,
	// 	}
	// })

	// pingPid := system.Root.Spawn(pingProps)

	// finish := make(chan os.Signal, 1)
	// signal.Notify(finish, os.Interrupt, os.Kill)
	// ticker := time.NewTicker(1*time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <- ticker.C:
	// 		system.Root.Send(pingPid, &ping{
	// 			data: "Aggregate Now",
	// 		})
	// 	case <- finish:
	// 		log.Print("finish")
	// 		return
	// 	}
	// }

	InitializeTop()

}	
