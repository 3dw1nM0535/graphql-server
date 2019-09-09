package utils

import (
	log "github.com/3dw1nM0535/go-gql-server/internal/logger"
	"os"
	"strconv"
)

// MustGet returns the env variable or panic if not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("ENV missing, key: " + k)
	}
	return v
}

// MustGetBool returns the env as boolean or panic if not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicf("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}
