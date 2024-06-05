import styles from "./Broadcast.module.less"
import {useState} from "react";
export function Broadcast(){
    const [roomName,setRoomName]=useState("")
    const [broadcastUrl,setBroadcastUrl]=useState("")
    return (
        <div className={styles.content}>
            <h1>我要开播</h1>
            <input value={roomName} onChange={(e)=>{setRoomName(()=>e.target.value)}} placeholder={"直播间名称"}/>
            <input value={broadcastUrl} onChange={(e)=>{setBroadcastUrl(()=>e.target.value)}} placeholder={"OBS推流地址"}/>
            <button>开启直播间</button>
        </div>
    )
}