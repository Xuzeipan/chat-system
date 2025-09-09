import { BrowserRouter, Routes, Route } from 'react-router'
import './App.css'
import AppLayout from './components/common/Layout/appLayout'
import Launch from './components/pages/LaunchPage/launch'
import Login from './components/pages/Auth/login'
import Register from './components/pages/Auth/register'

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route element={<AppLayout />}>
          <Route path='/' element={<Launch />} />
          <Route path='/login' element={<Login />} />
          <Route path='/register' element={<Register />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
