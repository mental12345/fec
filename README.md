# Fuel Economy Calculate

This program calculates the fuel economy at a steady cruise. SI units only the input for the program is a json file with the following values: 

```json
{
  "EngineEfficiency": 26,
  "DragForce" : 667.2332,
  "FuelEnergyDensity": 32414328.7844
}
```

Where:
* `EngineEfficiency` is how efficient is the engine.
* `DragForce` is the amount of force the vehicle has to apply. 
  "FuelEnergyDensity": 32414328.7844
* `FuelEnergyDensity` amount of energy in a liter of fuel, example gasoline has around 32414328.8 joules of energy in a liter. 

## How to compile
```bash
go build .  
```

## How to use
fec auto.json