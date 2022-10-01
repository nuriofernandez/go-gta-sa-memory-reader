# GTA SA Memory accessor

A simple and beautiful API to read GTA:SA game in memory data ✨

## Usage Example

### Code example

```go
func main() {
    gtaSa := new(gtaSa.GtaSa)
    gtaSa.Hook()
    
    money := gtaSa.GetMoney()
    pedLocation := gtaSa.GetPed().GetLocation()
    vehicleLocation := gtaSa.GetVehicle().GetLocation()
    pedStatus := gtaSa.GetPed().GetStatus()
    
    fmt.Println("Ped status: ", pedStatus)
    fmt.Println("User money: ", money)
    fmt.Printf("Ped location: '%f', '%f', '%f'\n", pedLocation.X, pedLocation.Y, pedLocation.Z)
    fmt.Printf("Vehicle location: '%f', '%f', '%f'\n", vehicleLocation.X, vehicleLocation.Y, vehicleLocation.Z)
}
```
### In-game screenshot

<img src="https://i.imgur.com/nboAIdd.png" alt="Boat Captcha resolve" width="100%"/>

### Code execution output

```bash
Trying to find a running GTA:SA instance.
Successfully hooked into a running GTA:SA instance.         
Ped status:  InFloor                                        
User money:  24968                                          
Ped location: '-2174.303467', '-428.641479', '35.335938'    
Vehicle location: '-2171.551025', '-436.249969', '35.251785'
```