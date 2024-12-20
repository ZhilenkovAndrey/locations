package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	landing   string
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (math.Round(c.d*1000)/1000 +
		math.Round(c.m/0.06)/1000 + math.Round(c.s/3.6)/1000)
}

func newLocation(landing string, lat, long coordinate) location {
	return location{landing, lat.decimal(), long.decimal()}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) (string, string, float64) {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return p1.landing, p2.landing, w.radius * math.Acos(s1*s2+c1*c2*clong)
	// fmt.Printf("Ditance between %s and %s is %f",
	// p1.landing, p2.landing, w.radius * math.Acos(s1*s2+c1*c2*clong))
}

func (w world) prinDistance(pointFirst, pointSecond string, distance float64) {
	fmt.Println()
	fmt.Printf("Ditance between %s and %s is %f ed.\n", pointFirst, pointSecond, distance)
	fmt.Println()
}

func main() {
	spirit := newLocation("Colambia Memorial Station",
		coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})

	oportunity := newLocation("Chalenger Memorial Station",
		coordinate{1, 556, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})

	curiosity := newLocation("Bradbury Landing",
		coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})

	inSight := newLocation("Elysium Planitia",
		coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	a := []location{spirit, oportunity, curiosity, inSight}

	for i := range a {
		fmt.Println()
		fmt.Println(a[i])
	}

	mars := world{3389.5}
	mars.prinDistance(mars.distance(spirit, oportunity))
	mars.prinDistance(mars.distance(spirit, curiosity))
	mars.prinDistance(mars.distance(spirit, inSight))
	mars.prinDistance(mars.distance(oportunity, curiosity))
	mars.prinDistance(mars.distance(oportunity, inSight))
	mars.prinDistance(mars.distance(curiosity, inSight))
}
