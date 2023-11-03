import React, { useState } from 'react';
import Profile from '../components/Profile';
import { Link } from 'react-router-dom';

// Dummy profile data   
const userProfileData = {
  name: 'John Doe',
  email: 'john@example.com',
};

const ProfilePage: React.FC = () => {
  const [profile, setProfile] = useState(userProfileData);

  const updateProfile = (name: string, email: string) => {
    setProfile({ ...profile, name, email });
  };

  return (
    <><div>
          <Profile
              name={profile.name}
              email={profile.email}
              onUpdate={updateProfile} />
      </div><div>
              <Link to="/">Back to Feed Page</Link>
          </div></>
  );
};

export default ProfilePage;
