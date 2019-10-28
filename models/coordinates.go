package models

type LatLng struct {
	Latitude  float64  `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}
