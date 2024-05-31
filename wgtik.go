package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var banyakData int

type pelaku struct {
	plat, nama, tipe                  string
	rute                              [10]string
	tanggal, pelanggaran, panjangRute int
}

func ExtractNumbersFromPlate(plate string) (int, error) {
	var numbers string
	for _, char := range plate {
		if unicode.IsDigit(char) {
			numbers += string(char)
		}
	}

	if numbers == "" {
		return 0, fmt.Errorf("no numbers found in the plate")
	}

	result, err := strconv.Atoi(numbers)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func cekPlat(plat string) bool {
	if strings.ToLower(plat) == "mobil" {
		return true
	} else {
		return false
	}
}

func CekPelanggaran(orang *[10]pelaku) {
	var digitTerakhir int
	for j := 1; j <= banyakData; j++ {
		if cekPlat(orang[j].tipe) {
			nomorPlat, err := ExtractNumbersFromPlate(orang[j].plat)
			if err != nil {

			} else {
				digitTerakhir = nomorPlat % 100
				if orang[j].tanggal%2 == 0 {
					if digitTerakhir%2 == 0 {

					} else {
						for i := 1; i <= orang[j].panjangRute; i++ {
							if strings.ToLower(orang[j].rute[i]) == "fatmawati" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "gajahMada" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "hayamwuruk" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "sisingamangaraja" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "panglimapolim" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "tomangraya" {
								orang[j].pelanggaran++
							}
						}
					}
				} else {
					if digitTerakhir%2 != 0 {

					} else {
						for i := 1; i <= orang[j].panjangRute; i++ {
							if strings.ToLower(orang[j].rute[i]) == "fatmawati" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "gajahMada" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "hayamwuruk" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "sisingamangaraja" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "panglimapolim" {
								orang[j].pelanggaran++
							} else if strings.ToLower(orang[j].rute[i]) == "tomangraya" {
								orang[j].pelanggaran++
							}
						}
					}
				}
			}
		} else {

		}

	}

}

func showPelanggar(orang [10]pelaku) {
	for i := 1; i <= banyakData; i++ {
		if orang[i].pelanggaran != 0 {
			fmt.Println("name :", orang[i].nama, ", tilang : ", orang[i].pelanggaran)
		} else {

		}
	}
}

func main() {
	var orang [10]pelaku

	fmt.Print("Masukkan banyak data : ")
	fmt.Scan(&banyakData)
	for x := 1; x <= banyakData; x++ {
		fmt.Print("Masukkan Tanggal : ")
		fmt.Scan(&orang[x].tanggal)
		if orang[x].tanggal < 1 || orang[x].tanggal > 31 {
			fmt.Println("tanggal tidak valid")
		} else {
			fmt.Print("Nama : ")
			fmt.Scan(&orang[x].nama)
			fmt.Print("Plat : ")
			fmt.Scan(&orang[x].plat)
			fmt.Print("tipe : ")
			fmt.Scan(&orang[x].tipe)
			fmt.Print("Masukkan panjang rute : ")
			fmt.Scan(&orang[x].panjangRute)
			orang[x].pelanggaran = 0

			for i := 1; i <= orang[x].panjangRute; i++ {
				fmt.Print("Masukkan rute ke-", i, ": ")
				fmt.Scan(&orang[x].rute[i])
			}
		}
		fmt.Println()
	}

	CekPelanggaran(&orang)

	showPelanggar(orang)

}
