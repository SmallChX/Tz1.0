import React from "react";
import image from "./images/about-us-content-image.jpg";
import { useInView } from 'react-intersection-observer';
import { motion } from 'framer-motion';

function Introduction() {
  const { ref, inView } = useInView({
    triggerOnce: true, // Chỉ kích hoạt hiệu ứng một lần khi phần tử lần đầu tiên vào viewport
    threshold: 0.5, // Phần tử sẽ được coi là trong viewport khi 50% của nó xuất hiện
  });
    return (
        <div className="container my-5" ref={ref}>
      <div className="row align-items-center">
        {/* Cột cho văn bản giới thiệu */}
        <motion.div
          className="col-md-6"
          initial={{ opacity: 0, x: -100 }}
          animate={inView ? { opacity: 1, x: 0 } : {}}
          transition={{ duration: 0.8 }}
        >
          <h2 className="mb-3">Job Fair 2024</h2>
          <h6>30.3.2024</h6>
          <p className="mb-4">
          “Ngày hội Việc làm - CSE Job Fair” đã từ lâu trở thành một chương trình thường niên của khoa Khoa học và Kỹ thuật Máy tính. Đây được xem như cầu nối doanh nghiệp và các bạn sinh viên, giúp sinh viên tìm hiểu về môi trường các chương trình thực tập, làm việc của doanh nghiệp. Đồng thời, mở rộng cơ hội để doanh nghiệp tiếp xúc với sinh viên. Theo sự chỉ đạo của Ban Chủ nhiệm Khoa, Đoàn thanh niên - Liên chi Hội Sinh viên khoa Khoa học và Kỹ thuật Máy tiếp tục tổ chức chương trình “Ngày hội Việc làm CSE Job Fair 2024”. </p>
          
          <div>
            <button className="btn btn-outline-dark btn-lg">Xem thêm</button>
          </div>
        </motion.div>
        {/* Cột cho hình ảnh */}
        <motion.div
          className="col-md-6"
          initial={{ opacity: 0, x: 100 }}
          animate={inView ? { opacity: 1, x: 0 } : {}}
          transition={{ duration: 0.8 }}
        >
          <img src={image} alt="Conference" className="img-fluid" />
          </motion.div>
      </div>
    </div>
    )
}

export default Introduction