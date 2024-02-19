import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Swal from 'sweetalert2';

const ProfilePage = () => {
  const userRole = localStorage.getItem('userRole');
  const userName = localStorage.getItem('userName');
  const [companyInfo, setCompanyInfo] = useState({});
  const [isEditing, setIsEditing] = useState(false);

  useEffect(() => {
    async function fetchData() {
      try {
        const response = await axios.get("/api/profile/company");
        if (response.status === 200) {
          setCompanyInfo(response.data.data);
        }
      } catch (error) {
        console.error("Failed to fetch data:", error);
      }
    }
    fetchData();
  }, [isEditing]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCompanyInfo(prevState => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleEditClick = () => {
    if (isEditing) {
      // Khi ở trạng thái chỉnh sửa và nhấn Save
      saveCompanyInfo();
    }
    setIsEditing(!isEditing);
  };

  const saveCompanyInfo = async () => {
    try {
      const response = await axios.put("/api/profile/company", {
        "represent_name": companyInfo.RepresentName,
        "represent_mail": companyInfo.RepresentMail,
        "represent_phone_number": companyInfo.RepresentPhoneNumber,
      });
      if (response.status === 200) {
        Swal.fire('Success!', 'Company information updated successfully.', 'success');
        localStorage.setItem("firstLogin", false);
      } else {
        Swal.fire('Error!', 'There was a problem updating your information.', 'error');
      }
    } catch (error) {
      console.error("Failed to save company info:", error);
      Swal.fire('Error!', 'There was a problem updating your information.', 'error');
    }
  };

  const renderProfile = () => {
    switch (userRole) {
      case 'admin':
        return (
          <div>
            <h2>Admin Profile</h2>
            <p>Name: {userName}</p>
          </div>
        );
      case 'company':
        return (
          <div>
            <h2>Company Profile</h2>
            <p>Name: {companyInfo.CompanyName}</p>
            {isEditing ? (
              <>
              <label htmlFor="RepresentName">Representative Name:</label>
                <input
                  type="text"
                  id="RepresentName"
                  name="RepresentName"
                  value={companyInfo.RepresentName || ''}
                  onChange={handleChange}
                />
                <label htmlFor="RepresentPhoneNumber">Representative Phone Number:</label>
                <input
                  type="text"
                  id="RepresentPhoneNumber"
                  name="RepresentPhoneNumber"
                  value={companyInfo.RepresentPhoneNumber || ''}
                  onChange={handleChange}
                />
                <label htmlFor="RepresentMail">Representative Email:</label>
                <input
                  type="email"
                  id="RepresentMail"
                  name="RepresentMail"
                  value={companyInfo.RepresentMail || ''}
                  onChange={handleChange}
                />
              </>
            ) : (
              <>
                <p>Representative Name: {companyInfo.represent_name}</p>
                <p>Representative Phone Number: {companyInfo.represent_phone_number}</p>
                <p>Representative Email: {companyInfo.represent_mail}</p>
              </>
            )}
            <button onClick={handleEditClick}>{isEditing ? 'Save' : 'Edit'}</button>
          </div>
        );
      default:
        return <p>No profile information available.</p>;
    }
  };

  return (
    <div>
      {renderProfile()}
    </div>
  );
};

export default ProfilePage;
