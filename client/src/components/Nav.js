import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from 'axios';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import React, { useState, useEffect } from 'react';


function NavbarDisplay() {
    const [navBackground, setNavBackground] = useState('transparent');
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        const changeBackground = () => {
            if (location.pathname === '/home') {
                if (window.scrollY >= 80) {
                    setNavBackground('solid');
                } else {
                    setNavBackground('transparent');
                }
            }
        };

        window.addEventListener('scroll', changeBackground);

        changeBackground();

        return () => window.removeEventListener('scroll', changeBackground);
    }, [location]);

    const handleLogout = async () => {
        try {
            const response = await axios.post('/api/auth/logout');
            
            if (response.status === 200) {
                localStorage.removeItem("userRole");
                navigate("/login", { replace: true });
            } 
        } catch (error) {
            console.error('Đã xảy ra lỗi', error);
        }
    }

    let currentUserRole;
    if (localStorage.getItem("userRole")) {
        currentUserRole = localStorage.getItem("userRole");
    }
    return (
        <div>
             <Navbar bg={location.pathname === '/home' && navBackground === 'transparent' ? 'transparent' : 'primary'} variant="dark" expand="lg" fixed="top" style={{ transition: '0.5s ease'}}>
                <Container>
                    <Navbar.Brand href="#home">Navbar</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="/home">Home</Nav.Link>
                            <Nav.Link href="/booth">Gian hàng</Nav.Link>
                            <Nav.Link href="/admin">Quản lý yêu cầu</Nav.Link>
                            <Nav.Link href="/account">Quản lý tài khoản</Nav.Link>
                            <Nav.Link href="/admin/booth"> Quản lý gian hàng</Nav.Link>
                        </Nav>
                        <Nav>
                            {currentUserRole && (
                                <>
                                    <Nav.Link href="/profile">Profile</Nav.Link>
                                    <button className="btn btn-outline-light ms-2" onClick={handleLogout}>
                                        Logout
                                    </button>
                                </>
                            )}
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>

        </div>

        
    )
}

export default NavbarDisplay