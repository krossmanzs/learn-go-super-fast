# 📘 Task-1: Dasar Sintaks & Eksekusi (Belajar Go)

Pada task ini, saya mempelajari dasar-dasar bahasa Go: struktur program, variabel, konstanta, operator, tipe data, dan pemisahan fungsi ke file/package lain.

---

## 🚀 Apa yang Dipelajari

1. **Struktur Dasar File Go**
   - Semua file Go dimulai dengan deklarasi `package`.
   - Menggunakan `import` untuk library atau package buatan sendiri.
   - Fungsi `main()` sebagai entry point program.

2. **Variabel & Konstanta**
   - Deklarasi klasik:
     ```go
     var name string = "Cornelius Linux"
     var age = 21
     var x, y int = 10, 20
     ```
   - Shorthand (hanya di dalam fungsi):
     ```go
     city := "Jakarta"
     ```
   - Konstanta:
     ```go
     const Pi = 3.14
     const Greeting = "Selamat datang!"
     ```

3. **Operator**
   - Aritmatika: `+ - * / %`
   - Logika dan pembanding: `== != > <`

4. **Output**
   - `fmt.Println()` untuk print biasa.
   - `fmt.Printf()` untuk format string.
   - Multiline string dengan backtick (\`...\`).

5. **Tipe Data Dasar**
   - Integer (`int`), Float (`float64`), Boolean (`bool`), String (`string`).
   - Operasi string: `len()`, akses index, substring.

6. **Struktur Data**
   - Array, Slice, Map, Struct, Pointer.

---

## 📂 Struktur Project

```swift
task-1/
├── go.mod
├── main.go # Program utama
├── operator.go # Package util: fungsi PrintSum
└── test.go # Package lain: array, slice, map, struct, pointer
```


---
