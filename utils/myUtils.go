package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
