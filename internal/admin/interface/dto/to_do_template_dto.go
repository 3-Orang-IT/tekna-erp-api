package dto

type CreateToDoTemplateInput struct {
	JobPositionID uint   `json:"job_position_id" binding:"required"`
	Activity      string `json:"activity" binding:"required"`
	Priority      int    `json:"priority" binding:"required"`
	OrderNumber   int    `json:"order_number" binding:"required"`
}

type UpdateToDoTemplateInput struct {
	JobPositionID uint   `json:"job_position_id"`
	Activity      string `json:"activity"`
	Priority      int    `json:"priority"`
	OrderNumber   int    `json:"order_number"`
}

type ToDoTemplateResponse struct {
	ID            uint   `json:"id"`
	JobPositionID uint   `json:"job_position_id"`
	JobPosition   string `json:"job_position"`
	Activity      string `json:"activity"`
	Priority      int    `json:"priority"`
	OrderNumber   int    `json:"order_number"`
}

type JobPositionOption struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GroupedToDoTemplateResponse struct {
	JobPositionID uint                   `json:"job_position_id"`
	JobPosition   string                 `json:"job_position"`
	Activities    []ToDoTemplateResponse `json:"activities"`
}
