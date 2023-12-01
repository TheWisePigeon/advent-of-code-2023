package main

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T){
  calibrationValue, err := getCalibrationValue()
  if err!= nil{
    t.Errorf("Error %q", err)
  }
  fmt.Println("Calibration value is ", calibrationValue)
}
