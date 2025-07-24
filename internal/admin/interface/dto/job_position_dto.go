package dto


type CreateJobPositionInput struct {
   Name string `json:"name" binding:"required"`
}

type UpdateJobPositionInput struct {
   Name string `json:"name"`
}


type JobPositionResponse struct {
   ID        uint   `json:"id"`
   Name      string `json:"name"`
   CreatedAt string `json:"created_at"`
   UpdatedAt string `json:"updated_at"`
}
