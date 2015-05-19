# intemp

# namespace 

**Documentation:** [![GoDoc](https://godoc.org/github.com/SeanDolphin/intemp?status.png)](http://godoc.org/github.com/SeanDolphin/intemp)  
**Build Status:** [![Build Status](https://travis-ci.org/SeanDolphin/intemp.svg?branch=master)](https://travis-ci.org/SeanDolphin/intemp)  
**Test Coverage:** [![Coverage Status](https://coveralls.io/repos/SeanDolphin/intemp/badge.svg?branch=master)](https://coveralls.io/r/SeanDolphin/intemp?branch=master)  
**License:**       [![License](http://img.shields.io/:license-mit-blue.svg)](http://opensource.org/licenses/MIT)


intemp creates and cleans up a temp directory.

## Usage
~~~ go
// main.go
package main

import (
	"log"
	"github.com/SeanDolphin/intemp"
)

func main() {
	err := intemp.New(context.Background(),func(ctx context.Context, tempDir string)error{
		//do something with the tempDir
		return nil
	})

	if err != nil{
		log.Fatal(err)
	}
}

~~~

You can also set the path associated with the temp folder

~~~ go
// main.go
package main

import (
	"log"
	"github.com/SeanDolphin/intemp"
)

func main() {
	ctx := context.Background()
	ctx = intemp.SetRoot(ctx, "/somepath")
	err := intemp.New(ctx,func(ctx context.Context, tempDir string)error{
		//do something with the tempDir
		return nil
	})

	if err != nil{
		log.Fatal(err)
	}
}

~~~


