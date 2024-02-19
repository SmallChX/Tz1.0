import { ButtonGroup, ToggleButton, Table, Button, Form, Row, Col } from "react-bootstrap";
import React, { useEffect, useState } from "react";
import axios from "axios";
import Swal from "sweetalert2"
import PaymentModal from "./Payment";
import ActionModal from "./ActionModal";
import "./booth_manager.css";

function BoothManager({ selectedBooths, ownedBooths, setAction, handleSubmit }) {
    const [requests, setRequests] = useState([]);
    const [deleted, setDeleted] = useState(false);
    const [showPayment, setShowPayment] = useState(false);
    const [requestPaymentID, setRequestPaymentID] = useState();
    const [showRejected, setShowRejected] = useState(false);
    
    // const [showActionModal, setShowActionModal] = useState(false);


    useEffect(() => {
        async function fetchData() {
            try {
                const response = await axios.get("/api/request/company")
                if (response.status === 200) {
                    const sortedData = response.data.data.sort((a, b) => {
                        const statusOrder = { "accepted": 1, "pending": 2, "rejected": 3 };
                        return statusOrder[a.status] - statusOrder[b.status];
                    });
                    setRequests(sortedData);
                } else {
                    console.error("fail fetch");
                }
            } catch (error) {
            console.error(error);
            }
        }
        fetchData();
    }, [deleted])

    function handleDelete(index) {
        Swal.fire({
            title: "Xác nhận",
            text: "Bạn chắc chắn muốn xóa yêu cầu này chứ?",
            icon: "warning",
            confirmButtonText: "Đồng ý",
            showCancelButton: true,
            preConfirm: async () => {
                try {
                    const response = await axios.delete(`/api/request/${index}`);
                    if (response.status === 200) {
                        Swal.fire("Thành công", "Đã xóa yêu cầu thành công", "success");
                        setDeleted(!deleted);
                    } else {
                        Swal.fire("Lỗi", "Đã xảy ra lỗi, vui lòng thử lại", "error");
                    }
                } catch (error) {
                    Swal.fire("Lỗi", error, 'error');
                }
            },
        })
    }

   

     // Hàm mở modal
  function handleOpenModal(requestID) {
    setRequestPaymentID(requestID);
    setShowPayment(true);
    }
  // Hàm đóng modal
  const handleCloseModal = () => {
    setShowPayment(false);
    setRequestPaymentID("");
  } 

  const tableWrapperStyle = {
    maxHeight: '400px', // Set this to the maximum height you want for the table
    overflowY: 'auto',   // This will add a scrollbar to the table when the content overflows
    overflowX: "hidden",
};

    return (
        <div className="mx-5" style={{width:"400px", minHeight:"300px", backgroundColor:"white", borderRadius:"15px"}}> 
            <h2>Danh sách Request</h2>
                <Row>
                <Col>
                    <p className="float-right">Hiển thị yêu cầu bị từ chối:</p>
                </Col>
                <Col xs="auto">
                    <Form.Check
                        type="switch"
                        id="custom-switch"
                        label=""
                        className="float-right"
                        checked={showRejected}
                        onChange={(e) => setShowRejected(e.currentTarget.checked)}
                    />
                </Col>
            </Row>
            <div style={tableWrapperStyle}>
            <PaymentModal
                show={showPayment}
                onHide={handleCloseModal}
                requestID={requestPaymentID}
            />
            <Table striped bordered hover className="custom-table">
                <thead className="table-fixed-header">
                    <tr>
                        <th>#</th>
                        <th>Loại</th>
                        <th>Gian Hàng</th>
                        <th>Trạng thái</th>
                        <th>Thao tác</th>
                    </tr>
                </thead>
                <tbody>
                    {requests
                      .filter(request => showRejected || request.status !== "rejected").map((request, index) => {
                        const rowClass = request.status === "rejected" ? "table-danger" 
                          : request.status === "accepted" ? "table-warning" 
                          : request.status === "finished" ? "table-success" : "";
                          return (
                        <tr key={index} className={rowClass}>
                            <td>{index + 1}</td>
                            <td>
                                {request.type === "regist" ? (<div>Đăng ký</div>) :
                                request.type === "change" ? (<div>Thay đổi</div>) :
                                (<div>Bỏ</div>)}</td>
                            <td><div className="booth-layoutrow">
                                {request.booth_id.map((id) => (
                                    <span key={id} className={`booth-layoutindividual-booth ${ownedBooths.includes(id) ? 'booth-layoutowned' : ''} `}><span>{id}</span></span>
                                ))}
                                </div>
                            </td>
                            {/* <td><div className="booth-layoutrow">
                                { request.des_booth_id ? (request.des_booth_id.map((id) => ( 
                                    <span key={id} className="booth-layoutindividual-booth"><span>{id}</span></span>)
                                )) : (<div></div>)}
                                </div>
                            </td> */}
                            <td>
                                {request.status === "accepted" ? (<div>Chờ thanh toán</div>) :
                                request.status === "rejected" ? (<div>Bị từ chối</div>) :
                                request.status === "finished" ? (<div>Đã hoàn thành</div>):
                                (<div>Chờ xử lý</div>)}
                            </td>
                            <td>
                                { request.status === "accepted" ? (<Button className="btn-booth" variant="success" onClick= {() => handleOpenModal(request.id)} >Thanh toán</Button>) :
                                request.status === "finished" || request.status === "rejected" ? (<div>Yêu cầu đã được xử lý</div>) :
                                (<Button variant="danger" className="btn-booth" onClick={() => handleDelete(request.id)}>Xóa yêu cầu</Button>)
                                }
                            </td>
                        </tr>
                    )})}
                </tbody>
            </Table>
            </div>
        </div>
    )
}

export default BoothManager