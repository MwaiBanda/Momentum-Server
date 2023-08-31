import axios from "axios";
import {useQuery} from "react-query";
import {Link} from "react-router-dom";

export const Demo = () => {
    const {data, isLoading} = useQuery("sermons", async () => {
        return axios.get("/api/v1/sermons")
    })
    
    if (isLoading) {
        return <> Loading... </>
    }
    
    return (
        <>
        <Link to="dashboard/services">
         <h1 className="text-3xl font-bold underline">Hello world!</h1>
        </Link>
        {JSON.stringify(data)}
        </>
    )
}


