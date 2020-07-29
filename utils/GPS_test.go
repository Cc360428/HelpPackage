package utils

import (
	"fmt"
	"testing"
)

func TestEarthDistance(t *testing.T) {
	//纬度
	lon1 := 120.074234
	lat1 := 30.29092
	lon2 := 120.075977
	lat2 := 30.286866
	distance := EarthDistance(lat1, lon1, lat2, lon2)
	distance1 := LatitudeLongitudeDistance(lon1, lat1, lon2, lat2)
	fmt.Println(distance)
	fmt.Println(distance1)
}
