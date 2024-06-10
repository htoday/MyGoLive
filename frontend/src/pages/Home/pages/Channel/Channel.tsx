import {useEffect, useState} from "react";
import Video from "./components/Video/Video.tsx";
import {Gift, GiftLayer} from "./components/GiftLayer/GiftLayer.tsx";
import {Communication} from "./components/Communication/Communication.tsx";
import {Room} from "../../../../api/room.ts";
import {baseData} from "../../../../data/BaseData.ts";

export function Channel(props:{
    room:Room|null
}){
    const [gifts,setGifts]=useState([] as Gift[])
    //const [pushAddress, setPushAddress] = useState("" as string)
    //const [channelKey, setChannelKey] = useState("" as string)
    let webSocket: WebSocket | null = new WebSocket(`ws://${document.location.host}/ws/${props.room!.roomId}`)
    webSocket.onopen = function () {
        console.log("webSocket打开成功")
    }
    let count=1
    if(props.room===null){
        alert("请先在主页选择一个房间加入!")
        return(
            <></>
        )
    }
    const room=props.room as Room
    /*const handleGetPushAddress = async () => {
        const url = baseData.roomApiServer.getBaseUrl() + "/getRoomPushAddress"
        await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(new GetRoomPushAddressRequest(room.roomId, localStorage.getItem("username") ?? ""))
        }).then(res => (res.json() as Promise<GetRoomPushAddressResponse>)).then(data => {
            setChannelKey(data.channelKey)
            setPushAddress(data.pushAddress)
        }).catch(() => {
            handleGetPushAddress()
        })
    }*/
    useEffect(() => {
        document.title="频道"

        setInterval(()=>{
            //console.log("gift")
            //setGifts(()=>[new Gift("sender","receiver",GiftId.COMMON, "普通礼物"+count)])
            count++
        },1000)
        return ()=>{
            document.title=""
        }
    }, []);
    const broadcastUrl=baseData.broadcastServer.getBaseUrl()+"/"+props.room.roomId+".flv"
    console.log(broadcastUrl)
    return (
        <>
            <Communication socket={webSocket}/>
            <GiftLayer gifts={gifts} clearGifts={()=>{
                setGifts([])
            }}></GiftLayer>
            {<Video
                title={room.roomName}
                url={broadcastUrl}
                owner={room.roomOwner}
                />

            }
        </>
    )
    //http://sf1-cdn-tos.huoshanstatic.com/obj/media-fe/xgplayer_doc_video/flv/xgplayer-demo-480p.flv
}