import { useEffect } from "react"
import { Nav } from "../components/nav"
import { useNavigate } from "react-router-dom";
  
export const Home = () => {
    const navigate = useNavigate();

    useEffect(() => {
        navigate('/messages')
    }, [])
    return (
        <div className="w-full">
        <Nav/>
        </div>
    )
}






