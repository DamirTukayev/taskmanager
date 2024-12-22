package models

type Entry struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description"` // Описание
	Date        string `json:"date"`        // Дата
	Time        string `json:"time"`        // Время
	Status      bool   `json:"status"`      // Статус (полностью/не полностью)
}
