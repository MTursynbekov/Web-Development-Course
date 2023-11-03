import React, { useState, useEffect } from 'react';
import { Link, useParams } from 'react-router-dom';
import Post from '../components/Post';
import { initialPosts } from './FeedPage';


const PostPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [post, setPost] = useState<{
    id: string; 
    content: string;
    likes: number;
    dislikes: number; 
} | null>(null);

  useEffect(() => {
    console.log("Post ID from URL:", id);
    const foundPost = initialPosts.find((post) => post.id === id);
    setPost(foundPost || null);
  }, [id]);

  const likePost = () => {
    if (post) {
        setPost({
          ...post,
          likes: post.likes + 1,
        });
      }
  };

  const dislikePost = () => {
    if (post) {
        setPost({
          ...post,
          dislikes: post.dislikes + 1,
        });
      }
  };

  if (!post) return <div>Post not found</div>;

  return (
    <><div>
      <Post
        id={post.id}
        content={post.content}
        likePost={likePost}
        dislikePost={dislikePost} likes={post.likes} dislikes={post.dislikes} />
    </div><div>
        <Link to="/">Back to Feed Page</Link>
      </div></>
  );
};

export default PostPage;
