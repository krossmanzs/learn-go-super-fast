# 📘 Task-3: Fungsi & Error (Belajar Go)

Pada task ini, saya mempelajari cara membuat fungsi di Go, penggunaan variadic function, multiple return values, dan cara melakukan error handling sesuai gaya idiomatis Go.

---

## 🚀 Apa yang Dipelajari

1. **Fungsi dengan Parameter & Return**
    - Fungsi harus didefinisikan tipe data parameternya.
    ```go
    func add(a int, b int) int {
        return a + b
    }
    ```

2. **Variadic Function**
    - Menerima jumlah argumen tak terbatas.
    - Diproses dengan for ... range.
    ```go
    func sum(nums ...int) int {
        total := 0
        for _, n := range nums {
            total += n
        }
        return total
    }
    ```

3. **Multiple Return Values**
    - Go mendukung return lebih dari satu nilai.
    ```go
    func minMax(a, b int) (int, int) {
        if a > b {
            return b, a
        }
        return a, b
    }
    ```

4. **Error Handling**
    - Gunakan ```errors.New()``` atau ```fmt.Errorf()```.
    - Pesan error tidak diawali huruf kapital (kecuali nama khusus).
    ```go
    func divide(a, b int) (int, error) {
        if b == 0 {
            return 0, errors.New("pembagi tidak boleh nol")
        }
        return a / b, nil
    }
    ```