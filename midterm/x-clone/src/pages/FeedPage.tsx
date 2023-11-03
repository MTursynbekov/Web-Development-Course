import React, { useState } from 'react';
import NewPost from '../components/NewPost';
import Post from '../components/Post';
import { Link } from 'react-router-dom';

// Dummy data
export const initialPosts = [
  { id: '1', content: 'This is a post 1', likes: 5, dislikes: 0 },
  { id: '2', content: 'This is a post 2' , likes: 10, dislikes: 3 },
  { id: '3', content: 'This is a post 3' , likes: 0, dislikes: 4 }
];

const FeedPage: React.FC = () => {
  const [posts, setPosts] = useState(initialPosts);

  const addPost = (content: string) => {
    const newPost = { id: (Math.random() * 1000).toString(), content, likes: 0, dislikes: 0 };
    setPosts([newPost, ...posts]);
  };

  const likePost = (postId: string) => {
    setPosts(posts.map(post => {
      if (post.id === postId) {
        return { ...post, likes: post.likes + 1 };
      }
      return post;
    }));
  };

  const dislikePost = (postId: string) => {
    setPosts(posts.map(post => {
      if (post.id === postId) {
        return { ...post, dislikes: post.dislikes + 1 };
      }
      return post;
    }));
  };

  return (
    <><div>
      <NewPost addPost={addPost} />
      {posts.map((post) => (
        <><Post key={post.id} {...post} likePost={() => likePost(post.id)} dislikePost={() => dislikePost(post.id)} />
        <Link to={`/post/${post.id}`}>Read more</Link></>
      ))}
    </div><div>
        <Link to="/profile">Go to Profile</Link>
      </div></>
  );
};

export default FeedPage;
