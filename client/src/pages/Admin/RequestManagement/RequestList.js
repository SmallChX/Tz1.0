import React, { useState, useEffect } from "react";
import { Container, Table, Form, Button, Row, Col } from 'react-bootstrap';
import axios from "axios";
import Swal from "sweetalert2";

function RequestList({ setSearchTerm, showRejected, setShowRejected, filteredRequests, isBoothAvailable, handleAccept, handleReject, handleFinish, 
    isEditing, selectedBooths, setSelectedBooths, setDeselectedBooths, setBooths, booths }) {
  const [requestActions, setRequestActions] = useState({});
  const [editedRequests, setEditedRequests] = useState([]);

  const handleSelectChange = (id, action) => {
    const request = filteredRequests.find(req => req.id === id);
    const booth = booths.find(b => b.ID === request.booth_id[0]); // Giả sử chỉ có một booth_id

    // Kiểm tra xem action là "accept" và request có hợp lệ không
    if (action === "accept") {
        if (request.type === "regist") {
            const isBoothIdExists = request.booth_id.some(boothId => selectedBooths.includes(boothId));
            if (isBoothIdExists) {
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'Booth ID already selected.'
                });
                return; // Không thực hiện gì nếu có lỗi
            } else {
                setSelectedBooths(prev => [...prev, ...request.booth_id]);
                setBooths(prev => prev.map(b => b.ID === booth.ID ? { ...b, company_info: { ...b.company_info, company_id: 0 } } : b));
            }
        } else if (request.type === "change") {
            const isDesBoothIdExists = request.des_booth_id.some(boothId => selectedBooths.includes(boothId));
            if (isDesBoothIdExists) {
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'Desired Booth ID already selected.'
                });
                return; // Không thực hiện gì nếu có lỗi
            } else {
                setSelectedBooths(prev => [...prev, ...request.des_booth_id]);
                setDeselectedBooths(prev => [...prev, ...request.booth_id]);
                setBooths(prev => prev.map(b => b.ID === booth.ID ? { ...b, company_info: { ...b.company_info, company_id: 0 } } : b));
            }
        } else if (request.type === "remove") {
            setDeselectedBooths(prev => [...prev, ...request.booth_id]);
            setBooths(prev => prev.map(b => b.ID === booth.ID ? { ...b, company_info: { ...b.company_info, company_id: 0 } } : b));
        }
    }

    // Kiểm tra xem action là "reject" và có phải là hành động từ trạng thái "accept" không
    if (action === "reject" || action === "") {
        const prevAction = requestActions[id];
        if (prevAction === "accept") {
            if (request.type === "regist" || request.type === "change") {
                setSelectedBooths(prev => prev.filter(boothId => !request.booth_id.includes(boothId)));
                setBooths(prev => prev.map(b => b.ID === booth.ID ? { ...b, company_info: { ...b.company_info, company_id: null } } : b));
            } else if (request.type === "remove") {
                setDeselectedBooths(prev => prev.filter(boothId => !request.booth_id.includes(boothId)));
                setBooths(prev => prev.map(b => b.ID === booth.ID ? { ...b, company_info: { ...b.company_info, company_id: null } } : b));
            }
        }
    }
    if (action === "") {
      // Loại bỏ request khỏi danh sách editedRequests
      setEditedRequests(prev => prev.filter(item => item.id !== id));
    }

    // Cập nhật hành động của request
    setRequestActions(prev => ({ ...prev, [id]: action }));

    // Kiểm tra xem request đã được chỉnh sửa trước đó chưa
    const existingIndex = editedRequests.findIndex(item => item.id === id);
    if (existingIndex !== -1) {
        // Nếu request đã tồn tại, cập nhật hành động của nó
        setEditedRequests(prev => prev.map(item => (item.id === id ? { ...item, action } : item)));
    } else {
        // Nếu không, thêm request mới vào mảng editedRequests
        setEditedRequests(prev => [...prev, { id, action }]);
    }
};



  const handleSubmitAction = (id) => {
    const action = requestActions[id];
    if(action === "accept") {
      handleAccept(id);
    } else if(action === "reject") {
      handleReject(id);
    } else if(action === "finish") {
      handleFinish(id);
    }
  };

  // Kiểm tra xem booth có hợp lệ không
  const isBoothValid = (request) => {
    if (request.type === "regist" || request.type === "change") {
      const boothIDs = request.type === "regist" ? request.booth_id : request.des_booth_id;
      return boothIDs.every(isBoothAvailable);
    } else if (request.type === "remove") {
      // Đối với remove, booth phải đã đăng ký (không hợp lệ nếu booth trống)
      return request.booth_id.some(isBoothAvailable);
    }
    return true;
  };

  const handleSubmitAll = () => {
    // Lọc ra các request đã chỉnh sửa
    const editedRequestsToSend = filteredRequests.filter(request => editedRequests.includes(request.id));
    // Thêm thông tin về hành động đã chọn vào mỗi request
    const requestsWithAction = editedRequestsToSend.map(request => ({
      ...request,
      action: requestActions[request.id] // Thêm thông tin về hành động vào từng request
    }));
    console.log(editedRequests);
    // Gửi các request đã chỉnh sửa tới endpoint
    axios.put('/api/request/handle-list', editedRequests)
      .then(response => {
        if (response.status === 200) {
          Swal.fire("Thành công", "Cập nhật thông tin thành công", "success");
          // window.location.reload();
        }
      })
      .catch(error => {
        Swal.fire("Oops!", error, "warning");
      });
  };

  return (
    <div>
      <Container>
        <h1>Quản lý các request</h1>
        <Row className="mb-3">
          <Col>
            <Form.Control
              type="text"
              placeholder="Tìm kiếm theo tên công ty..."
              onChange={e => setSearchTerm(e.target.value)}
            />
          </Col>
          <Col>
            <Form.Check
              type="switch"
              id="custom-switch"
              label="Hiển thị yêu cầu bị từ chối"
              checked={showRejected}
              onChange={() => setShowRejected(!showRejected)}
            />
          </Col>
        </Row>
        <div style={{overflow:"auto", maxHeight:"600px"}}>
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>#</th>
              <th>Loại</th>
              <th>ID Gian Hàng</th>
              <th>ID Gian Hàng Mới</th>
              <th>Tên công ty</th>
              <th>Thao tác</th>
            </tr>
          </thead>
          <tbody>
            {filteredRequests.map((request, index) => {
              const validBooth = isBoothValid(request);
              return (
                <tr key={index} className={editedRequests.some(item => item.id === request.id) ? "table-success" : ""}>
                  <td>{request.id}</td>
                  <td>{request.type}</td>
                  <td>{request.booth_id.join(", ")}</td>
                  <td>{request.des_booth_id ? request.des_booth_id.join(", ") : "N/A"}</td>
                  <td>{request.company_name}</td>
                  <td>
                    {!validBooth && request.status === "pending" ? (
                        <div>
                            <span>Không hợp lệ</span>
                            <Button variant="warning" onClick={() => handleReject(request.id)}>Từ chối</Button>
                        </div>
                    ) : (
                      <>
                        {request.status === "pending" && (
                          <>
                            <Form.Select aria-label="Default select example" value={requestActions[request.id] || ""} onChange={(e) => isEditing && handleSelectChange(request.id, e.target.value)}>
                              <option value="">Chọn hành động</option>
                              <option value="accept">Chấp nhận</option>
                              <option value="reject">Từ chối</option>
                            </Form.Select>
                            <Button variant="primary" onClick={() => handleSubmitAction(request.id)}>Xác nhận</Button>
                          </>
                        )}
                        {request.status === "accepted" && (
                          <Button variant='success' onClick={() => handleFinish(request.id)}>Hoàn thành</Button>
                        )}
                      </>
                    )}
                    {request.status === "rejected" && <span>Đã từ chối</span>}
                    {request.status === "finished" && <span>Đã hoàn thành</span>}
                  </td>
                </tr>
              );
            })}
          </tbody>
        </Table>
        </div>
        {isEditing && (
          <Button variant="primary" onClick={handleSubmitAll}>Chấp nhận tất cả</Button>
        )}
      </Container>
    </div>
  );
}

export default RequestList;
