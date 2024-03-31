package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	db := InitDB("tasks.db")
	defer db.Close()

	CreateTable(db)

	for {
		fmt.Println("Какие действия вы хотите произвести с задачами?")
		fmt.Printf("1 = Добавить новую задачу\n2 = Показать все задачи\n3 = Обновить задачу\n4 = Удалить задачу\n5 = выход\n")

		choice := readLine()
		switch choice {
		case "1":
			fmt.Println("Ввидите загаловок задачи:")
			title := readLine()

			fmt.Println("Ввидите описание задачи:")
			description := readLine()

			var category_id int
			fmt.Println("Ввидите категорию задачи:")
			fmt.Scan(&category_id)

			err := CreateTask(db, title, description, category_id)
			if err != nil {
				log.Fatalf("Ошибка добавления задачи: %v", err)
			}

		case "2":
			tasks, err := GetAllTasks(db)
			if err != nil {
				log.Fatal("Ошибка получения списка задач", err)
			}

			for _, task := range tasks {
				fmt.Printf("ID: %d, Title: %s, Description: %s, CategoryId: %d, Completed: %t\n", task.ID, task.Title, task.Description, task.CategoryID, task.Completed)
			}

		case "3":
			var id, category_id int
			fmt.Println("Ввидите id задачи, которую хотите изменить:")
			fmt.Scan(&id)

			fmt.Println("Ввидите заголовок задачи:")
			title := readLine()

			fmt.Println("Ввидите описание задачи:")
			description := readLine()

			fmt.Println("Ввидите категорию задачи")
			fmt.Scan(&category_id)

			fmt.Println("Вы хотите отметить ее как выполненую ?: (да/нет)")
			choice := readLine()
			var completed bool

			switch choice {
			case "да":
				completed = true
			case "нет":
				completed = false
			default:
				fmt.Println("Вы ввели неверное значение")
			}

			if err := UpdateTask(db, id, title, description, category_id, completed); err != nil {
				log.Fatalf("Ошибка обновления задачи: %v", err)
			}

		case "4":
			fmt.Println("Ввидите id задачи, которую вы хотите удалить:")
			var id int
			fmt.Scan(&id)

			if err := DeleteTask(db, id); err != nil {
				log.Fatalf("Ошибка удаления задачи: %v", err)
			}

		case "5":
			fmt.Println("Выход из списка задач")
			return
		default:
			fmt.Println("Не верный выбор, попробуйте снова.")
		}
	}
}
