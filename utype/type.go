package utype

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Uint | Int | Float
}
