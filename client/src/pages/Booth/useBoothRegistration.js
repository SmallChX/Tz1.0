// useBoothRegistration.js
import { useState, useCallback, useEffect } from 'react';
import axios from 'axios';
import Swal from 'sweetalert2';

function useBoothRegistration(ownedBooths) {
    const [selectedBooths, setSelectedBooths] = useState([]);
    const [deselectedBooths, setDeselectedBooths] = useState([]);
    const [action, setAction] = useState('');

    useEffect(() => {
        setSelectedBooths([]);
        setDeselectedBooths([]);
    }, [action]);


  const toggleBoothSelection = useCallback((id) => {
    if (action === 'change') {
        if (ownedBooths.includes(id)) {
            // Kiểm tra nếu id đã tồn tại trong danh sách deselectedBooths
            if (deselectedBooths.includes(id)) {
                // Nếu đã tồn tại, loại bỏ khỏi danh sách và cho phép thao tác này mà không cần kiểm tra số lượng
                setDeselectedBooths(prev => prev.filter(boothId => boothId !== id));
            } else {
                // Kiểm tra trước khi thêm vào danh sách deselectedBooths
                if (selectedBooths.length + 1 > deselectedBooths.length + 1) {
                    Swal.fire("Lỗi", "Vui lòng bỏ chọn gian hàng mới trước khi thêm lại gian hàng cũ", "warning");
                } else {
                    setDeselectedBooths(prev => [...prev, id]);
                }
            }
        } else {
            // Kiểm tra nếu id không tồn tại trong danh sách selectedBooths và đảm bảo rằng
            // số lượng gian hàng chọn mới không vượt quá số lượng gian hàng bỏ chọn
            if (!selectedBooths.includes(id) && selectedBooths.length + 1 > deselectedBooths.length) {
                Swal.fire("Lỗi", "Vui lòng bỏ chọn số lượng gian hàng của bạn trước khi chọn gian hàng mới", "warning");
            } else {
                // Toggle trong danh sách selectedBooths cho các hành động khác
                setSelectedBooths(prev => prev.includes(id) ? prev.filter(boothId => boothId !== id) : [...prev, id]);
            }
        }
    } else if (action === 'remove') {
        // Xử lý cho đăng ký hoặc xóa
        if (!ownedBooths.includes(id)) {
            Swal.fire("Lỗi", "Vui lòng chọn gian hàng của bạn đang sở hữu để xóa", 'warning');
            return;
        }
        else setDeselectedBooths(prev => prev.includes(id) ? prev.filter(boothId => boothId !== id) : [...prev, id]);
    } else if (action === 'register') {
        if (!ownedBooths.includes(id)) {
            setSelectedBooths(prev => prev.includes(id) ? prev.filter(boothId => boothId !== id) : [...prev, id]);
        } 
    } else {
        Swal.fire("Lỗi", "Vui lòng chọn hành động ở bên để thực hiện", 'warning');
    }
        // console.log(deselectedBooths);
}, [selectedBooths, deselectedBooths, ownedBooths, action]);

    async function handleSubmit() {
        try {
            if (action === 'register') {
                const sortedSelectedBooths = [...selectedBooths].sort((a, b) => a - b);
                // Check if the selected booths are consecutive
                const selectedBoothIds = sortedSelectedBooths.map(Number);
                if (selectedBoothIds.length === 0) {
                    Swal.fire('Warning!', "Vui lòng chọn ít nhất một gian hàng để đăng ký", 'warning');
                    setSelectedBooths([]);
                    return;
                }
                console.log(sortedSelectedBooths);
                var sortedFinalBoothsID
                if (ownedBooths) {
                    const finalBoothsID = Array.from(new Set(selectedBoothIds.concat(ownedBooths.map(Number))));
                    sortedFinalBoothsID = [...finalBoothsID].sort((a, b) => a - b);
                }

                for (let i = 1; i < sortedFinalBoothsID.length; i++) {
                    if (sortedFinalBoothsID[i] !== sortedFinalBoothsID[i - 1] + 1) {
                        Swal.fire('Warning!',"Các gian hàng phải được chọn liên tiếp.", 'warning');
                        setSelectedBooths([]);
                        return;
                    }
                }
    
                if (sortedFinalBoothsID.length > 3) {
                    Swal.fire('Warning!',"Bạn chỉ có thể đăng ký tối đa 3 gian hàng.", 'warning');
                    setSelectedBooths([]);
                    return;
                }

                Swal.fire({
                    title: "Xác nhận",
                    html: "Bạn đăng ký " + (ownedBooths.length !== 0 ? "thêm" : "")  + " gian hàng <b>" + selectedBoothIds + "</b>?",
                    icon: "question",
                    confirmButtonText:"Xác nhận!",
                    preConfirm: async () => {
                        const response = await axios.post('/api/request', {
                            booth_id: selectedBoothIds,
                            type: 'regist',
                        })
                        if (response.status === 200) {
                            Swal.fire('Thành công', "Bạn đã đăng ký gian hàng thành công", 'success');
                        } else {
                            Swal.fire('Đã có vấn đề gì đấy?', 'Vui lòng thử lại', 'error');
                        }
                        
                    },
                    showCancelButton: true, 
                }).then((result) => {
                    if (result.isConfirmed) {
                        return;
                      } else if (result.isDenied) {
                        Swal.fire("Changes are not saved", "", "info");
                      }
                 
                });
                
            } else if (action === 'change') {
                const afterBooths = Array.from(new Set(selectedBooths.concat(ownedBooths.filter(boothId => !deselectedBooths.includes(boothId)))));
                const afterBoothsID = afterBooths.map(Number)
                console.log(afterBoothsID);
                const sortedBoothsID = [...afterBoothsID].sort((a, b) => a - b);
                Swal.fire({
                    title: 'Xác nhận',
                    html: 'Bạn xác nhận đổi từ gian hàng <b>' + ownedBooths.map(Number) + '</b> thành gian hàng <b>' + sortedBoothsID + '</b> chứ?',
                    icon: "question",
                    confirmButtonText:"Xác nhận!",
                    showCancelButton: true, 
                    preConfirm: async () => {
                        for (let i = 1; i < sortedBoothsID.length; i++) {
                            if (sortedBoothsID[i] !== sortedBoothsID[i - 1] + 1) {
                                Swal.fire('Warning!',"Các booth phải được chọn liên tiếp.", 'warning');
                                setSelectedBooths([]);
                                setDeselectedBooths([]);
                                return;
                            }
                        }
                        if (selectedBooths.length !== deselectedBooths.length) {
                            Swal.fire("warning!", "Số lượng gian hàng thay đổi phải như nhau", "warning");
                            setSelectedBooths([]);
                            setDeselectedBooths([]);
                            return;
                        }
                        const response = await axios.post('/api/request', {
                            booth_id: deselectedBooths.map(Number),
                            des_booth_id: selectedBooths.map(Number),
                            type: "change"
                        })
        
                        if (response.status === 200) {
                            Swal.fire('Thành công', "Bạn đã đăng ký gian hàng thành công", 'success');
                        } else {
                            Swal.fire('Đã có vấn đề gì đấy?', 'Vui lòng thử lại', 'error');
                        }
                    }
                }).then((result) => {
                    if (result.isConfirmed) {
                        return;
                    }
                });
                
            } else {
                Swal.fire({
                    title: "Xác nhận",
                    html: "Bạn có chắc muốn hủy gian hàng <b>" + deselectedBooths.map(Number) + "</b> chứ?",
                    input: "textarea",
                    inputLabel: "Vui lòng nói rõ lý do",
                    inputPlaceholder: "Điền lý do ở đây...",
                    showCancelButton: true,
                    preConfirm: async (reason) => {
                        const response = await axios.post('/api/request', {
                            booth_id: deselectedBooths.map(Number),
                            reason: reason,
                            type: 'remove'
                        })
                        if (response.status === 200) {
                            Swal.fire('Thành công', "Bạn đã gửi yêu cầu thành công", 'success');
                        } else {
                            Swal.fire('Đã có vấn đề gì đấy?', 'Vui lòng thử lại', 'error');
                        }
                    }
                }).then((result) => {
                    if (result.isConfirmed) {
                        return;
                    }
                });
                
            }
        } catch (error) {
            console.error("Lỗi", error);
        }
        setSelectedBooths([]);
        setDeselectedBooths([]);
    }

    return { selectedBooths, deselectedBooths, toggleBoothSelection, handleSubmit, setAction, action };
};

export default useBoothRegistration;
