package main

import "fmt"

type Kullanici struct {
	Ad      string
	Soyad   string
	Takipci []Kullanici
}

//normal receiver degerleri okuma icin
func (k Kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

//pointer receiver degisiklikler icin
func (k *Kullanici) TakipciEkle(t Kullanici) {
	if k.Takipci == nil {
		k.Takipci = []Kullanici{}
	}

	k.Takipci = append(k.Takipci, t)
}

func main() {

	k := Kullanici{
		Ad:    "salih",
		Soyad: "asd",
	}

	t := Kullanici{
		Ad:    "ferhunde",
		Soyad: "asd",
	}

	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())
	k.TakipciEkle(t)
	fmt.Println(k)
	fmt.Printf("Takipci sayisi: %d\n", k.TakipciSayisi())
}
