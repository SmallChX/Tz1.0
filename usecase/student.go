package usecase

// Student đăng nhập thông qua gmail. Đối với Student mới đăng nhập lần đầu, tạo trong database và yêu cầu cập nhật thông tin.
type StudentUsecase interface {
	GetStudent()    // Lấy thông tin Student.
	GetAllStudent() // Lấy danh sách tất cả Student.
	CreateStudent() // Tạo Student mới.
	UpdateStudent() // Cập nhật thông tin Student. (Yêu cầu đối với đăng nhập lần đầu)
	DeleteStudent() // Xóa thông tin Student.
}

// Xác thực role từ jwt xem có phải admin hoặc id có phải student hiện tại
// => lấy từ database => return
func GetStudent() {}

// Xác thực role từ jwt xem có phải admin
// => lấy từ database => return
func GetAllStudent() {}

// Xác thực role từ jwt xem có phải admin hoặc id có phải student hiện tại
// => nhận và xác thực thông tin => lưu vào database => return.
func UpdateStudent() {}

// Xác thực role từ jwt xem có phải admin
// => Xóa thông tin user, bao gồm cả mail?
func DeleteStudent() {}

// Đối với Student đăng nhập lần đầu (kiểm tra ở Login)
// Tạo ở bảng Student Info, sau đó yêu cầu Update thông tin.
func CreateStudent() {}
