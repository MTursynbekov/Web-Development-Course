import React from 'react';

interface PostProps {
  id: string;
  content: string;
  likes: number; 
  dislikes: number; 
  likePost: () => void;
  dislikePost: () => void;
}

const Post: React.FC<PostProps> = ({ id, content, likes, dislikes, likePost, dislikePost }) => {
  return (
    <div>
      <div>{content}</div>
      <div>
        Likes: {likes}
        <button onClick={likePost}>Like</button>
      </div>
      <div>
        Dislikes: {dislikes}
        <button onClick={dislikePost}>Dislike</button>
      </div>
    </div>
  );
};

export default Post;
