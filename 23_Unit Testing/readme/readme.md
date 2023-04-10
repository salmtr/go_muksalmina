Rangkuman

1.unit testing adalah suatu proses untuk menganalisa suatu software / sistem untuk menditeksi fitur saat ini dan fitur yang diinginkan oleh user.

2.unit testing Go dituliskan dalam bentuk fungsi, yang memiliki parameter yang bertipe *testing.T, dengan nama fungsi harus diawali kata Test (pastikan sudah meng-import package testing sebelumnya). Lewat parameter tersebut, kita bisa mengakses method-method untuk keperluan testing

3.method standar testing yang bisa digunakan di Go.

Method	Kegunaan
Log()	Menampilkan log
Logf()	Menampilkan log menggunakan format
Fail()	Menandakan terjadi Fail() dan proses testing fungsi tetap diteruskan
FailNow()	Menandakan terjadi Fail() dan proses testing fungsi dihentikan
Failed()	Menampilkan laporan fail
Error()	Log() diikuti dengan Fail()
Errorf()	Logf() diikuti dengan Fail()
Fatal()	Log() diikuti dengan failNow()
Fatalf()	Logf() diikuti dengan failNow()
Skip()	Log() diikuti dengan SkipNow()
Skipf()	Logf() diikuti dengan SkipNow()
SkipNow()	Menghentikan proses testing fungsi, dilanjutkan ke testing fungsi setelahnya
Skiped()	Menampilkan laporan skip
Parallel()	Menge-set bahwa eksekusi testing adalah paralle