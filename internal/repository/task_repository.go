package repository

import (
    "database/sql"
    "log"
    "my-go-api/internal/models"
)

type TaskRepository struct {
    db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
    return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
    log.Println("Menjalankan query SELECT id, name, is_completed FROM tasks")
    rows, err := r.db.Query("SELECT id, name, is_completed FROM tasks")
    if err != nil {
        log.Println("Query error di GetAllTasks:", err)
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var t models.Task
        if err := rows.Scan(&t.ID, &t.Name, &t.IsCompleted); err != nil {
            log.Println("Scan error di GetAllTasks:", err)
            return nil, err
        }
        tasks = append(tasks, t)
    }
    log.Println("Tasks ditemukan:", len(tasks))
    return tasks, nil
}

func (r *TaskRepository) AddTask(task models.Task) (models.Task, error) {
    log.Println("Menjalankan query INSERT INTO tasks, name:", task.Name)
    var id int
    err := r.db.QueryRow("INSERT INTO tasks (name, is_completed) VALUES ($1, $2) RETURNING id", task.Name, task.IsCompleted).Scan(&id)
    if err != nil {
        log.Println("Query error di AddTask:", err)
        return models.Task{}, err
    }
    task.ID = id
    log.Println("Task berhasil ditambah, ID:", id)
    return task, nil
}

func (r *TaskRepository) UpdateTask(id int, task models.Task) (models.Task, error) {
    log.Println("Menjalankan query UPDATE tasks, ID:", id)
    result, err := r.db.Exec("UPDATE tasks SET name=$1, is_completed=$2 WHERE id=$3", task.Name, task.IsCompleted, id)
    if err != nil {
        log.Println("Query error di UpdateTask:", err)
        return models.Task{}, err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Println("RowsAffected error di UpdateTask:", err)
        return models.Task{}, err
    }
    if rowsAffected == 0 {
        log.Println("Task tidak ditemukan, ID:", id)
        return models.Task{}, sql.ErrNoRows
    }
    task.ID = id
    log.Println("Task berhasil diupdate, ID:", id)
    return task, nil
}

func (r *TaskRepository) DeleteTask(id int) error {
    log.Println("Menjalankan query DELETE FROM tasks, ID:", id)
    result, err := r.db.Exec("DELETE FROM tasks WHERE id=$1", id)
    if err != nil {
        log.Println("Query error di DeleteTask:", err)
        return err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Println("RowsAffected error di DeleteTask:", err)
        return err
    }
    if rowsAffected == 0 {
        log.Println("Task tidak ditemukan, ID:", id)
        return sql.ErrNoRows
    }
    log.Println("Task berhasil dihapus, ID:", id)
    return nil
}