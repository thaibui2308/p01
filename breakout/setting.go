package breakout

import (
	"image/color"
	"math"
)

const (
	// Screen layout's settings
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 240

	// Player's settings
	PLAYER_HEIGHT       = SCREEN_HEIGHT / 24
	PLAYER_WIDTH        = SCREEN_WIDTH / 12
	PLAYER_X            = (SCREEN_WIDTH - PLAYER_WIDTH) / 2
	PLAYER_Y            = SCREEN_HEIGHT * .8
	PLAYER_MOVING_SPEED = 3

	// Bullet's settings
	BULLET_HEIGHT        = PLAYER_HEIGHT / 2
	BULLET_WIDTH         = BULLET_HEIGHT
	BULLET_X             = PLAYER_X + (PLAYER_WIDTH-BULLET_WIDTH)/2
	BULLET_Y             = PLAYER_Y - BULLET_HEIGHT
	RIGHT_UPWARD_ANGLE   = -45
	RIGHT_DOWNWARD_ANGLE = 20
	LEFT_UPWARD_ANGLE    = 20
	LEFT_DOWNWARD_ANGLE  = -20
	BULLET_SPEED         = 2

	// Targer and TargetSprite' settings
	NUM_ROWS                   = 4
	NUM_TARGETS                = 7
	SCREEN_HORIZONTAL_PADDING  = SCREEN_WIDTH / 10
	SCREEN_VERTICAL_PADDING    = SCREEN_HEIGHT / 5
	TARGET_WIDTH               = (SCREEN_WIDTH - 2*SCREEN_HORIZONTAL_PADDING) / (1.25*NUM_TARGETS - .25)
	TARGET_HEIGHT              = TARGET_WIDTH / 3
	SPRITE_HORIZONTAL_DISTANCE = TARGET_WIDTH / 4
	SPRITE_VERTICAL_DISTANCE   = (.6 * SCREEN_HEIGHT) * .15

	// Particle settings
	PARTICLE_WIDTH  = BULLET_HEIGHT / 2.5
	PARTICLE_HEIGHT = BULLET_WIDTH / 2.5
)

var (
	INITIAL_BULLET_DIRECTION = &Coordinate{
		X: BULLET_SPEED / math.Sqrt(math.Tan(RIGHT_UPWARD_ANGLE)*math.Tan(RIGHT_UPWARD_ANGLE)+1),
		Y: math.Tan(RIGHT_UPWARD_ANGLE) * (BULLET_SPEED / math.Sqrt(math.Tan(RIGHT_UPWARD_ANGLE)*math.Tan(RIGHT_UPWARD_ANGLE)+1)),
	}
)

var (
	PALETTE = []color.RGBA{
		{255, 0, 0, 0xff},
		{255, 127, 0, 0xff},
		{255, 255, 0, 0xff},
		{0, 255, 0, 0xff},
	}
)
