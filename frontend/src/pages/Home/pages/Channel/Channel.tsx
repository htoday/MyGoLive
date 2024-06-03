import {useEffect, useState} from "react";
import Video from "./components/Video/Video.tsx";
import {Gift, GiftId, GiftLayer} from "./components/GiftLayer/GiftLayer.tsx";
import {Communication} from "./components/Communication/Communication.tsx";

export function Channel(){
    const [gifts,setGifts]=useState([] as Gift[])
    let count=1
    useEffect(() => {
        document.title="频道"

        setInterval(()=>{
            //console.log("gift")
            setGifts(()=>[new Gift("sender","receiver",GiftId.COMMON, "普通礼物"+count)])
            count++
        },1000)
        return ()=>{
            document.title=""
        }
    }, []);
    return (

        <>
            <Communication/>
            <GiftLayer gifts={gifts} clearGifts={()=>{
                setGifts([])
            }}></GiftLayer>
            {<Video title={"Welcome to Aqua's channel"} url={"http://sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/flv/xgplayer-demo-480p.flv"} owner={"Aqua"}/>}
        </>
    )
    //http://sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/flv/xgplayer-demo-480p.flv
}