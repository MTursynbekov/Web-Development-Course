import React, { useState } from 'react';

interface ProfileProps {
  name: string;
  email: string;
  onUpdate: (name: string, email: string) => void;
}

const Profile: React.FC<ProfileProps> = ({ name, email, onUpdate }) => {
  const [editable, setEditable] = useState(false);
  
  const [profileName, setProfileName] = useState(name);
  const [profileEmail, setProfileEmail] = useState(email);

  const handleUpdate = () => {
    onUpdate(name, email);
    setEditable(false);
  };

  return (
    <div>
      {editable ? (
        <>
          <input value={profileName} onChange={(e) => setProfileName(e.target.value)} />
          <input value={profileEmail} onChange={(e) => setProfileEmail(e.target.value)} />
          <button onClick={handleUpdate}>Update</button>
        </>
      ) : (
        <>
          <p>{profileName}</p>
          <p>{profileEmail}</p>
          <button onClick={() => setEditable(true)}>Edit</button>
        </>
      )}
    </div>
  );
};

export default Profile;
