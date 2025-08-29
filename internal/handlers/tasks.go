package handlers

import (
    "database/sql"
    "encoding/json"
    "log"
    "my-go-api/internal/models"
    "my-go-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

var taskService *services.TaskService

func SetTaskService(s *services.TaskService) {
    taskService = s
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    tasks := taskService.GetAllTasks()
    if err := json.NewEncoder(w).Encode(tasks); err != nil {
        log.Println("Error encoding JSON di GetTasks:", err)
        http.Error(w, "Gagal encode JSON", http.StatusInternalServerError)
        return
    }
    log.Println("Berhasil get tasks, jumlah:", len(tasks))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newTask models.Task
    if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
        log.Println("Error decoding JSON di CreateTask:", err)
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    if newTask.Name == "" {
        log.Println("Name kosong di CreateTask")
        http.Error(w, "Name wajib diisi", http.StatusBadRequest)
        return
    }
    task := taskService.AddTask(newTask)
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(task); err != nil {
        log.Println("Error encoding JSON di CreateTask:", err)
        http.Error(w, "Gagal encode JSON", http.StatusInternalServerError)
        return
    }
    log.Println("Task berhasil dibuat, ID:", task.ID)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Println("ID tidak valid di UpdateTask:", vars["id"])
        http.Error(w, "ID tidak valid", http.StatusBadRequest)
        return
    }
    var updatedTask models.Task
    if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
        log.Println("Error decoding JSON di UpdateTask:", err)
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    if updatedTask.Name == "" {
        log.Println("Name kosong di UpdateTask")
        http.Error(w, "Name wajib diisi", http.StatusBadRequest)
        return
    }
    task, err := taskService.UpdateTask(id, updatedTask)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("Task tidak ditemukan di UpdateTask, ID:", id)
            http.Error(w, "Task tidak ditemukan", http.StatusNotFound)
            return
        }
        log.Println("Error di UpdateTask:", err)
        http.Error(w, "Gagal update task: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(task); err != nil {
        log.Println("Error encoding JSON di UpdateTask:", err)
        http.Error(w, "Gagal encode JSON", http.StatusInternalServerError)
        return
    }
    log.Println("Task berhasil diupdate, ID:", id)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Println("ID tidak valid di DeleteTask:", vars["id"])
        http.Error(w, "ID tidak valid", http.StatusBadRequest)
        return
    }
    if err := taskService.DeleteTask(id); err != nil {
        if err == sql.ErrNoRows {
            log.Println("Task tidak ditemukan di DeleteTask, ID:", id)
            http.Error(w, "Task tidak ditemukan", http.StatusNotFound)
            return
        }
        log.Println("Error di DeleteTask:", err)
        http.Error(w, "Gagal hapus task: "+err.Error(), http.StatusInternalServerError)
        return
    }
    log.Println("Task berhasil dihapus, ID:", id)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task berhasil dihapus"})
}