package ipdb

type IpSegment struct {
	Start uint32
	End   uint32
}

type IpSegments []IpSegment

type IpDB struct {
	IpSegments IpSegments
	count      uint32
}
