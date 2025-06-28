import { Route, Routes } from "react-router-dom"
import Login from "./components/Login"
import Signup from "./components/Signup"
import Home from "./components/Home"
import { ProtectedRoute } from "./components/ProtectedRoute"
import ReceptionHome from "./components/ReceptionHome"
import DoctorHome from "./components/DoctorHome"
import PatientDetails from "./components/PatientDetails"
import NavBar from "./components/Navbar"
import "./App.css"


function App() {

  return (
    <>
      <NavBar />
      <Routes>
        <Route path="/login" element={<Login />}></Route>
        <Route path="/signup" element={<Signup />}></Route>

        {/* protected routes */}
        <Route element={<ProtectedRoute/>}>
          <Route path="/" element={<Home />}></Route>
        </Route>

        <Route element={<ProtectedRoute requiredRole="receptionist"/>}>
          <Route path="/reception" element={<ReceptionHome />}></Route>
        </Route>

        <Route element={<ProtectedRoute requiredRole="doctor"/>}>
          <Route path="/doctors" element={<DoctorHome />}></Route>
        </Route>

        <Route element={<ProtectedRoute/>}>
          <Route path="/patient/:id" element={<PatientDetails />}></Route>
        </Route>



      </Routes>
    </>
  )
}

export default App
