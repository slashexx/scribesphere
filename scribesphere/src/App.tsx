import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home';
import BlogList from './components/BlogList';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blogs" element={<BlogList/>} />
      </Routes>
    </Router>
  );
}

export default App;