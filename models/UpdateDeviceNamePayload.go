package models

type UpdateDeviceNamePayload struct {
	StudentName string `json:"student_name"`
	ClassName string `json:"class_name"`
}