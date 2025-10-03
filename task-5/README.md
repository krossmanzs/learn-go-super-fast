# 📘 Task-5: Slice, Map, Array (Belajar Go)

Pada task ini, saya mempelajari **array**, **slice**, dan **map** di Go.  
Array digunakan untuk data berukuran tetap, slice untuk data dinamis (yang paling sering dipakai), dan map untuk struktur key-value seperti dictionary.

---

## 🚀 Apa yang Dipelajari

1. **Array (Fixed Size)**
    - Array memiliki ukuran tetap.
    ```go
    arr := [3]int{1, 2, 3}
    fmt.Println("Array:", arr)
    fmt.Println("Index ke-1:", arr[1])
    ```

2. **Slice (Dynamic)**
    - Slice fleksibel, bisa ditambah elemen dengan ```append```.
    ```go
    nums := []int{1, 2, 3}
    nums = append(nums, 4, 5)
    fmt.Println("Slice:", nums)
    ```

3. **Struct + Slice**
    - Struct digunakan untuk membuat tipe data baru.
    - Slice bisa menyimpan kumpulan struct.
    ```go
    type User struct {
        Username string
        Email    string
    }

    users := []User{
        {"Linux", "linux@mail.com"},
        {"Budi", "budi@mail.com"},
    }

    for _, u := range users {
        fmt.Println("User:", u.Username, "Email:", u.Email)
    }
    ```