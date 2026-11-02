package parser

import (
	"strings"
	"redis-lite/internal/protocol"
)


func Parser (message string) (*protocol.COMMAND, error){
	command := &protocol.COMMAND{}
	
	FirstSpace := strings.Index(message, " ")
	

	if FirstSpace == -1 {
		return nil, INVALIDMESSAGE
	} else {
		//there is a possiblity of index out of range here fix this
		command.Type = protocol.COMMANDTYPE(message[:FirstSpace])
		// trim space will remove all space
		command.Arg = strings.TrimSpace(message[FirstSpace+1:])
		
	}
	ValidCommand := true
	if !(command.Type == protocol.SET || command.Type == protocol.GET || command.Type == protocol.DELETE) || command.Arg == ""{
		ValidCommand = false
	}
	
	if !ValidCommand{
		return nil, INVALIDMESSAGE
	}
	return command, nil
}