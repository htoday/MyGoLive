import {useEffect} from "react";

export function Channel(){
    useEffect(() => {
        document.title="频道"
        return ()=>{
            document.title=""
        }
    }, []);
    return(

        <div>111</div>
    )
}