import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import FeedPage from './pages/FeedPage';
import PostPage from './pages/PostPage';
import ProfilePage from './pages/ProfilePage';


const App: React.FC = () => {
  return (
    <Router>
      <Routes>
          <Route index element={<FeedPage />} />
          <Route path="/post/:id" element={<PostPage />} />
          <Route path="/profile" element={<ProfilePage />} />
      </Routes>
    </Router>
    );
};

export default App;
