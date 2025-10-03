# 📘 Task-7: Package & Mini Project (Belajar Go)

Pada task ini, saya mempelajari cara menggunakan **package standar**, membuat **package sendiri**, dan membangun **mini project To-Do CLI** sederhana.

---

## 🚀 Apa yang Dipelajari

1. **Package Standar**
   - Go sudah menyediakan banyak package bawaan seperti:
     - `fmt` → untuk input/output teks.
     - `math` → fungsi matematika.
     - `time` → bekerja dengan waktu.
   - Contoh penggunaan:
     ```go
     fmt.Println("Akar kuadrat 16 =", math.Sqrt(16))
     fmt.Println("Waktu sekarang:", time.Now())
     ```

2. **Package Buatan Sendiri**
   - Membuat package `mymath` dengan fungsi `Add`.
    ```go
    package mymath

    func Add(a, b int) int {
        return a + b
    }
    ```

3. **Mini Project: To-Do CLI**
    - Package cli berisi fungsi untuk:
      - Menambah task (addTask)
      - Menampilkan task (listTasks)
      - Menghapus task (deleteTask)
      - Menjalankan CLI (RunCli)
    - Program utama (main.go) memanggil cli.RunCli() untuk menjalankan aplikasi interaktif berbasis terminal.

**📂 Struktur Project**
```swift
task-7/
├── go.mod
├── main.go       # entry point, pakai math, time, mymath, dan cli
├── mymath/
│   └── mymath.go # package custom: fungsi Add
└── cli/
    └── cli.go    # package CLI untuk To-Do list
```