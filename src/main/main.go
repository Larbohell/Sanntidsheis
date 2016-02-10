package main
import "driver"
//import "fmt"

func main() {
	driver.Io_init()
	
	for {
		if (driver.Elevator_is_button_pushed(driver.BUTTON_OUTSIDE_UP, 2)){
			driver.Elevator_set_motor_direction(driver.MOTOR_DIRECTION_DOWN)
		}
		if (driver.Elevator_is_button_pushed(driver.BUTTON_OUTSIDE_DOWN, 2)){
			driver.Elevator_set_motor_direction(driver.MOTOR_DIRECTION_UP)
		}
		
		if (driver.Elevator_get_stop_signal()){
			driver.Elevator_set_motor_direction(driver.MOTOR_DIRECTION_STOP)
		}

		for floor := 0; floor < 4; floor++ {
			for button := 0; button < 3; button++ {
			if (driver.Elevator_is_button_pushed(driver.Button(button), floor)){
				driver.Elevator_set_button_lamp(driver.Button(button), floor, 1)
				}
			}
		}
		if (driver.Elevator_get_floor_sensor_signal() != -1){
			driver.Elevator_set_floor_indicator(driver.Elevator_get_floor_sensor_signal())
		}

		if (driver.Elevator_is_button_pushed(driver.BUTTON_INSIDE_COMMAND, 3)){
			driver.Elevator_set_door_open_lamp(1)
		}
	}
}