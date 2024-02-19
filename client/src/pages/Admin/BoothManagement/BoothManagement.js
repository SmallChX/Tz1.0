import React, { useState, useEffect } from 'react';
import BoothList from './BoothList';
import BoothEdit from './EditBooth';
import axios from 'axios';

function BoothManager() {
  const [companies, setCompanies] = useState([]);
  const [booths, setBooths] = useState([]);
  const [selectedBooth, setSelectedBooth] = useState(null);
  const [showEditModal, setShowEditModal] = useState(false);

  useEffect(() => {
    async function fetchCompanyData() {
      try {
        const resCompany = await axios.get('/api/booth/company');
        if (resCompany.status === 200) {
          setCompanies(resCompany.data.data);
        } else {
          console.error("fail fetch company");
        }
      } catch (error) {
        console.error("error fetching company data:", error);
      }
    }
    fetchCompanyData();
  }, []);

  useEffect(() => {
    async function fetchBoothData() {
      try {
        const resBooth = await axios.get('/api/booth/get-all-booth');
        if (resBooth.status === 200) {
          // Lấy danh sách booths từ response
          const boothData = resBooth.data.data;
          // Sắp xếp booths theo ID trước khi set vào state
          boothData.sort((a, b) => a.ID - b.ID);
          // Set danh sách booths đã được sắp xếp vào state
          setBooths(boothData);
        } else {
          console.error("fail fetch company");
        }
      } catch (error) {
        console.error("error fetching booth data:", error);
      }
    }
    fetchBoothData();
  }, []);

  function handleEdit(booth) {
    setSelectedBooth(booth);
    setShowEditModal(true);
  }

  const handleSave = (updatedBooth) => {
    // Xử lý lưu trữ thông tin gian hàng đã chỉnh sửa
    setShowEditModal(false);
  };

  const handleCloseEditModal = () => {
    setSelectedBooth(null);
    setShowEditModal(false);
  };

  return (
    <div>
      <BoothList booths={booths} onEdit={handleEdit} />
      <BoothEdit
        show={showEditModal}
        onHide={handleCloseEditModal}
        booth={selectedBooth}
        companies={companies}
        onSave={handleSave}
      />
    </div>
  );
}

export default BoothManager;
