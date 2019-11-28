package main

import (
	"flag"
	"fmt"
)

func main() {
	testBool := flag.Bool("testBool", false, "start test bool")
	var testBoolVar bool
	flag.BoolVar(&testBoolVar, "testBoolVar", false, "start test boolVar")
	testString := flag.String("testString", "testStringValue", "start test String")
	var testStringVar string
	flag.StringVar(&testStringVar, "testStringVar", "testStringVarValue", "start test StringVar")
	testInt := flag.Int("testInt", -1, "start test Int")
	var testIntVar int
	flag.IntVar(&testIntVar, "testIntVar", -1, "start test IntVar")
	flag.Parse()

	fmt.Printf("testBool:%t \ntestBoolVar:%t \ntestString:%s \ntestStringVar:%s \ntestInt:%d \ntestIntVar:%d\n",
				*testBool, testBoolVar, *testString, testStringVar, *testInt, testIntVar)
}
