import {useEffect, useState} from "react";
import Video from "./components/Video/Video.tsx";
import { GiftLayer} from "./components/GiftLayer/GiftLayer.tsx";
import {Communication} from "./components/Communication/Communication.tsx";
import {Room} from "../../../../api/room.ts";
import {baseData} from "../../../../data/BaseData.ts";
import {Message, MessageType} from "../../../../api/communication.ts";
import {getGiftByName, Gift} from "../../../../api/gift.ts";
import {GiftDisplay} from "./components/GiftDisplay/GiftDisplay.tsx";
import {Raffle} from "./components/Raffle/Raffle.tsx";
const defaultMessage:Message[]=[]
export function Channel(props:{
    room:Room|null
}){
    const [gifts,setGifts]=useState([] as Gift[])
    const [message,setMessage]=useState(defaultMessage)
    const [webSocket,setWebSocket]=useState(null as WebSocket|null)
    const initializeWebSocket=()=>{
        const url=baseData.webSocketServer.getBaseUrl()+"/ws/"+room.roomId
        const socket=new WebSocket(url)
        console.log("webSocket连接地址:"+url)
        socket.onopen = function () {
            console.log("webSocket打开成功")
        }
        socket.onmessage = function (event) {
            const data=JSON.parse(event.data) as Message
            switch (data.msgType){
                case MessageType.TEXT.valueOf():{
                    setMessage((pre)=>[...pre,data])
                    break
                }
                case MessageType.GIFT.valueOf():{

                    const gift=new Gift(data.name,room.roomOwner,getGiftByName(data.content)!,"")
                    setGifts(()=>[gift])

                    break
                }
                default:{
                    //TODO 未知类型的MessageType
                }
            }
        }
        setWebSocket(socket)
    }
    if(props.room===null){
        alert("请先在主页选择一个房间加入!")
        return(
            <></>
        )
    }
    const room=props.room as Room
    useEffect(() => {
        document.title="频道"
        initializeWebSocket()
        return ()=>{
            document.title=""
        }
    }, []);
    const broadcastUrl=baseData.broadcastServer.getBaseUrl()+"/"+props.room.roomId+".flv"
    console.log(broadcastUrl)
    return (
        <>
            <Communication messages={message} websocket={webSocket}/>
            <GiftLayer gifts={gifts} clearGifts={()=>{
                setGifts([])
            }}></GiftLayer>
            <GiftDisplay room={room}/>
            <Raffle visible={true} content={"奖品"} endTime={0} joined={true} winnerNumber={2}/>
            <Video title={room.roomName} url={broadcastUrl} owner={room.roomOwner}/>
        </>
    )
}