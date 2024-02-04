import React, { useState } from 'react'
import axios from 'axios'

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [userRole, setUserRole] = useState('');

    const handleLogin = async () => {
        try {
            const response = await axios.post('/api/auth/login', {
                username, 
                password,
            });

            if (response.status === 200) {
                const data = response.data;
                setUserRole(data.role);

                localStorage.setItem('userRole', userRole);
            } else {
                console.error('Đăng nhập thất bại');
            }
        } catch (error) {
            console.error('Đã xảy ra lỗi', error);
        }
    };

    return (
        <div>
            <h1>Đăng nhập</h1>
            <div>
                <label>Username:</label>
                <input type="text" value={username} onChange={(e) => setUsername(e.target.value)}/>
            </div>
            <div>
                <label>Password:</label>
                <input type="text" value={password} onChange={(e) => setPassword(e.target.value)}/>
            </div>
            <button onClick={handleLogin}>Đăng nhập</button>
        </div>
    )
}

export default Login;