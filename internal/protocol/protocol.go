package protocol


type COMMANDTYPE string

//store const
const (
	GET COMMANDTYPE = "GET" //GET key
	SET COMMANDTYPE = "SET"  // set key val -> 2 in args
	DELETE COMMANDTYPE = "DELETE" // DELECT key
	CREATESTORE COMMANDTYPE = "CREATESTORE" // CREATESTORE name
	SELECT COMMANDTYPE = "SELECT" // SELECT stotename
	DROP COMMANDTYPE = "DROP" // DROP storename

	// this would be no args command
	ALLSTORES COMMANDTYPE = "ALLSTORES" //single command
)

//set const

const (

	//set command type ex SADD setname item
	//this would be command onearg
	SCREATE COMMANDTYPE = "SCREATE"
	SDROP COMMANDTYPE = "SDROP"
	SCOUNT COMMANDTYPE = "SCOUNT"
	SEXIST COMMANDTYPE = "SEXIST"
	SALL COMMANDTYPE = "SALL" 


	//this would be command 2 args
	SADD COMMANDTYPE = "SADD" // 2 args
	SDELETE COMMANDTYPE = "SDELETE" // 2 args
	
	//IMPLEMENT THIS LATER IN V2
	SINTERSECT COMMANDTYPE = "SINTERSECT"
	SUNION COMMANDTYPE = "SUNION"
)


type COMMAND struct {
	Type COMMANDTYPE
	Arg string
}


var ValidCommand = map[COMMANDTYPE] struct{}{
	GET: {},
	SET: {},
	DELETE: {},
	CREATESTORE: {},
	SELECT: {},
	ALLSTORES: {},
	
	//SET COMMANDS
	
	SCREATE: {},
	SADD: {},
	SDELETE: {},
	SEXIST: {},
	SDROP: {},
	SALL: {},
}

var NoArgsValidCommand = map[COMMANDTYPE] struct{}{
	ALLSTORES: {},
	SALL: {},
}

var OneArgsValidCommand = map[COMMANDTYPE] struct{}{
	GET: {},
	DELETE: {},
	CREATESTORE: {},
	SELECT: {},
	DROP: {},
	SDROP: {},
	SCREATE: {},
	SCOUNT: {},
	SEXIST: {},
}

var TwoArgsValidCommand = map[COMMANDTYPE] struct{}{
	SET: {},
	SADD: {},
	SDELETE: {},

}
