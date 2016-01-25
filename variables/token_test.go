package variables

import (
	"errors"
	"os"
	"testing"
)

func TestLoadToken(t *testing.T) {
	os.Setenv(EnvTokenSignKey, "../"+TokenSignKeyDefault)
	os.Setenv(EnvTokenVerifyKey, "../"+TokenVerifyKeyDefault)
	err := LoadTokenKeys()
	if err != nil {
		panic(err)
	}
	if BackendToken == "" {
		panic(errors.New("BackendToken is empty"))
	}
}
