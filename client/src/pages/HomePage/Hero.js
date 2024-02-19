import React from 'react';
import { Container, Row, Col, Button } from 'react-bootstrap';
import Countdown from './Countdown';

function HeroContent() {
  return (
    <div className="hero-content">
      <Container>
        <Row>
          <Col xs={12} style={{ position: 'relative', zIndex: 1 }} >
            
            
            
            
            {/* <div className="entry-footer mt-3">
              <Button variant="primary" className="mr-2">Buy Tickets</Button>
              <Button variant="secondary">See Lineup</Button>
            </div> */}
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default HeroContent;
