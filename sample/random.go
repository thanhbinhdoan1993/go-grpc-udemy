package sample

import (
	"math/rand"
	"time"

	"github.com/thanhbinhdoan1993/go-grpc-udemy/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7 9750M",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 2500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}