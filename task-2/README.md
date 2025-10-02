# 📘 Task-2: Kontrol Alur (Belajar Go)

Pada task ini, saya mempelajari **kontrol alur** dalam bahasa Go: `if/else`, `switch`, dan variasi `for loop` (karena Go tidak punya `while` atau `do while`, semuanya disimulasikan dengan `for`).

---

## 🚀 Apa yang Dipelajari

1. **If / Else**
   - Mengecek kondisi sederhana.
   - Nested dengan `else if`.
   ```go
   umur := 21
   if umur >= 18 {
       fmt.Println("Dewasa")
   } else {
       fmt.Println("Anak-anak")
   }
   ```

2. **If-Else If-Else**
    ```go
    var a, b float64 = 2, 5
    sum := b / a

    if sum == 5 {
        fmt.Println("penjumlahan = 5", sum)
    } else if sum < 5 {
        fmt.Println("penjumlahan < 5", sum)
    } else {
        fmt.Println("penjumlahan > 5", sum)
    }
    ```

3. **Switch**
    ```go
    hari := 3
    switch hari {
        case 1: fmt.Println("Senin")
        case 2: fmt.Println("Selasa")
        case 3: fmt.Println("Rabu")
        ...
        default: fmt.Println("Nomor hari tidak valid")
    }
    ```

4. **For Loop**
    ```go
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
    ```
    - Dengan break dan continue:
    ```go
    for i := 1; i < 10; i++ {
        if i == 5 {
            break
        }
        if i % 2 == 0 {
            continue
        }
        fmt.Println(i)
    }
    ```

5. **While Style**
    - Go tidak punya ```while```, jadi gunakan ```for``` dengan kondisi.
    ```go
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
    ```

6. **Do-While Style**
    - Disimulasikan dengan ```for {}``` + ```break```.
    ```go
    j := 0
    for {
        fmt.Println(j)
        j++
        if j >= 5 {
            break
        }
    }
    ```

7. **Infinite Loop**
    ```go
    for {
        fmt.Println("jalan terus...")
        break
    }
    ```
