package variables

const (
	// APIPathCallServer is the path to make a call
	APIPathCallServer = "/call/:number"
	// APIPathCall is the path to make a call
	APIPathCall = "/api" + APIPathCallServer
	// APIPathEndServer is the path to end a call
	APIPathEndServer = "/end/:lock"
	// APIPathEnd is the path to end a call
	APIPathEnd = "/api" + APIPathEndServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
