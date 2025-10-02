# 🚀 Learning Path Fundamental Go (1 Hari)

Belajar Golang with ChatGPT.

---

## 📌 Task 1: Dasar Sintaks & Eksekusi
- Install Go terbaru.  
- Buat project baru dengan `go mod init`.  
- Coba `go run main.go` dengan program sederhana “Hello, World!”.  
- Pelajari:
  - Variabel & konstanta  
  - Tipe data dasar (`int`, `string`, `bool`, `float`)  
  - Operator aritmatika & logika  
- **Latihan:** Buat program kalkulator sederhana (tambah, kurang, kali, bagi).  

---

## 📌 Task 2: Kontrol Alur
- Gunakan `if/else` dan `switch`.  
- Pelajari looping dengan `for`, termasuk `break` dan `continue`.  
- **Latihan:**  
  - Program untuk cek bilangan genap/ganjil.  
  - Print angka 1–100, skip kelipatan 3.  

---

## 📌 Task 3: Fungsi & Error
- Buat fungsi dengan parameter & return.  
- Belajar variadic function (`func sum(nums ...int) int`).  
- Multiple return values.  
- Error handling dengan `errors.New` dan `fmt.Errorf`.  
- **Latihan:** Fungsi `divide(a, b)` yang return hasil + error kalau `b == 0`.  

---

## 📌 Task 4: Struct & Pointer
- Buat `struct` untuk mendefinisikan tipe data sendiri.  
- Belajar pointer (`&` dan `*`).  
- Tambahkan method di struct.  
- **Latihan:** Struct `User {ID, Name}` dengan method `SayHello()`.  

---

## 📌 Task 5: Slice, Map, Array
- Belajar array (fixed size).  
- Slice (dynamic, paling sering dipakai).  
- Map (key-value seperti dictionary).  
- **Latihan:**  
  - Simpan daftar user dalam slice.  
  - Buat map dari `username -> email`.  

---

## 📌 Task 6: Interface & Goroutine
- Apa itu interface dan kenapa penting.  
- Implementasi polymorphism sederhana dengan interface.  
- Jalankan goroutine (`go func() {}`).  
- Belajar channel dasar.  
- **Latihan:** Jalankan dua goroutine: satu print angka, satu print huruf.  

---

## 📌 Task 7: Package & Mini Project
- Import package standar (`fmt`, `math`, `time`).  
- Buat package sendiri (`mymath` dengan fungsi `Add`).  
- **Mini Project (To-Do CLI):**  
  - Tambah task  
  - List task  
  - Hapus task  

---

# 🎯 Hasil Akhir
- Menguasai dasar Go (syntax, fungsi, struct, interface, goroutine).  
- Punya project kecil berbentuk CLI To-Do List.  
- Fondasi cukup kuat untuk lanjut belajar framework **Echo**.  

---