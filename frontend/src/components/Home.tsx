import { Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

const Home = () => {
  const { user } = useAuth();

  let targetPath = "/login";

  if (user?.role === "receptionist") {
    targetPath = "/reception";
  } else if (user?.role === "doctor") {
    targetPath = "/doctors";
  }

  return (
    <div className="home-container">
      <h1 className="welcome-title">
        Welcome{user?.username ? `, ${user.username}` : ""}!
      </h1>
      <p className="welcome-message">
        {user
          ? "Go to your dashboard here: "
          : "Please log in to continue."}
        {user && (
          <Link className="dashboard-link" to={targetPath}>
            {targetPath.replace("/", "")}
          </Link>
        )}
      </p>
    </div>
  );
};

export default Home;
