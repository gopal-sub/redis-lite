package store


import(
	"errors"
)


var INVALIKEYORVAL = errors.New("Invalid key or value input")
var KEYDOESNOTEXIST = errors.New("Key does not exist")