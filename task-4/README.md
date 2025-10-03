# 📘 Task-4: Struct & Pointer (Belajar Go)

Pada task ini, saya mempelajari cara membuat **struct** di Go, bermain dengan **pointer**, dan menambahkan **method** pada struct.  

---

## 🚀 Apa yang Dipelajari

1. **Struct**
    - Struct digunakan untuk mendefinisikan tipe data baru dengan beberapa field.
    ```go
    type User struct {
        ID   int
        Name string
    }
    ```

     - Inisialisasi struct:
    ```go
    u1 := user.User{ID: 1, Name: "Linux"}
    u2 := user.User{2, "Budi"}
    ```

2. **Method di Struct**
    - Method ditambahkan dengan receiver pada struct.
    ```go
    func (u User) SayHello() {
        fmt.Println("Hello,", u.Name)
    }
    ```

3. **Pointer**
    - ```&``` untuk mengambil alamat memori.
    - ```*``` untuk mengambil nilai dari alamat.
    ```go
    x := 10
    p := &x
    fmt.Println(*p) // akses nilai
    *p = 20         // ubah nilai melalui pointer
    fmt.Println(x)  // x sekarang 20
    ```

---

**📂 Struktur Project**

```swift
task-4/
├── go.mod
├── main.go      # program utama
└── user/
    └── user.go  # definisi struct User dan method SayHello
```