import { Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

const NavBar = () => {
  const { isAuthenticated, logout, user } = useAuth();

  return (
    <nav className="navbar">
      <h2 className="logo">MedPortal</h2>

      <div className="nav-links">
        {!isAuthenticated ? (
          <>
            <Link to="/login">Login</Link>
            <Link to="/signup">Signup</Link>
          </>
        ) : (
          <>
            <span className="user-info">Hello, {user?.username} ({user?.role})</span>
            <button onClick={logout}>Logout</button>
          </>
        )}
      </div>
    </nav>
  );
};

export default NavBar;
