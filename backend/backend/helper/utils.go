package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomPaymentID() string {
	return fmt.Sprintf("TR-%d", rand.Intn(900000)+100000)
}

func RandomStatus() string {
	r := rand.Intn(100)

	switch {
	case r < 5:
		return "pending"
	case r < 10:
		return "failed"
	case r < 40:
		return "processing"
	default:
		return "completed"
	}
}

func RandomMerchant() string {
	companies := []string{
		"PT Sumber Jaya Abadi",
		"CV Mitra Sejahtera",
		"PT Nusantara Digital",
		"PT Maju Bersama",
		"PT Global Teknologi",
		"CV Karya Mandiri",
		"PT Sentosa Makmur",
		"PT Indo Solusi",
		"PT Prima Utama",
		"PT Cipta Digital",
	}

	return companies[rand.Intn(len(companies))]
}

func RandomCreatedAt() time.Time {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Now()

	diff := end.Unix() - start.Unix()
	sec := rand.Int63n(diff)

	return time.Unix(start.Unix()+sec, 0)
}
