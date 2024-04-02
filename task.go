package main

import (
	"database/sql"
)

type Tasks struct {
	ID          int    `db:"idTask"` // Указываем, что ID соответствует колонке idTask в БД
	Title       string `db:"title"`
	Description string `db:"description"`
	CategoryID  int    `db:"category_id"`
	Completed   bool   `db:"completed"`
}


// Создаем новую задачу
func CreateTask(db *sql.DB, title string, description string, categoryID int) error {
	statement, err := db.Prepare("INSERT INTO tasks (title, description, category_id, completed) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(title, description, categoryID, 0)
	if err != nil {
		return err
	}

	return nil
}

// Обновляем задачу
func UpdateTask(db *sql.DB, id int, title string, description string, categoryID int, completed bool) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, category_id = ?, completed = ? WHERE idTask = ?", title, description, categoryID, completed, id)
	return err
}

// Удаляем задачу
func DeleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE idTask = ?", id)
	return err
}

// Выводим весь список задач
func GetAllTasks(db *sql.DB) ([]Tasks, error) {
	rows, err := db.Query("SELECT idTask, title, description, category_id, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Tasks

	for rows.Next() {
		var t Tasks
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CategoryID, &t.Completed)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
