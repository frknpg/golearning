package main

import (
	"errors"
	"fmt"
	"log"
)

type Kullanici struct {
	Ad      string
	Soyad   string
	Takipci []Kullanici
}

func (k Kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

func (k *Kullanici) TakipciEkle(t Kullanici) error {

	if k.TakipciSayisi() == 2 {
		return errors.New("maximum takipci sayisina(2) ulasildi")
	}

	if k.Takipci == nil {
		k.Takipci = make([]Kullanici, 0, 2)
	}

	k.Takipci = append(k.Takipci, t)

	return nil
}

func main() {

	k := Kullanici{
		Ad:    "salih",
		Soyad: "asd",
		Takipci: []Kullanici{
			{
				Ad:    "takipci",
				Soyad: "bir",
			},
			{
				Ad:    "takipci",
				Soyad: "iki",
			},
		},
	}

	t := Kullanici{
		Ad:    "ferhunde",
		Soyad: "asd",
	}

	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())

	if err := k.TakipciEkle(t); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())
}
