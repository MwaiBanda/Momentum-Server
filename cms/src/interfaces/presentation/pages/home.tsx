import { useEffect } from "react"
import { MomentumNavigation } from "../components/navigation"
import { useNavigate } from "react-router-dom";
  
export const Home = () => {
    const navigate = useNavigate();

    useEffect(() => {
        navigate('/messages')
    }, [])
    return (
        <div className="w-full">
        <MomentumNavigation />
        </div>
    )
}






