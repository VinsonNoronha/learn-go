package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

//struct methods value revievers pointer revievers
// value reciever methods
func (c car) kmh() float64 {
	return float64(c.gas_pedal)*(c.top_speed_kmh/usixteenbitmax) 
}
func (c car) mph() float64 {
	return float64(c.gas_pedal)*(c.top_speed_kmh/usixteenbitmax/kmh_multiple) 
}
//pointer recievers
func (c *car) new_top_speed(newspeed float64)  {
	c.top_speed_kmh = newspeed
}


//defining struct
type car struct {
	gas_pedal uint16 
	break_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func main() {
// struct
a_car := car{
	gas_pedal: 2231,
	break_pedal: 0, 
	steering_wheel: 12561, 
	top_speed_kmh: 225.0,
}

fmt.Println(a_car.gas_pedal)
fmt.Println(a_car.kmh()) // value recievers
fmt.Println(a_car.mph())
a_car.new_top_speed(500) // pointer recievers
fmt.Println(a_car.kmh())
fmt.Println(a_car.mph())

}