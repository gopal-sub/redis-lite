package parser

import (
	"redis-lite/internal/protocol"
	"strings"
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
	switch command.Type {
	case protocol.SET, protocol.GET, protocol.DELETE, protocol.SELECT, protocol.CREATESTORE:
		if command.Type == protocol.SET{
			// SET needs to be SET key value -> arg needs to have a space in bw
			isValidGet := strings.Index(command.Arg, " ")
			if isValidGet == -1 {
				ValidCommand = false
			}
		}
		if command.Arg == "" {
			ValidCommand = false
		}
	default:
		ValidCommand = false

	}
	
	if !ValidCommand{
		return nil, INVALIDMESSAGE
	}
	return command, nil
}