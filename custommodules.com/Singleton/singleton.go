package singleton

//Non Exportable
type singleton struct {
	value int
}

//Exportable
var Singleton = singleton{value: 2}

// Only uppercased stuff is exportable
