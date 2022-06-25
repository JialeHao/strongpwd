package main

import (
	"flag"
	"os"

	"github.com/JialeHao/strongpwd/generator"
)

var length = flag.Int("l", 12, "The length of the generated password.(max 1024)")

func main() {
	flag.Parse()
	bf := generator.GenerateComplexPassword(*length)
	bf.WriteTo(os.Stdout)
}