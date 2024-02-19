import React, { useEffect, useState } from 'react';
import { Modal, Button } from 'react-bootstrap';
import axios from 'axios';

const PaymentModal = ({ show, onHide, requestID }) => {
  const [paymentInfo, setPaymentInfo] = useState();
  useEffect(() => {
    async function fetchData() {
      const response = await axios.get(`/api/request/${requestID}/payment`)

      if (response.status === 200) {
        setPaymentInfo(response.data.data);
        console.log(response);
      } else {
        console.error("looix");
      }
    }
    if (requestID) {
      fetchData();
      console.log(requestID);
    }
  }, [requestID])

  return (
    <Modal show={show} onHide={onHide} centered>
      <Modal.Header closeButton>
        <Modal.Title>Thanh toán gian hàng</Modal.Title>
      </Modal.Header>
      <Modal.Body>
          {paymentInfo ? (
            <div>
              <h5>Tên công ty: {paymentInfo.company_name}</h5>
              <p>ID gian hàng: {paymentInfo.booths_id}</p>
              <p>Số tiền: {paymentInfo.amount} VND</p>
            </div>
          ) : (<></>)}
        <div className="text-center">
          {/* Giả định đây là QR code cho việc chuyển khoản */}
          <img src="path_to_your_qr_code_image" alt="QR Code" style={{ width: '200px', height: '200px' }} />
        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="primary" onClick={onHide}>Xác nhận</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default PaymentModal;
