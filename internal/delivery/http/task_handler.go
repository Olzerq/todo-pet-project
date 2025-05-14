package http

import (
	"encoding/json"
	"github.com/olzerq/todo-pet-project/internal/usecase"
	"net/http"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(useCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: useCase,
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *TaskHandler) PostTask(w http.ResponseWriter, r *http.Request) {
	var taskReq struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	task, err := h.taskUseCase.CreateTask(r.Context(), taskReq.Title, taskReq.Description)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusCreated, task)
}
