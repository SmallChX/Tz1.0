import { Navigate } from "react-router-dom";
import { useLocation } from "react-router-dom";
import Swal from 'sweetalert2';

const RequireAuth = ({ children, userroles }) => {
    let currentUserRole;
    let firstLogin;

    // Lấy giá trị currentUserRole từ localStorage
    if (localStorage.getItem("userRole")) {
        currentUserRole = localStorage.getItem("userRole");
    }

    // Lấy giá trị firstLogin từ localStorage
    if (localStorage.getItem("firstLogin")) {
        firstLogin = localStorage.getItem("firstLogin") === 'true';
    }

    const location = useLocation();

    // Kiểm tra nếu người dùng chưa đăng nhập
    if (!currentUserRole) {
        return <Navigate to="/login" state={{ from: location }} />
    }

    // Kiểm tra nếu là lần đầu đăng nhập
    if (firstLogin && location.pathname !== "/profile") {
        Swal.fire('Vui lòng điền thông tin', "", 'info');
        return <Navigate to="/profile" state={{ from: location }} />;
    }

    // Kiểm tra quyền truy cập dựa trên vai trò người dùng
    if (userroles) {
        if (userroles.includes(currentUserRole)) {
            return children;
        } else {
            Swal.fire('Access Denied !', "", 'warning');
            return <Navigate to="/home" />;
        }
    } else {
        return children;
    }
};

export default RequireAuth;
