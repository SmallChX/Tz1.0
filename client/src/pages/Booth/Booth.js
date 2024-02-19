import React, { useState, useEffect } from "react";
import BoothsLayout from "./BoothsLayout";
import useBoothRegistration from "./useBoothRegistration";
import BoothManager from "./BoothManager";
import axios from "axios";
import { Container, Row, Col } from 'react-bootstrap';
import ActionModal from "./ActionModal"
import BoothNote from "./BoothNote";

function Booth() {
    const [ownedBooths, setOwnedBooths] = useState([]);
    const [booths, setBooths] = useState([]);
    // Bạn cần lấy danh sách ownedBooths từ API hoặc từ một nguồn khác

    useEffect(() => {
        const fetchOwnedBooths = async () => {
            try {
                const response = await axios.get('/api/booth/company-owned-booth');
                if (response.status === 200) {
                    const sortedBooths = response.data.data.sort((a, b) => a.ID - b.ID);
                    setOwnedBooths(sortedBooths);
                } else {
                    console.error("Failed to fetch booths");
                }
            } catch (error) {
                console.error("Error fetching booths", error);
            }
        };
        const fetchAllBooths = async () => {
            try {
                const response = await axios.get('/api/booth/get-all-booth');
                if (response.status === 200) {
                    const sortedBooths = response.data.data.sort((a, b) => a.ID - b.ID);
                    setBooths(sortedBooths);
                } else {
                    console.error("Failed to fetch booths");
                }
            } catch (error) {
                console.error("Error fetching booths", error);
            }
        };
        fetchOwnedBooths();
        fetchAllBooths();
    }, []);

    const style = {
        height: '100vh', // This makes sure that the row is as tall as the viewport
        display: 'flex',
        flexDirection: 'column', // Stack children vertically
        justifyContent: 'center', // Center children vertically
    };

    const parentStyle = {
        paddingTop:"40px",
        display: 'flex',
        alignItems: 'center', // Center vertically
        justifyContent: 'center', // Center horizontally
        minHeight: '100vh', // Full viewport height
        backgroundColor: "#eff7f6",
        margin: "0",
        padding: "0",
    };

    const { selectedBooths, deselectedBooths, toggleBoothSelection, handleSubmit, setAction, action } = useBoothRegistration(ownedBooths);
    return (
        <div style={parentStyle}>
            <Container fluid style={{paddingTop:"0px"}}>
            <Row className="g-4 justify-content-center justify-content-center"> {/* g-4 provides gutters i.e., spacing between cols */}
             
                <Col xs={12} md={10} lg={7} className="d-flex justify-content-center mt-0">
                    {/* xs=12 ensures it takes full width on smaller screens */}
                    <Row>
                        <BoothsLayout
                            booths = {booths}
                            ownedBooths={ownedBooths}
                            selectedBooths={selectedBooths}
                            deselectedBooths={deselectedBooths}
                            toggleBoothSelection={toggleBoothSelection}
                            action={action}
                        />
                        <BoothNote />
                            
                    </Row>
                    
                
                </Col>

                <Col xs={12} md={10} lg={3} className="d-flex row justify-content-center mx-5 ">
                    <ActionModal 
                    show={true}
                    booths={booths}
                    ownedBooths={ownedBooths}
                    actionType={action}
                    setAction={setAction}
                    selectedBooths={selectedBooths}
                    deselectedBooths={deselectedBooths}
                    handleSubmit={handleSubmit}
                    // Các props khác cần thiết cho modal
                    />
                    <BoothManager
                        selectedBooths={selectedBooths}
                        ownedBooths={ownedBooths}
                        setAction={setAction}
                        handleSubmit={handleSubmit}
                    />
                </Col>
            </Row>
        </Container>
        </div>
    )
}

export default Booth