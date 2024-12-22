package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description"` // Описание
	Date        string `json:"date"`        // Дата и время
	Time        string `json:"time"`        // Дата и время
	Status      bool   `json:"status"`      // Статус (полностью/не полностью)
}
