package model

type CompanyContributionLevel string

const (
	PlatinumContributor CompanyContributionLevel = "platinum_contributor"
	GoldContributor     CompanyContributionLevel = "gold_contributor"
	SilverContributor   CompanyContributionLevel = "silver_contributor"
)

type CompanyInformation struct {
	UserID      int64
	IdCompany   int64
	CompanyName string

	// Contact Represent = HR
	RepresentName        string
	RepresentPhoneNumber string
	RepresentMail        string

	// Level in Contribution
	ContributionLevel CompanyContributionLevel
	Booths            []Booth        `gorm:"foreignKey:CompanyID"` // Booth Own
	BoothRequests     []BoothRequest `gorm:"foreignKey:RequestID"`
}
