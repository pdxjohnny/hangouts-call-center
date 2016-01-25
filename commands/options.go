package commands

// ConfigOptions is used to set viper defaults
var ConfigOptions = map[string]interface{}{
	"get": map[string]interface{}{
		"collection": map[string]interface{}{
			"value": "",
			"help":  "The collection to look in",
		},
		"id": map[string]interface{}{
			"value": "",
			"help":  "The id's doc to return",
		},
	},
	"save": map[string]interface{}{
		"collection": map[string]interface{}{
			"value": "",
			"help":  "The collection to put in",
		},
	},
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
		"static": map[string]interface{}{
			"value": "static",
			"help":  "Directory which holds static content",
		},
	},
	"db": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 42345,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/db/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/db/key.pem",
			"help":  "Key for https server",
		},
	},
	"user": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 32345,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/user/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/user/key.pem",
			"help":  "Key for https server",
		},
	},
}
