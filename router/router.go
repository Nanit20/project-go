package router

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Matricula string `gorm:"not null" json:"matricula"`
	Precio    string `json:"precio"`
	Ensubasta bool   `gorm:"default:false" json:"ensubasta"`
}


func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Hola, Mundo!"))
	})

	r.Get("/tasks", getTasks(db))
	r.Post("/tasks", createTask(db))
	r.Get("/tasks/{id}", getTask(db))
	r.Put("/tasks/{id}", updateTask(db))
	r.Delete("/tasks/{id}", deleteTask(db))

	return r
}

func getTasks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks []Task
		if err := db.Find(&tasks).Error; err != nil {
			http.Error(w, "Error al obtener tareas", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func createTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Error al leer datos", http.StatusBadRequest)
			return
		}
		if err := db.Create(&task).Error; err != nil {
			http.Error(w, "Error al guardar tarea", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}


func getTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			http.Error(w, "Tarea no encontrada", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}


func updateTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			http.Error(w, "Tarea no encontrada", http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Error al leer datos", http.StatusBadRequest)
			return
		}
		if err := db.Save(&task).Error; err != nil {
			http.Error(w, "Error al actualizar tarea", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}


func deleteTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := db.Delete(&Task{}, id).Error; err != nil {
			http.Error(w, "Error al eliminar tarea", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
