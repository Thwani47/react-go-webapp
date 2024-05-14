import './App.css';
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import CreateUser from './components/CreateUser';
import MainChat from './components/MainChat';
import Login from './components/Login';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/create-user" element={<CreateUser />} />
        <Route path='/chat' element={<MainChat />} />
        <Route path='/chat/:channelId' element={<MainChat />} />
        <Route path='/' element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;
