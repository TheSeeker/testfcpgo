// testfcp project main.go
package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
		foo, err := fcp.NewClient("127.0.0.1:9481", false, "LOL Test")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		foo.Connect()
		time.Sleep(3 * time.Second)

		upload := foo.NewClientPut()
		upload.SetURI("SSK@")
		upload.SetUploadFrom("disk")
		upload.SetFilename("f:\\curstats.txt")
		verbosity := byte(1 | 64 | 128)
		println(verbosity)
		upload.SetVerbosity(verbosity)
		upload.SetIdentifier("testing123")

		foo.DoInsert(upload)

		time.Sleep(5 * time.Minute)
		err = foo.Disconnect()
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(3 * time.Second)

}
