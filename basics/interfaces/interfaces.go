package main

import (
	"errors"
	"fmt"
	"log"
)

// type ismi buyuk harfle basliyor ise publicdir
type Kullanici interface {
	TakipciSayisi() int
	TakipciEkle(t kullanici) error
}

// type ismi buyuk harfle basliyor ise privatedir, sadece bulundugu paketten erisilebilir
type kullanici struct {
	Ad      string
	Soyad   string
	Takipci []kullanici
}

func (k kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

func (k *kullanici) TakipciEkle(t kullanici) error {

	if k.TakipciSayisi() == 3 {
		return errors.New("maximum takipci sayisina(3) ulasildi")
	}

	if k.Takipci == nil {
		k.Takipci = make([]kullanici, 0, 3)
	}

	k.Takipci = append(k.Takipci, t)

	return nil
}

var k Kullanici

func main() {

	k := kullanici{
		Ad:    "salih",
		Soyad: "asd",
		Takipci: []kullanici{
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

	t := kullanici{
		Ad:    "ferhunde",
		Soyad: "asd",
	}

	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())

	if err := k.TakipciEkle(t); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())

	fmt.Println(buyuktur(&k, &t))
}

func buyuktur(k1, k2 Kullanici) bool {
	return k1.TakipciSayisi() > k2.TakipciSayisi()
}
