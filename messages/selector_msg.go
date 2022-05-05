package messages

import "github.com/asynkron/protoactor-go/actor"

type ConnectLater struct{
	Time string
}

type ConnectionRefused struct{
	Reason string
}

type ConnectionAccepted struct{
	Instructions string
}

type AvailableDevices struct{
	Devcies [] *actor.PID
}