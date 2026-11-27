package parser

import (
	"fmt"
	"redis-lite/internal/protocol"
	"strings"
)


func Parser (message string) (*protocol.COMMAND, error){
	command := &protocol.COMMAND{}
	
	FirstSpace := strings.Index(message, " ")
	
	// this would be for listing all sets or stores
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

	fmt.Println("command")
	fmt.Println(command)
	if _, exists := protocol.ValidCommand[command.Type]; exists {
		if _, NoArgsValid := protocol.NoArgsValidCommand[command.Type]; NoArgsValid{
			if command.Arg != ""{
				ValidCommand = false
			}
		}else if _, OneArgsValid := protocol.OneArgsValidCommand[command.Type]; OneArgsValid{
			spaceArg := strings.Index(command.Arg, " ")
			if spaceArg != -1 {
				ValidCommand = false
			}
		}else if _, TwoArgsValid := protocol.TwoArgsValidCommand[command.Type]; TwoArgsValid{
			spaceArg := strings.Index(command.Arg, " ")
			if spaceArg == -1 {
				ValidCommand = false
			}
		}else {
			ValidCommand = false
		}
	}
	fmt.Println(ValidCommand)
	if !ValidCommand{
		return nil, INVALIDMESSAGE
	}
	return command, nil
}