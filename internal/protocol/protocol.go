package protocol


type COMMANDTYPE string


const (
	GET COMMANDTYPE = "GET"
	SET COMMANDTYPE = "SET"
	DELETE COMMANDTYPE = "DELETE"
)

type COMMAND struct {
	Type COMMANDTYPE
	Arg string
}


