package usecase

// Quyền: admin, company.
// Đối với Admin: quản lý tất cả company đăng ký tham gia JobFair và tạo account của Company
// Policy: Gửi Proposal kèm thông tin tài khoản và phương thức đăng ký cho Company.
// Đối với các công ty không nhận được proposal, sẽ có phương thức liên lạc với admin để nhận account.
type CompanyUsecase interface {
	CreateCompanyAccount() // Quyền xử lý: admin. Tạo tài khoản cho Company.
	UpdateCompany()        // Quyền xử lý: company, admin. Cập nhật thông tin Company. Sẽ được yêu cầu đối với Company đăng nhập lần đầu.
	DeleteCompany()        // Quyền xử lý: admin. Xóa Company ra khỏi web: Chỉ xóa thôn tin cơ bản, vẫn giữ account.
	GetCompany()           // Quyền: admin, company(xác thực id qua jwt). Lấy thôn tin Company.
	GetAllCompany()        // Quyền: admin. Lấy danh sách tất cả Company.
}
