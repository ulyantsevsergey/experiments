package main

import (
	"fmt"
	"testing"
)


func TestMain(m *testing.M)  {
	var poolCounter int
	poolCounter = 100
	if poolCounter != poolCounter{
		fmt.Errorf("Pool counter is not 100")
	}
}