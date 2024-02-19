import React, { useState } from 'react';
import { ButtonGroup, ToggleButton, Table, Button, Form, Row, Col, Card } from "react-bootstrap";

function ActionCard({ actionType, setAction, booths, ownedBooths, selectedBooths, deselectedBooths, handleSubmit}) {
    const [radioValue, setRadioValue] = useState('');
    const radios_button = [
        { name: 'Đăng ký', value: '1', type: "register", variant: "success"},
        { name: 'Thay đổi', value: '2', type: "change", variant: "primary"},
        { name: 'Xóa', value: '3', type: "remove", variant: "warning"}
    ];

    const fixedStyle = {
        backgroundColor: "white",
        borderRadius: "15px",
        height: "300px", // Fixed height
        width: "500px", // Fixed width
        overflow: "auto", // Add scrollbars when content exceeds the container's height
        padding: "1rem",
        display: "flex",
    };

    const calculateBoothsPrice = (boothIds) => {
        const priceCounts = {};
    
        // Tính số lượng gian hàng cho mỗi mức giá
        boothIds.forEach((boothId) => {
            const booth = booths.find(booth => booth.ID === boothId);
            if (booth) {
                priceCounts[booth.price] = (priceCounts[booth.price] || 0) + 1;
            }
        });
    
        // Tính toán tổng giá tiền
        let totalPrice = 0;
        Object.entries(priceCounts).forEach(([price, count]) => {
            const priceInt = parseInt(price);
            // Áp dụng giá đặc biệt cho mỗi cặp gian hàng
            const pairs = Math.floor(count / 2);
            totalPrice += pairs * (5 / 3) * priceInt;
            // Nếu còn lại số lẻ gian hàng, thêm vào giá thông thường
            totalPrice += (count % 2) * priceInt;
        });
    
        return totalPrice;
    };

    const calculatePriceDifference = () => {
        const selectedPrice = calculateBoothsPrice(selectedBooths);
        const deselectedPrice = calculateBoothsPrice(deselectedBooths);
    
        return selectedPrice - deselectedPrice;
    };
    
    const formatCurrency = (value) => {
        // 'vi-VN' là mã ngôn ngữ và địa phương cho tiếng Việt, 'VND' là mã tiền tệ Việt Nam đồng
        return value.toLocaleString('vi-VN', { style: 'currency', currency: 'VND' });
    };

    const priceDifference = calculatePriceDifference();
    let priceMessage = "Không có sự thay đổi về giá.";
    if (priceDifference > 0) {
        priceMessage = `Trả thêm: ${formatCurrency(priceDifference)}`;
    } else if (priceDifference < 0) {
        priceMessage = `Được hoàn lại: ${formatCurrency(Math.abs(priceDifference))}`;
    }


    let cardTitle = "";
    let cardBody = null;

    if (!actionType) {
        return (
            <div style={fixedStyle} >
                <div style={{flex: 8}}>
                    <h2>Danh sách gian hàng sở hữu:</h2>
                    {ownedBooths.length > 0 ? (
                            <div className='booth-layoutrow'>
                                {ownedBooths.map((ID) => (
                                    <div key={ID} className={`booth-layoutindividual-booth ${ownedBooths.includes(ID) ? 'booth-layoutowned' : ''}`}>
                                        <span>{ID}</span>
                                    </div>
                                ))}
                                </div>
                    ) : (
                        <p>Bạn chưa sở hữu gian hàng nào.</p>
                    )}
                </div>
                <div style={{ flex: 1.5, marginTop: 'auto', display: 'flex', justifyContent: 'center', bottom:"0px"}}>
                    {radios_button.map((radio, idx) => (
                        <Button
                            key={idx}
                            id={`button-radio-${idx}`}
                            type="radio"
                            variant={radio.variant}
                            name="radio-button"
                            value={radio.value}
                            checked={radioValue === radio.value}
                            onChange={(e) => setRadioValue(e.currentTarget.value)}
                            onClick={() => setAction(radio.type)}
                            style={{ marginRight: '5px' }} // Adjust margin as needed
                        >
                            {radio.name}
                        </Button>
                    ))}
                </div>
            </div>
        );
    }

    // Hàm tính tổng giá tiền cho các gian hàng đã chọn
    const calculateTotalPrice = () => {
        return calculateBoothsPrice(selectedBooths);
    };

    switch (actionType) {
        case "register":
            cardTitle = "Đăng ký mới";
            cardBody = (
                <>
                    <Form.Group controlId="selectedBooth">
                        <Form.Label>Gian hàng đang chọn:</Form.Label>
                        <div className="booth-layout row">
                            {!selectedBooths || selectedBooths.length === 0 ? (
                                <div>Chưa có gian hàng nào được chọn</div>
                            ) : (
                                <div className='booth-layoutrow'>
                                {selectedBooths.map((ID) => (
                                    <div key={ID} className={`booth-layoutindividual-booth ${ownedBooths.includes(ID) ? 'booth-layout owned' : ''}`}>
                                        <span>{ID}</span>
                                    </div>
                                ))}
                                </div>
                            )}
                        </div>
                    </Form.Group>
                    <Form.Group controlId="price">
                        <Form.Label>Báo giá:</Form.Label>
                        <Form.Control type="text" value={formatCurrency(calculateTotalPrice())} readOnly />
                    </Form.Group>
                </>
            );
            break;
        case "change":
            cardTitle = "Thay đổi gian hàng";
            cardBody = (
                <>
                    <Form.Group controlId="deselectedBooth">
                        <Form.Label>Gian hàng muốn bỏ:</Form.Label>
                        <div className="booth-layoutrow">
                                {deselectedBooths.map((ID) => (
                                    <div key={ID} className={`booth-layoutindividual-booth ${ownedBooths.includes(ID) ? 'booth-layoutowned' : ''}`}>
                                        <span>{ID}</span>
                                    </div>
                                ))}
                        </div>
                    </Form.Group>
                    <Form.Group controlId="selectedBooth">
                        <Form.Label>Gian hàng mới:</Form.Label>
                        <div className='booth-layoutrow'>
                                {selectedBooths.map((ID) => (
                                    <div key={ID} className={`booth-layoutindividual-booth ${ownedBooths.includes(ID) ? 'booth-layoutowned' : ''}`}>
                                        <span>{ID}</span>
                                    </div>
                                ))}
                        </div>
                    </Form.Group>
                    <Form.Group controlId="priceDifference">
                        <Form.Label>Báo giá:</Form.Label>
                        <Form.Control type="text" value={priceMessage} readOnly />
                    </Form.Group>
                </>
            );
            break;
        case "remove":
            cardTitle = "Xóa gian hàng";
            cardBody = (
                <>
                    <Form.Group>
                        <Form.Label>Chọn gian hàng muốn xóa:</Form.Label>
                        <div className='booth-layoutrow'>
                                {deselectedBooths.map((ID) => (
                                    <div key={ID} className={`booth-layoutindividual-booth ${ownedBooths.includes(ID) ? 'booth-layoutowned' : ''}`}>
                                        <span>{ID}</span>
                                    </div>
                                ))}
                        </div>
                        <Form.Label>Tiền hoàn lại:</Form.Label>
                            <p>{formatCurrency(calculateRefund())}</p>
                    </Form.Group>
                </>
            );
            break;
        // Các case khác cho các actionType khác (nếu có)
    }

    // Hàm tính tổng tiền được hoàn lại khi bỏ chọn các gian hàng
    const calculateRefund = () => {
        return calculateBoothsPrice(deselectedBooths);
    };

    function submit() {
        handleSubmit();
        setAction("");
    }

    return (
        <div style={fixedStyle}>
            <Card>
                <Card.Header>
                    {cardTitle}
                </Card.Header>
                <Card.Body>{cardBody}</Card.Body>
                <Card.Footer>
                    <Button onClick={submit}>Xác nhận</Button>
                    <Button variant="secondary" onClick={() => setAction("")}>
                        Hủy
                    </Button>
                </Card.Footer>
            </Card>
        </div>
    );
}

export default ActionCard;
