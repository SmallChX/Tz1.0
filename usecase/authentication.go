package usecase

// Xử lý về phân quyền và xác thực người dùng
type AuthenticationUsecase interface {
	LoginWithAccount()
	Logout()
	RefreshToken()
}

// Đăng nhập với tài khoản: dành cho admin và company.
// Kiểm tra thông tin đăng nhập: useaccount và password trong database
// Nếu đúng, trả về role và phân hạng đối với company
func LoginWithAccount() {}

// Xóa jwt token
func Logout() {}

// Tạo lại jwt token
func RefreshToken() {}
