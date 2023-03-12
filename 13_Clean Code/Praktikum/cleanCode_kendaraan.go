package main

//struct ini menunjukkan sebuah kendaraan yang mempunyai 2 komponen
type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

//struct ini menunjukkan bahwa kendaraan tersebut berupa mobil
type mobil struct {
	kendaraan
}

//menambah kecepatan mobil sebanyak 10
func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

//menambah kecepatan mobil dengan cara menambah kecepatan baru
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

//func ini menampilkan kecepatan mobile tersebut berjalan
func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}
