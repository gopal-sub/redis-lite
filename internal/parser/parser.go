package parser

import (
	"redis-lite/internal/protocol"
	"strings"
)


func Parser (message string) (*protocol.COMMAND, error){
	command := &protocol.COMMAND{}
	
	FirstSpace := strings.Index(message, " ")
	

	if FirstSpace == -1 {
		command.Type = protocol.COMMANDTYPE(message[:])
		command.Arg = ""
	} else {
		//there is a possiblity of index out of range here fix this
		command.Type = protocol.COMMANDTYPE(message[:FirstSpace])
		// trim space will remove all space
		command.Arg = strings.TrimSpace(message[FirstSpace+1:])
		
	}
	ValidCommand := true
	if _, exists := protocol.ValidCommand[command.Type]; exists {
		if command.Type == protocol.SET{
			// SET needs to be SET key value -> arg needs to have a space in bw
			//cannot set a blank val
			isValidGet := strings.Index(command.Arg, " ")
			if isValidGet == -1 {
				ValidCommand = false
			}
		}
		if command.Type != protocol.ALLSTORES && command.Arg == "" {
			ValidCommand = false
		}
	}
	
	if !ValidCommand{
		return nil, INVALIDMESSAGE
	}
	return command, nil
}