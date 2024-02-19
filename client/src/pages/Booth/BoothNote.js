import React from "react";
import { Container, Col, Row } from "react-bootstrap";

function BoothNote() {
    return (
        <Container className="mt-5" style={{backgroundColor:"white", borderRadius:"20px"}}>
            <Row>
                <Col>
                    <Row className="align-items-center h-100">
                        <Col style={{maxWidth:"75px"}}>
                            <div className="booth-layoutgrass  plaque" style={{float:"right", height:"2em", width:"2em", borderRadius:"15%"}}></div>
                        </Col>
                        <Col className="px-0" style={{width:"100px"}}>
                            <p className="mb-0 p" style={{fontSize:"14px", fontWeight:"500"}}>Cỏ đường đi</p>
                           
                        </Col>

                    </Row>
                </Col>
                <Col>
                    <Row className="align-items-center">
                        <Col style={{maxWidth:"75px"}}>
                            <div className="booth-layoutindividual-booth cheap-booth" style={{float:"right"}}><span>0</span></div>
                        </Col>
                        <Col className="px-0 pt-2" style={{minWidth:"200px"}}>
                            <p className="mb-0" style={{fontSize:"14px", fontWeight:"500"}}>Gian hàng đơn: 12 000 000 VNĐ</p>
                            <p className="mb-2" style={{fontSize:"14px", fontWeight:"500"}}>Gian hàng đôi: 20 000 000 VNĐ</p>
                        </Col>

                    </Row>
                </Col>
                <Col>
                    <Row className="align-items-center">
                    <Col style={{maxWidth:"75px"}}>
                        <div className="booth-layoutindividual-booth expensive-booth" style={{float:"right"}}><span>0</span></div>
                    </Col>
                    <Col className="px-0 pt-2" style={{minWidth:"200px"}}>
                            <p className="mb-0" style={{fontSize:"14px", fontWeight:"500"}}>Gian hàng đơn: 15 000 000 VNĐ</p>
                            <p className="mb-2" style={{fontSize:"14px", fontWeight:"500"}}>Gian hàng đôi: 25 000 000 VNĐ</p>
                    </Col>
                    </Row>
                </Col>
            </Row>
        </Container>
    )
}

export default BoothNote;