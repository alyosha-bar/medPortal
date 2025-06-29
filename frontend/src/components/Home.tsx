import { useAuth } from "../hooks/useAuth";

const Home = () => {

    // import user info
    const { user } = useAuth();

    return ( 
        <>
            <div> Home Page </div>
            <div> {user?.username} </div>
        </>
    );
}
 
export default Home;