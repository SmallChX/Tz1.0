import React, { useState, useEffect } from 'react';
import { Form, Button, Modal } from 'react-bootstrap';
import axios from 'axios';

function BoothEdit({ show, onHide, companies, booth, onSave }) {
  const initialBoothState = {
    ID: booth ? booth.ID : null,
    company_info: {
      company_id: booth && booth.company_info ? booth.company_info.company_id : null,
      name: booth && booth.company_info ? booth.company_info.name : ""
    },
    level: booth ? booth.level : null,
    price: booth ? booth.price : null
  };
 
  const [editedBooth, setEditedBooth] = useState(initialBoothState);
  const [selectedCompany, setSelectedCompany] = useState(editedBooth.company_info.company_id || '');

  useEffect(() => {
    console.error(initialBoothState);
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name === "company_info.company_id") {
      // Xử lý cập nhật cho company_id
      setEditedBooth(prevState => ({
        ...prevState,
        company_info: {
          ...prevState.company_info,
          company_id: value
        }
      }));
    } else {
      // Xử lý cập nhật cho các trường khác (price và level)
      setEditedBooth(prevState => ({
        ...prevState,
        [name]: value
      }));
    }
  };
  

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      console.error(editedBooth);
      // Gửi dữ liệu booth đã chỉnh sửa tới endpoint /api/booth với phương thức PUT
      await axios.put('/api/booth', {
        booth_id: editedBooth.ID,
        // company_id: editedBooth.company_info.company_id,
        // company_name: booth.company_info.company_name,
        level: editedBooth.level,
        price: editedBooth.price,
      });
      // Gọi hàm onSave để thông báo rằng dữ liệu đã được lưu thành công
      onSave(editedBooth);
    } catch (error) {
      console.error('Error saving booth:', error);
    }
    setEditedBooth(initialBoothState);
  };

  if (!show) {
    return null;
  }

  return (
    <Modal show={show} onHide={onHide}>
      <Modal.Header closeButton>
        <Modal.Title>Edit Booth</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={handleSubmit}>
        <Form.Group controlId="company">
            <Form.Label>Company</Form.Label>
            <Form.Control as="select" name="company_info.company_id" value={selectedCompany} onChange={handleChange}>
              {/* Option cho phép không có công ty nào */}
              <option value="">Select company</option>
              {/* Duyệt qua danh sách công ty để tạo option */}
              {companies.map((company) => (
                <option key={company.id} value={company.id}>{company.name}</option>
              ))}
            </Form.Control>
          </Form.Group>
          <Form.Group controlId="price">
            <Form.Label>Price</Form.Label>
            <Form.Control type="number" name="price" value={editedBooth.price || ''} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="level">
            <Form.Label>Level</Form.Label>
            <Form.Control type="number" name="level" value={editedBooth.level || ''} onChange={handleChange} />
          </Form.Group>
          <Button variant="primary" type="submit">Save</Button>
        </Form>
      </Modal.Body>
    </Modal>
  );
};

export default BoothEdit;
