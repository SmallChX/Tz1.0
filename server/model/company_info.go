package model

type CompanyContributionLevel string

const (
	PlatinumContributor CompanyContributionLevel = "platinum_contributor"
	GoldContributor     CompanyContributionLevel = "gold_contributor"
	SilverContributor   CompanyContributionLevel = "silver_contributor"
)

type CompanyInformation struct {
	ID   int64 `gorm:"unique"`
	CompanyName string

	UserAccountID int64
	UserAccount   UserAccount `gorm:"foreignKey:UserAccountID"`

	// Contact Represent = HR
	RepresentName        string
	RepresentPhoneNumber string
	RepresentMail        string

	// Level in Contribution
	ContributionLevel CompanyContributionLevel
	Booths            []Booth        `gorm:"foreignKey:CompanyID;references:ID"` // Booth Own
	BoothRequests     []BoothRequest `gorm:"foreignKey:CompanyID;references:ID"`
}
