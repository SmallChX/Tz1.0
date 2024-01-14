package usecase

// Đây là Request của Company đặt Booth.
// Quyền xử lý: Admin. Quyền thêm, sửa: Company.
type BoothRequestUsecase interface {
	GetRequest()    // Lấy một Request
	GetAllRequest() // Lấy tất cả Request
	CreateRequest() // Thêm Request (Company Regist Booths), một Company có thể có nhiều Request.
	// Một Request có thể có nhiều hơn 1 Booth.
	// Khi đăng ký, khởi tạo Request với status là Pending.
	// Trong quá trình Pending, Admin sẽ xử lý, dựa vào policy và thanh toán.
	AcceptRequest() // Quyền xử lý: admin. Chuyển status của Request sang Accepted.
	RejectRequest() //  Quyền xử lý: admin. Chuyển status của Request sang Rejected.
	DeleteRequest() // Quyền xử lý: company. Chuyển status của Request sang Deleted.
	// Đối với Reject và Delete, không xóa mà chỉ chuyển status => Xử lý và đối chứng sau này.
}
