package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Kullanici struct {
	Ad      string      `json:"adi"`               //tag
	Soyad   string      `json:"-"`                 // - json ciktisinda soyad olmasini engelliyor
	Takipci []Kullanici `json:"takipci,omitempty"` // omitempty zero valuelar json da gosterilmeyecek
	// Bageni []struct {
	// 	Tarih time.Time
	// }
	Begeni []Begen
}

type Begen struct {
	Tarih time.Time
}

func main() {
	k := Kullanici{
		Ad:      "salih",
		Soyad:   "asd",
		Takipci: make([]Kullanici, 0),
	}

	arr, _ := json.Marshal(k)
	fmt.Println(string(arr))

	fmt.Println(k)
}
