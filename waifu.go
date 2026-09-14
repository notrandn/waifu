package modulewaifu

import (
	"fmt"
	"math"
	"strconv"
)

type Waifu struct {
	Name, DereType, MBTI, Address, Nationality, Hobby string
	Age, Height , Weight int // Height: cm, Weigth: kg
}

type bmi struct {
	value float64
}

func (waifu *Waifu) Introduce() {
	fmt.Println(
		"Hi! Namaku " + waifu.Name + ", usiaku " + strconv.Itoa(waifu.Age) + " tahun, hobiku adalah " +
		waifu.Hobby + " dan aku tinggal di " + waifu.Address + ".",
)
}

func (waifu *Waifu) CalculateBMI() bmi {
	HeigthMeter := float64(waifu.Height) / 100.0
	val := float64(waifu.Weight) / math.Pow(float64(HeigthMeter), float64(2.0))
	bmi := bmi{value: val}
	return bmi
}


func (bmi bmi) GetCategories() string {
	if bmi.value < 18.5 {
		return "Kurus (Kekurangan berat badan)"
	} else if bmi.value >= 18.5 && bmi.value <= 22.9 {
		return "Normal (Ideal)"
	} else if bmi.value >= 23.0 && bmi.value <= 24.9 {
		return "Kelebihan berat badan (Overweight)"
	} else if bmi.value >= 25.0 && bmi.value <= 29.9 {
		return "Obesitas Tingkat 1"
	} else {
		return "Obesitas Tingkat 2"
	}
}