# Basic Cloud Server Model For Federated Learning
## Usage
To run the server, open the folder in terminal, run the following commands.

```
go mod tidy
go build
./system
```

## Dependencies

* Go 1.18 (Preferred)
* ProtoActor-Go

## Output

The system will output log of all the server activity going on. I have used randomness in trainers to generate different 
conditions on each run. The parameters, wherever required, are hardcoded and can be changed to observe the resultant behaviour.
