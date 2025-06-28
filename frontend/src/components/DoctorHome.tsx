import { useAuth } from "../hooks/useAuth";

const DoctorHome = () => {

    // import user info
    const { user } = useAuth();

    // based on role redirect to receptionist home or doctor's home

    return ( 
        <>
            <div> Doctor's Home Page </div>
            <div> {user?.username} </div>
            <div> {user?.role} </div>
        </>
    );
}
 
export default DoctorHome;