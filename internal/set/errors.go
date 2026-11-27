package set


import "errors"


var KEYEXISTS  = errors.New("Key exists")
var KEYDOESNOTEXISTS  = errors.New("Key does not exist")