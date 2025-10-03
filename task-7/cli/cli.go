package cli

import "fmt"

var tasks []string

func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Task ditambahkan: ", task)
}

func listTasks() {
	fmt.Println("Daftar Task: ")
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func deleteTask(index int) {
	if index < 1 || index > len(tasks) {
		fmt.Println("Index tidak valid")
		return
	}

	removed := tasks[index-1]
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task dihapus: ", removed)
}

func RunCli() {
	for {
		fmt.Println("\n=== To-Do CLI===")
		fmt.Println("1. Tambah Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Hapus Task")
		fmt.Println("4. Keluar")
		fmt.Println("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				var task string
				fmt.Print("Masukkan task: ")
				fmt.Scan(&task)
				addTask(task)
			case 2:
				listTasks()
			case 3:
				var index int
				fmt.Print("Masukkan nomor task: ")
				fmt.Scan(&index)
				deleteTask(index)
			case 4:
				fmt.Println("Keluar...")
				return
			default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}