import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { Form, Button, Table } from 'react-bootstrap';
import Swal from 'sweetalert2';

function AccountManager() {
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('');
    const [checked, setChecked] = useState(false);

    const [users, setUsers] = useState([]);

    useEffect(() => {
        async function fetchData() {
            const response = await axios.get("/api/admin/account/get-all-info")
            if (response.status === 200) {
                setUsers(response.data.data);
            } else {
                console.error("Looix");
            }
        }

        fetchData();
    }, [checked])

    async function handleSubmit() {
        
            const response = await axios.post('/api/admin/account/',{
                name: name,
                username: username,
                password: password,
                role: role,
            })
            if (response.status === 200) {
                Swal.fire("Thành công", "Đã tạo tài khoản thành công", "success");
                console.log(response)
                await setChecked(!checked);
            } else {
                Swal.fire("Lỗi", "Đã xảy ra lỗi, vui lòng thử lại", "error");
            }
        setName("");
        setUsername("");
        setPassword("");
        setRole("");
    };

    function handleResetPassword(userID, name) {
        Swal.fire({
            title: "Bạn chắc chứ?",
            text: "Bạn có muốn đổi lại mật khẩu của " + {name} + " chứ?",
            preConfirm: async () => {
                const response = await axios.post("/api/admin/account/reset-password",{
                    user_id: userID,
                })
                if (response.status === 200) {
                    Swal.fire("Thành công", "Đã reset mật khẩu thành công", "success");
                    await setChecked(!checked);
                } else {
                    Swal.fire("Lỗi", "Đã xảy ra lỗi, vui lòng thử lại", "error");
                }
            },
            confirmButtonText: "Xác nhận",
            showCancelButton: true,
        })
    }

    function handleDeleteAccount(userID, name) {
        Swal.fire({
            title: "Bạn chắc chứ?", 
            text: "Bạn có muốn đổi xóa tài khoản của " + {name} + " chứ?",
            preConfirm: async () => {
                const response = await axios.delete(`/api/admin/account/${userID}`)
                if (response.status === 200) {
                    Swal.fire("Thành công", "Đã xóa tài khoản thành công", "success");
                } else {
                    Swal.fire("Lỗi", "Đã xảy ra lỗi, vui lòng thử lại", "error");
                }
            },
            confirmButtonText: "Xác nhận",
            showCancelButton: true,
        })
        setChecked(!checked);
    }

    return (
        <div>
             <div className='container mb-3 col-5'>
            <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3" controlId="formName">
                <Form.Label>Tên</Form.Label>
                <Form.Control type="text" placeholder="Nhập tên" value={name} onChange={(e) => setName(e.target.value)} />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formUsername">
                <Form.Label>Username</Form.Label>
                <Form.Control type="text" placeholder="Nhập username" value={username} onChange={(e) => setUsername(e.target.value)} />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formPassword">
                <Form.Label>Password</Form.Label>
                <Form.Control type="password" placeholder="Nhập password" value={password} onChange={(e) => setPassword(e.target.value)} />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formRole">
                <Form.Label>Role</Form.Label>
                <Form.Select value={role} onChange={(e) => setRole(e.target.value)}>
                    <option>Chọn role...</option>
                    <option value="admin">Admin</option>
                    <option value="company">Doanh nghiệp</option>
                </Form.Select>
            </Form.Group>

            <Button variant="primary" onClick={handleSubmit}>
                Tạo Tài Khoản
            </Button>
        </Form>
        </div>
             <Table striped bordered hover>
            <thead>
                <tr>
                    <th>Tên</th>
                    <th>Username</th>
                    <th>Role</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {users.map(user => (
                    <tr key={user.id}>
                        <td>{user.name}</td>
                        <td>{user.username}</td>
                        <td>{user.role}</td>
                        <td>
                            <Button variant="warning" onClick={() => handleResetPassword(user.user_id, user.name)}>Reset Password</Button>
                            <Button variant='danger' onClick={() => handleDeleteAccount(user.user_id, user.name)}>Xóa tài khoản</Button>
                        </td>
                    </tr>
                ))}
            </tbody>
        </Table>
        </div>
       
    );
}

export default AccountManager;
