# 2311102048_modul12_alpro2
latihan soal modul 12 selection short algoritma pemrograman 
1. Program bertujuan untuk mengurutkan nomor rumah di beberapa daerah berdasarkan aturan tertentu:
Bilangan ganjil diurutkan secara membesar.
Bilangan genap diurutkan secara mengecil.
Penjelasan Langkah:
Masukan Data : Program menerima jumlah daerah sebagai masukan pertama.
Untuk setiap daerah, masukan jumlah rumah dan nomor-nomor rumah di daerah tersebut.
Pemrosesan Data
Nomor rumah dipisahkan menjadi dua kelompok: ganjil dan genap.
Kelompok ganjil diurutkan membesar menggunakan selection sort.
Kelompok genap diurutkan mengecil menggunakan selection sort terbalik.
Penggabungan Hasil
Hasil pengurutan ganjil dan genap digabung, dengan ganjil berada di depan genap.
Keluaran : Program mencetak nomor rumah yang telah diurutkan untuk masing-masing daerah dalam format satu baris per daerah.

2. Program membaca data dari input pengguna untuk memproses dan mengurutkan nomor rumah di beberapa daerah. Berikut adalah langkah-langkah yang dilakukan oleh program:
1. **Input Jumlah Daerah**: Program meminta pengguna untuk memasukkan jumlah daerah yang akan diproses.
2. **Proses Setiap Daerah**:
   - Untuk setiap daerah, program meminta jumlah rumah dan memverifikasi bahwa jumlah tersebut berada dalam rentang 1 hingga 999999.
   - Program kemudian membaca nomor rumah untuk daerah tersebut.
3. **Pisahkan dan Urutkan**:
   - Nomor rumah dipisahkan menjadi dua kelompok: ganjil dan genap.
   - Nomor rumah ganjil diurutkan secara menaik menggunakan algoritma selection sort.
   - Nomor rumah genap diurutkan secara menurun menggunakan algoritma selection sort terbalik.
4. **Output**:
   - Program mencetak nomor rumah yang telah diurutkan untuk setiap daerah.
Program ini menggunakan fungsi `pisahkanGanjilGenap` untuk memisahkan nomor rumah, `urutkanSelectionSort` untuk mengurutkan nomor ganjil, dan `urutkanSelectionSortTerbalik` untuk mengurutkan nomor genap.

3. Program menghitung median dari serangkaian angka yang dimasukkan oleh pengguna. Berikut adalah deskripsi singkat dari cara kerja program:
1. **Input Pengguna**: Program meminta pengguna untuk memasukkan data berupa angka-angka yang dipisahkan oleh spasi. Input berakhir ketika pengguna memasukkan `-5313`.
2. **Pengolahan Data**: Program membaca input sebagai string dan mengonversinya menjadi slice dari integer. Setiap angka dipisahkan oleh spasi, dan program mengabaikan spasi tersebut.
3. **Penghitungan Median**: Program menggunakan algoritma bubble sort untuk mengurutkan data. Setelah data diurutkan, program menghitung median:
   - Jika jumlah data ganjil, median adalah elemen tengah.
   - Jika jumlah data genap, median adalah rata-rata dari dua elemen tengah.
4. **Output**: Program mencetak median dari data yang dimasukkan. Jika data kosong, program menampilkan pesan bahwa median tidak dapat dihitung.
5. **Penghentian**: Program berhenti memproses input ketika mendeteksi karakter `-` dalam input.
