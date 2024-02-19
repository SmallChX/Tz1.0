import axios from "axios";
import React from "react";
import { Table, Button, Container } from "react-bootstrap";

function BoothList({ booths, onEdit }) {

    function handleEditClick(booth) {
        onEdit(booth);
    }
    return (
        <Container style={{marginTop:"100px"}}>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Company Name</th>
                        <th>Price</th>
                        <th>Level</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {booths.map((booth) => (
                        <tr key={booth.ID}>
                            <td>{booth.ID}</td>
                            <td>{booth.company_info.name}</td>
                            <td>{booth.price}</td>
                            <td>{booth.level}</td>
                            <td>
                                <Button variant="primary" onClick={() => handleEditClick(booth)}>Edit</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
        </Container>
    )
}

export default BoothList;