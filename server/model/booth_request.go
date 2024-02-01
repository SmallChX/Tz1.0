package model

import "time"

type TypeRequest string

const (
	RegistTypeRequest TypeRequest = "regist"
	ChangeTypeRequest TypeRequest = "change"
	RemoveTypeRequest TypeRequest = "remove"
)

type StatusRequest string

const (
	PedingRequest   StatusRequest = "pending"
	AcceptedRequest StatusRequest = "accepted"
	FinishedRequest StatusRequest = "finished" // Finish payment
	RejectedRequest StatusRequest = "rejected"
)

type BoothRequest struct {
	RequestID int64   `gorm:"primaryKey"`
	Booths    []Booth `gorm:"many2many:request_booths;"`
	CompanyID int64
	Status    StatusRequest
	Type      TypeRequest
	CreateAt  time.Time
	// Base on type, there are some extend information from booth request
	Reason string // request remove
	// Sử dụng một struct riêng biệt để biểu diễn quan hệ
	DestinationBooths []Booth `gorm:"many2many:des_request_booths;"`
}
