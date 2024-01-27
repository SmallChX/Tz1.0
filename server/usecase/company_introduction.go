package usecase

// Gửi kèm proposal. Hình thức mới để giới thiệu Company, không thông qua Facebook nữa.
type CompanyIntroductionUsecase interface {
	GetAllCompanyIntroduction()  // Lấy danh sách tất cả Company Introduction.
	GetCompanyIntroduction()     // Lấy thông tin Company Introduction.
	CreateCompanyIntroduction()  // Quyền: admin, company. Tạo Company Introduction.
	DeleteCompanyIntrocduction() // Quyền: admin, company(xác thực id với jwt). Xóa Company Introduction.
	UpdateCompanyIntroduction()  // Quyền: conmpany. Cập nhật thông tin Company Introduction.
}


