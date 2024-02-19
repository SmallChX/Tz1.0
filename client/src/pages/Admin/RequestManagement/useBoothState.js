import { useState } from 'react';

const useBoothsState = (initialBooths) => {
  // Khởi tạo trạng thái cho selectedBooths, deselectedBooths và booths
  const [selectedBooths, setSelectedBooths] = useState([]);
  const [deselectedBooths, setDeselectedBooths] = useState([]);
  const [booths, setBooths] = useState(initialBooths); // Sử dụng initialBooths từ BoothManager

  // Tạo danh sách gian hàng ảo từ booths truyền vào của boothmanager
  const virtualBooths = JSON.parse(JSON.stringify(initialBooths));


  // Trả về các trạng thái và hàm cập nhật tương ứng
  return {
    selectedBooths,
    setSelectedBooths,
    deselectedBooths,
    setDeselectedBooths,
    booths: virtualBooths, // Trả về danh sách gian hàng ảo
    setBooths,
  };
};

export default useBoothsState;
