package variables

const (
	// APIPathLoginServer authenticates a client
	APIPathLoginServer = "/login"
	// APIPathLogin is the path to login
	APIPathLogin = "/api" + APIPathLoginServer
	// APIPathRefreshServer updates a clients token
	APIPathRefreshServer = "/refresh"
	// APIPathRefresh is the path to login
	APIPathRefresh = "/api" + APIPathRefreshServer
	// APIPathCallServer is the path to make a call
	APIPathCallServer = "/call/:number"
	// APIPathCall is the path to make a call
	APIPathCall = "/api" + APIPathCallServer
	// APIPathEndServer is the path to end a call
	APIPathEndServer = "/end/:lock"
	// APIPathEnd is the path to end a call
	APIPathEnd = "/api" + APIPathEndServer
	// APIPathCallerServer is the path callers connect to via websocket
	APIPathCallerServer = "/caller"
	// APIPathCaller is the path callers connect to via websocket
	APIPathCaller = "/api" + APIPathCallerServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
