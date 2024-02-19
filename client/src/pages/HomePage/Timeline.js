import React from 'react';
import { Container, Row, Col, Card, Button } from 'react-bootstrap';
import { useInView } from 'react-intersection-observer';
import { motion } from 'framer-motion';
import jamilaImage from './images/black-chick.jpg';
import sandraImage from './images/mathew-kane.jpg';
import djCrazyheadImage from './images/eric-ward.jpg';
import ArtistCard from './TimeLineCard';

function LineupArtists() {

    const artists = [
        {
          name: "Khai mạc Ngày hội Việc làm CSE Job Fair 2024",
          image: jamilaImage,
          description: "Đây là mô tả của Khai mạc Ngày hội Việc làm CSE Job Fair 2024",
          link: "#"
        },
        {
          name: "Gian hàng triển lãm và Không gian Doanh nghiệp",
          image: sandraImage,
          description: "Đây là mô tả của gian hàng triễn làm và không gian doanh nghiệp",
          link: "#"
        },
        {
          name: "Hội thảo Công nghệ và Định hướng cho sinh viên khóa 2023",
          image: djCrazyheadImage,
          description: "Đây là mô tả của Hội thảo Công nghệ và Định hướng cho sinh viên khóa 2023.",
          link: "#"
        }
        // Thêm các nghệ sĩ khác theo cùng định dạng
      ];

  return (
        <Container>
            <Row>
                <Col className='lineup-artists-headline pt-5'>
                <div className="entry-title">
                    <h2>Nội dung chương trình</h2>
                    {/* <p>Job Fair</p> */}
                </div>
                </Col>
            </Row>
            <div className='lineup-artists'>
                {artists.map((artist, index) => (
                    <ArtistCard key={index} artist={artist} index={index} />
                ))}
            </div>
        </Container>

  );
}

export default LineupArtists;
