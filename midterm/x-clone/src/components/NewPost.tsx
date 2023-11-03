import React, { useState } from 'react';

interface NewPostProps {
  addPost: (content: string) => void;
}

const NewPost: React.FC<NewPostProps> = ({ addPost }) => {
  const [content, setContent] = useState('');

  const handleSubmit = () => {
    addPost(content);
    setContent('');
  };

  return (
    <div>
      <textarea value={content} onChange={(e) => setContent(e.target.value)}></textarea>
      <button onClick={handleSubmit}>Post</button>
    </div>
  );
};

export default NewPost;
