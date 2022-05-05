package main

import (
	"fmt"
	"log"
	"time"

	"dev.com/system/messages"
	"github.com/asynkron/protoactor-go/actor"
)

type SelectorActor struct{
	MinDevices int
	MaxDevices int
	AcceptDevice bool
	selectionCriteria string
	devices [] *actor.PID
}

func checkDeviceElegibility(s *SelectorActor,pid *actor.PID) bool{
	//code to check eligibility
	if len(s.devices) >= s.MaxDevices{
		return false
	}
	return true
}

func (s *SelectorActor) Receive(ctx actor.Context){
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Println("Starting the selector.....")
		//time.Sleep(10*time.Second)
	case *messages.StartAcceptingConnections:
		log.Println(ctx.Message().(*messages.StartAcceptingConnections).Params)
		s.selectionCriteria = ctx.Message().(*messages.StartAcceptingConnections).Params
		s.AcceptDevice = true
		s.devices = nil
	case *messages.AttemptConnection:
		if s.AcceptDevice == false {
			
			ctx.Respond(&messages.ConnectLater{
				Time: "Better Luck Next Time :(",
			
			})
		} else{
			
			if(checkDeviceElegibility(s, ctx.Sender())){
				s.devices = append(s.devices, ctx.Sender())
				ctx.Respond(&messages.ConnectionAccepted{
					Instructions: "Execute these instructions on device to prepare for round.",
				
				})
			} else{
				
				ctx.Respond(&messages.ConnectionRefused{
					Reason: "You are not eligible to participate in the round.",
				
				})
			}
		}
	
	case *messages.ReturnDevices:
		
		var t time.Duration
		for len(s.devices) < s.MinDevices && t < time.Minute{
			t += time.Microsecond
		}
		if(len(s.devices) > s.MinDevices){
			ctx.Respond(messages.AvailableDevices{
				Devcies: s.devices,
			})
		} else{
			// stop this selector
			s.AcceptDevice = false
			s.devices = nil
			fmt.Println("Not enough devices yet!")
		}
		
	
	default:
		fmt.Print("Selector: ")
		fmt.Println(ctx.Message())
		
	}
}