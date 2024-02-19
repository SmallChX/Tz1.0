import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Swal from 'sweetalert2';
import RequestList from "./RequestList"
import BoothsLayout from '../../Booth/BoothsLayout';
import { Col, Container, FormCheck, Row } from 'react-bootstrap';
import useBoothsState from './useBoothState'; // Import hook

function RequestManagement() {
  const [requests, setRequests] = useState([]);
  const [handling, setHandling] = useState(false);
  const [booths, setBooths] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [showRejected, setShowRejected] = useState(false);
  const [ownedBooths, setOwnedBooths] = useState([]);
  const [isEditing, setIsEditing] = useState(false);

  // Function to reset all states to initial state
  const resetStates = () => {
    window.location.reload();
  };

  // Function to handle toggle edit mode
  const handleToggleEdit = () => {
    setIsEditing(prevState => !prevState);
    if (isEditing) {
      resetStates();
    }
  };

  
  useEffect(() => {
    // Gọi API để lấy danh sách các request từ /api/request/get-all-request
    axios.get('/api/request/get-all-request')
      .then(response => {
        setRequests(response.data.data);
      })
      .catch(error => {
        console.error(error);
      });

      axios.get('/api/booth/get-all-booth')
      .then(response => {
        const sortedBooths = response.data.data.sort((a, b) => a.ID - b.ID);
        setBooths(sortedBooths);
      }).catch(error => {
        console.error(error);
      })
      
  }, [handling]);

  function isBoothAvailable(boothID) {
    const booth = booths.find(b => b.ID === boothID);
    return booth && (booth.company_info.company_id === 0 || booth.company_info.company_id === null);
  }

  function handleAccept(requestId) {
    Swal.fire({
      title: "Xác nhận đồng ý",
      text: "Bạn đồng ý với yêu cầu này?",
      icon: "question",
      showCancelButton: true,
      confirmButtonText: "Đồng ý",
      preConfirm: async () => {
        try {
          const response = await axios.put("/api/request/accept", {
            request_id: requestId,
          }) 
          if (response.status === 200) {
            Swal.fire("Thành công", "Xác nhận thành công", "success");
            setHandling(!handling);
          } else {
            Swal.fire("Thất bại", "Đã có vấn đề, vui lòng kiểm tra lại", "error");
          }
        } catch (error) {
          Swal.fire("Oops", "Lỗi hệ thống", "error");
        }
      }
      
    })
  };

  const handleReject = (requestId) => {
    Swal.fire({
      title: "Xác nhận từ chối",
      text: "Bạn từ chối với yêu cầu này?",
      icon: "question",
      showCancelButton: true,
      confirmButtonText: "Đồng ý",
      preConfirm: async () => {
        try {
          const response = await axios.put("/api/request/reject", {
            request_id: requestId,
          }) 
          if (response.status === 200) {
            Swal.fire("Thành công", "Xử lý thành công", "success");
            setHandling(!handling);
          } else {
            Swal.fire("Thất bại", "Đã có vấn đề, vui lòng kiểm tra lại", "error");
          }
        } catch (error) {
          Swal.fire("Oops", "Lỗi hệ thống", "error");
        }
      }})
  };

  const handleFinish = (requestId) => {
    Swal.fire({
      title: "Xác nhận đăng ký",
      text: "Bạn chắc chắn hoàn thành yêu cầu này?",
      icon: "question",
      showCancelButton: true,
      confirmButtonText: "Đồng ý",
      preConfirm: async () => {
        try {
          const response = await axios.put("/api/request/finish", {
            request_id: requestId,
          }) 
          if (response.status === 200) {
            Swal.fire("Thành công", "Xử lý thành công", "success");
            setHandling(!handling);
          } else {
            Swal.fire("Thất bại", "Đã có vấn đề, vui lòng kiểm tra lại", "error");
          }
        } catch (error) {
          Swal.fire("Oops", "Lỗi hệ thống", "error");
        }
      }})
  };

  const filteredRequests = requests.filter(request => {
    return request.company_name.toLowerCase().includes(searchTerm.toLowerCase()) &&
           (showRejected || request.status !== "rejected");
  });

  function toggleBoothSelection() {

  }

  const {
    selectedBooths,
    setSelectedBooths,
    deselectedBooths,
    setDeselectedBooths,
    virtualBooths,
    setVirtualBooths,
  } = useBoothsState(booths); // Sử dụng hook

  return (
      <Row style={{paddingTop:"100px", height:"100hv"}}>
        <Col style={{maxWidth:"400px"}}>
            
          <RequestList 
          setSearchTerm ={setSearchTerm}
          showRejected = {showRejected} 
          setShowRejected = {setShowRejected} 
          filteredRequests = {filteredRequests} 
          isBoothAvailable= {isBoothAvailable} 
          handleAccept = {handleAccept} 
          handleReject = {handleReject} 
          handleFinish = {handleFinish}
          isEditing = {isEditing}
          setSelectedBooths={setSelectedBooths}
          setDeselectedBooths={setDeselectedBooths}
          selectedBooths={selectedBooths}
          setBooths={setBooths}
          booths={booths}
          />
  
        </Col>
        <Col>
        <div className='mt-5'>
            Mode chỉnh sửa
          <FormCheck 
            type="switch"
            id="custom-switch"
            label={isEditing ? 'Disable Edit Mode' : 'Enable Edit Mode'}
            onChange={handleToggleEdit}
            checked={isEditing}
            />
          </div>
        <BoothsLayout 
          booths={booths}
          selectedBooths={selectedBooths}
          deselectedBooths={deselectedBooths}
          ownedBooths={ownedBooths}
          toggleBoothSelection={toggleBoothSelection}
        />
        
        </Col>
      </Row>
      
  );
};

export default RequestManagement;
