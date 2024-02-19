import React, { useEffect, useState} from 'react'
import { useNavigate } from 'react-router-dom';
import axios from 'axios'
import { Container, Row, Col, Form, Button } from 'react-bootstrap';
import Swal from 'sweetalert2';

function Login() {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [userRole, setUserRole] = useState('');

    useEffect(() => {
        if (localStorage.getItem('userRole')) {
        }
    });
    
    const handleLogin = async () => {
        try {
            const response = await axios.post('/api/auth/login', {
                username, 
                password,
            });
            
            if (response.status === 200) {
                const data = response.data.data;
                setUserRole(data);
                
                localStorage.setItem('userRole', data.user_role);
                localStorage.setItem('firstLogin', data.first_login);
                
                navigate("/");
            } else {
                console.error('Đăng nhập thất bại');
            }
        } catch (error) {
            console.error('Đã xảy ra lỗi', error);
        }
        
    };

    const handleLoginWithGoogle = () => {
        Swal.fire("Tính năng trong tương lại", "Hiện tại trang web chưa phổ biến đến sinh viên, bạn vui lòng quay lại sau", "info");
    };

    return (
        <Container style={{marginTop: "200px"}}>
        <Row>
            <Col>
                <h1>Với doanh nghiệp</h1>
                <Form>
                    <Form.Group controlId="formBasicUsername">
                        <Form.Label>Username:</Form.Label>
                        <Form.Control type="text" value={username} onChange={(e) => setUsername(e.target.value)}/>
                    </Form.Group>
                    <Form.Group controlId="formBasicPassword">
                        <Form.Label>Password:</Form.Label>
                        <Form.Control type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
                    </Form.Group>
                    <Button variant="primary" onClick={handleLogin}>Đăng nhập</Button>
                </Form>
            </Col>
            <Col ></Col>
            <Col>
                <h1>Với sinh viên</h1>
                <Button variant="primary" onClick={handleLoginWithGoogle}>Đăng nhập bằng Gmail</Button>
            </Col>
        </Row>
    </Container>
    )
}

export default Login;