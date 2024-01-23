package model

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
	RequestID int64
	Booths    []Booth `gorm:"many2many:request_booths;"`
	CompanyID int64
	Status    StatusRequest
	Type      TypeRequest
	// Base on type, there are some extend information from booth request
	Reason             string // request remove
	DestinationBoothID int64  // request change
}
