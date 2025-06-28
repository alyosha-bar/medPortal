import { useAuth } from "../hooks/useAuth";

const Home = () => {

    // import user info
    const { user, isAuthenticated, logout } = useAuth();

    // based on role redirect to receptionist home or doctor's home

    return ( 
        <>
            <div> Home Page </div>
            <div> {user?.username} </div>
        </>
    );
}
 
export default Home;