package main

import "math/rand/v2"

func RandomUintInRange(min, max uint16) uint16 {
	return uint16(rand.IntN(maxHitPoints-minHitPoints+1) + minHitPoints)
}
