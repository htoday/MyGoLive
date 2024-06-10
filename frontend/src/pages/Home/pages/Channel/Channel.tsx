import {useEffect, useState} from "react";
import Video from "./components/Video/Video.tsx";
import { GiftLayer} from "./components/GiftLayer/GiftLayer.tsx";
import {Communication} from "./components/Communication/Communication.tsx";
import {Room} from "../../../../api/room.ts";
import {baseData} from "../../../../data/BaseData.ts";
import {Message, MessageType} from "../../../../api/communication.ts";
import {getGiftByName, Gift} from "../../../../api/gift.ts";
const defaultMessage=[
    new Message(MessageType.TEXT,"你好啊","系2统"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
]
export function Channel(props:{
    room:Room|null
}){
    const [gifts,setGifts]=useState([] as Gift[])
    const [message,setMessage]=useState(defaultMessage)
    let webSocket: WebSocket | null=null
    const initializeWebSocket=()=>{
        const url=baseData.webSocketServer.getBaseUrl()+"/ws/"+room.roomId
        webSocket = new WebSocket(url)
        webSocket.onopen = function () {
            console.log("webSocket打开成功")
        }

        webSocket.onmessage = function (event) {
            const data=JSON.parse(event.data) as Message
            switch (data.msgType){
                case MessageType.TEXT.valueOf():{
                    setMessage((pre)=>[...pre,data])
                    break
                }
                case MessageType.GIFT.valueOf():{

                    const gift=new Gift(data.name,room.roomOwner,getGiftByName(data.content)!,"")
                    setGifts((pre)=>[...pre,gift])

                    break
                }
                default:{
                    //TODO 未知类型的MessageType
                }
            }
        }
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
            {<Video
                title={room.roomName}
                url={broadcastUrl}
                owner={room.roomOwner}
                />

            }
        </>
    )
}