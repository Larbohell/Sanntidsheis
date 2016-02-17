package driver

const N_FLOORS = 4;

type Button int

const (
	BUTTON_OUTSIDE_UP Button = 0
	BUTTON_OUTSIDE_DOWN Button = 1
	BUTTON_INSIDE_COMMAND Button = 2
)

type Lamp int

const (
	LAMP_OUTSIDE_UP__BUTTON = 0
	LAMP_OUTSIDE_DOWN__BUTTON = 1
	LAMP_FLOOR_INDICATOR = 2
	LAMP_INSIDE_COMMAND_BUTTON = 3
	LAMP_INSIDE_STOP_BUTTON = 4
	LAMP_INSIDE_DOOR_INDICATOR = 5
)


type Motor_direction int

const (
	MOTOR_DIRECTION_DOWN Motor_direction = -1
	MOTOR_DIRECTION_UP Motor_direction = 1
	MOTOR_DIRECTION_STOP Motor_direction = 0
)

const lamp_channel_matrix = make([][]int)

func elevator_init() bool {
	 if (Io_init()){
	 	// Shit failed, handle this
	 	return false
	 }

	 return true
}

func lamp_on(lamp Lamp, floor int) {
	Io_set_bit()
}

func lamp_off(lamp Lamp, floor int) {

}