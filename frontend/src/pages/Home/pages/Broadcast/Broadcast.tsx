import styles from "./Broadcast.module.less"
import {useState} from "react";
import {oneRunningAsync} from "../../../../utils/Utils.ts";
import {baseData} from "../../../../data/BaseData.ts";
import {CreateRoomRequest, CreateRoomResponse} from "../../../../api/room.ts";
export function Broadcast(){
    const [roomName,setRoomName]=useState("")
    const [broadcastUrl,setBroadcastUrl]=useState("")
    const [channelKey,setChannelKey]=useState("")
    const handleOpenBroadcast = oneRunningAsync(async () => {
        const url = baseData.roomApiServer.getBaseUrl() + "/createRoom"
        return await fetch(url,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(new CreateRoomRequest(localStorage.getItem("username")!, roomName))
            }
        ).then((res) => {
            return res.json() as Promise<CreateRoomResponse>
        }).then((data)=>{
            if (data.status === 200) {
                alert("开播成功")
            } else {
                alert("开播失败")
            }
        }).catch((e)=>{
            alert("开播失败!:"+e.message)
        })
    }, () => {
        alert("请勿重复点击此按钮")
    })
    return (
        <div className={styles.content}>
            <h1>我要开播</h1>
            <input value={roomName} onChange={(e)=>{setRoomName(()=>e.target.value)}} placeholder={"直播间名称"}/>
            <input value={broadcastUrl} onChange={(e)=>{setBroadcastUrl(()=>e.target.value)}} placeholder={"OBS推流地址"}/>
            <input value={channelKey} onChange={(e)=>{setChannelKey(()=>e.target.value)}} placeholder={"推流码"}/>
            <button onClick={handleOpenBroadcast}>开启直播间</button>
        </div>
    )
}