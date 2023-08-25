import axios from "axios";
import {useQuery} from "react-query";

export const Demo = () => {
    const {data, isLoading} = useQuery("sermons", async () => {
        return axios.get("/api/v1/sermons")
    })
    return (
        <>
         <h1 className="text-3xl font-bold underline">Hello world!</h1>
        {isLoading ? <>Loading...</> : JSON.stringify(data)}
        </>
    )
}


