// package main

// import "fmt"



// func main() {
// 	x := 15
// 	a :=& x
// 	fmt.Println(a)
// 	fmt.Println(*a)
// }

// const usixteenbitmax float64 = 65535
// const kmh_mulitple float64 = 1.60934

// type car struct {
// 	gas_pedal uint16
// 	brake_pedal uint16
// 	steering_wheel int16
// 	top_speed_kmh float64
// }

// func (c car) kmh() float64 {
// 	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
// }

// func (c car) mph() float64 {
// 	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_mulitple)
// }

// func (c *car) new_top_spead(newspeed float64){
// 	c.top_speed_kmh=newspeed
// }

// func main(){
// 	a_car := car{22341, 0, 55, 140}
// 	fmt.Println(a_car.gas_pedal)
// 	fmt.Println(a_car.kmh())
// 	fmt.Println(a_car.mph())
// 	a_car.new_top_spead(500)
// 	fmt.Println(a_car.kmh())
// 	fmt.Println(a_car.mph())
// }

