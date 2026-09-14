package modulewaifu

import (
	"fmt"
	"strconv"
)

type Waifu struct {
	Name, DereType, MBTI, Address, Nationality, Hobby string
	Age, Height, Weight int
}

func (waifu *Waifu) Introduce() {
	fmt.Println(
		"Hi! Namaku " + waifu.Name + ", usiaku " + strconv.Itoa(waifu.Age) + " tahun, hobiku adalah " +
		waifu.Hobby + " dan aku tinggal di " + waifu.Address + ".",
)
}