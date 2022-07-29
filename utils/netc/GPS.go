// GPS 定位计算距离
package netc

import "math"

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6378137) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// 转化为弧度(rad)
func rad(d float64) float64 {
	// 3.14159265358979323846264338327950288419716939937510582097494459
	r := d * math.Pi / 180.0
	return r
}

func LatitudeLongitudeDistance(lon1, lat1, lon2, lat2 float64) (distance float64) {
	// 赤道半径(单位m)
	const EarthRadius = 6378137
	// 转换为弧度
	radLat1 := rad(lat1)
	radLon1 := rad(lon1)
	radLat2 := rad(lat2)
	radLon2 := rad(lon2)
	if radLat1 < 0 {
		radLat1 = math.Pi/2 + math.Abs(radLat1)
	}
	if radLat1 > 0 {
		radLat1 = math.Pi/2 - math.Abs(radLat1)
	}
	if radLon1 < 0 {
		radLon1 = math.Pi*2 - math.Abs(radLon1)
	}
	if radLat2 < 0 {
		radLat2 = math.Pi/2 + math.Abs(radLat2)
	}
	if radLat2 > 0 {
		radLat2 = math.Pi/2 - math.Abs(radLat2)
	}
	if radLon2 < 0 {
		radLon2 = math.Pi*2 - math.Abs(radLon2)
	}
	// 		地球半径				余弦值				正弦值
	x1 := EarthRadius * math.Cos(radLon1) * math.Sin(radLat1)
	y1 := EarthRadius * math.Sin(radLon1) * math.Sin(radLat1)
	z1 := EarthRadius * math.Cos(radLat1)

	x2 := EarthRadius * math.Cos(radLon2) * math.Sin(radLat2)
	y2 := EarthRadius * math.Sin(radLon2) * math.Sin(radLat2)
	z2 := EarthRadius * math.Cos(radLat2)
	// 		计算平方根
	d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	//  		弧度的x的反余弦值
	theta := math.Acos((EarthRadius*EarthRadius + EarthRadius*EarthRadius - d*d) / (2 * EarthRadius * EarthRadius))
	distance = theta * EarthRadius
	return
}
