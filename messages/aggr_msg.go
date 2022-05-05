package messages



type SendModelToMaster struct{
	
}

type SendFLPlanToDevice struct{

}

type ReturnDevices struct {

}

type StopTraining struct{
	Message string
}

type TrainModel struct{
	FlPlan string
}

type LocalTrainedModels struct{
	Models [] string
}
type NotEnoughDevices struct{
	
}
