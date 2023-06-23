package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FuelEconomy struct {
	EngineEfficiency  float32 `json:"EngineEfficiency"`
	DragForce         float32 `json:"DragForce"`
	FuelEnergyDensity float32 `json:FuelEnergyDensity`
}

func readJson(filename string) (float32, float32, float32) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json file", err)
		os.Exit(1)
	}

	var data FuelEconomy
	json.Unmarshal(byteValue, &data)
	return data.EngineEfficiency, data.DragForce, data.FuelEnergyDensity
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error reading argument file")
		fmt.Println("Usage: fec myFile.json ")
		os.Exit(1)
	}
	filename := os.Args[1]
	EffEng, Drag, energyDensity := readJson(filename)

	energyEff := (EffEng * energyDensity) / 100
	amountWork := 1000 * Drag
	kmperlit := energyEff / amountWork
	fmt.Println("km per liter", kmperlit)
	litperhundred := 100 / kmperlit
	fmt.Println("liters per 100km", litperhundred)
}
