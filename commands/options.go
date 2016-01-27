package commands

// ConfigOptions is used to set viper defaults
var ConfigOptions = map[string]interface{}{
	"http": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 8080,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/http/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/http/key.pem",
			"help":  "Key for https server",
		},
		"username": map[string]interface{}{
			"value": "user",
			"help":  "Username to check authenticate requests against",
		},
		"password": map[string]interface{}{
			"value": "pass",
			"help":  "Password to check authenticate requests against",
		},
		"gmail_username": map[string]interface{}{
			"value": "user@gmail.com",
			"help":  "Username to authenticate with",
		},
		"gmail_password": map[string]interface{}{
			"value": "gmailpass",
			"help":  "Password to authenticate with",
		},
	},
	"call": map[string]interface{}{
		"host": map[string]interface{}{
			"value": "http://localhost:8080",
			"help":  "Host that the call center is running on",
		},
		"number": map[string]interface{}{
			"value": "1234567890",
			"help":  "Number to call, will return a lock on that call for ending",
		},
		"username": map[string]interface{}{
			"value": "user",
			"help":  "Username to authenticate with",
		},
		"password": map[string]interface{}{
			"value": "pass",
			"help":  "Password to authenticate with",
		},
	},
	"end": map[string]interface{}{
		"host": map[string]interface{}{
			"value": "http://localhost:8080",
			"help":  "Host that the call center is running on",
		},
		"lock": map[string]interface{}{
			"value": "somelock",
			"help":  "End a call by providing a lock",
		},
		"username": map[string]interface{}{
			"value": "user",
			"help":  "Username to authenticate with",
		},
		"password": map[string]interface{}{
			"value": "pass",
			"help":  "Password to authenticate with",
		},
	},
	"caller": map[string]interface{}{
		"host": map[string]interface{}{
			"value": "http://localhost:8080",
			"help":  "Host that the call center is running on",
		},
		"username": map[string]interface{}{
			"value": "user",
			"help":  "Username to authenticate with",
		},
		"password": map[string]interface{}{
			"value": "pass",
			"help":  "Password to authenticate with",
		},
	},
}
