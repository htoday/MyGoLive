import styles from "./Broadcast.module.less"
import {useEffect, useState} from "react";
import {oneRunningAsync} from "../../../../utils/Utils.ts";
import {baseData} from "../../../../data/BaseData.ts";
import {
    CloseRoomRequest,
    CreateRoomRequest,
    CreateRoomResponse,
    GetRoomPushAddressRequest,
    GetRoomPushAddressResponse
} from "../../../../api/room.ts";
export function Broadcast() {
    const [roomName, setRoomName] = useState("")
    const [broadcastUrl, setBroadcastUrl] = useState("")
    const [channelKey, setChannelKey] = useState("")
    const [broadcasting,setBroadcasting] = useState(false)

    const editClipboard = (text: string) => {
        navigator.clipboard.writeText(text)
            .then(() => {
                alert("已复制到剪贴板")
            })
            .catch((e) => {
                alert("复制失败:" + e.message)
            })
    }
    const handleGetRoomStateRequest = async () => {
        const url = baseData.roomApiServer.getBaseUrl() + "/getRoomPushAddress"
        return await fetch(url,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                     "Authorization": localStorage.getItem("token")!
                },
                body: JSON.stringify(new GetRoomPushAddressRequest(localStorage.getItem("username")!))
            }).then((res) => {
                if(res.status !== 200){
                    alert("获取房间状态失败:" + res.status)
                    return
                }
                return res.json() as Promise<GetRoomPushAddressResponse>
        }).then(
            (data) => {
                if (data!.status === 200) {
                    if(data!.roomId!==-1){
                        setBroadcasting(true)
                        setBroadcastUrl(data!.pushAddress)
                        setChannelKey(data!.channelKey)
                        setRoomName(data!.roomName)
                    }else{
                        setBroadcasting(false)
                        setBroadcastUrl("")
                        setChannelKey("")
                    }

                } else {
                    //alert("获取房间状态失败")
                    alert("获取房间状态失败 请刷新网页重试!")
                    //handleGetRoomStateRequest()
                }
            }
        ).catch(()=>{
            alert("获取房间状态失败 请刷新网页重试!")
            //handleGetRoomStateRequest()
        })
    }
    const handleCloseBroadcast = oneRunningAsync(async () => {
        const url = baseData.roomApiServer.getBaseUrl() + "/closeRoom"
        fetch(url,{
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": localStorage.getItem("token")!
            },
            body: JSON.stringify(new CloseRoomRequest(localStorage.getItem("username")!))
        }).then(res=>{
            if(res.status === 200){
                return res.json() as Promise<CreateRoomResponse>
            }else{
                throw new Error("关闭房间失败")
            }
        }).then((data)=>{
            if(data.status!==200){
                throw new Error("关闭房间失败")
            }else{
                alert("关闭房间成功!")
            }
            handleGetRoomStateRequest()
        }).catch((e)=>{
            alert("关闭房间失败:" + e.message)
        })

    }, () => {
        alert("请勿重复点击此按钮")
    })
    const handleOpenBroadcast = oneRunningAsync(async () => {
        const url = baseData.roomApiServer.getBaseUrl() + "/createRoom"
        return await fetch(url,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": localStorage.getItem("token")!
                },
                body: JSON.stringify(new CreateRoomRequest(localStorage.getItem("username")!, roomName))
            }
        ).then((res) => {
            console.log(res.body)
            return res.json() as Promise<CreateRoomResponse>
        }).then((data) => {
            if (data.status === 200) {
                alert("开播成功")
            } else {
                alert("开播失败")
            }
            handleGetRoomStateRequest()
        }).catch((e) => {
            alert("开播失败!:" + e.message)
        })
    }, () => {
        alert("请勿重复点击此按钮")
    })
    useEffect(() => {
        handleGetRoomStateRequest()
    }, []);
    return (
        <div className={styles.content}>
            <h1>我要开播</h1>
            <input value={roomName} disabled={broadcasting} onChange={(e) => {
                setRoomName(() => e.target.value)
            }} placeholder={"直播间名称"}/>
            <p onClick={async () => {
                if(!broadcasting) {
                    alert("请先开播 再进行复制")
                    return
                }
                editClipboard(broadcastUrl)}}>
                {!broadcasting ? "开启房间后 即可获取直播间推流地址" :"(点击复制到剪贴板)"+ "推流地址:" + broadcastUrl }
            </p>
            <p onClick={async () => {
                if(!broadcasting) {
                    alert("请先开播 再进行复制")
                    return
                }
                editClipboard(channelKey)}}>
                {!broadcasting ? "开启房间后 即可获取直播间推流码" :"(点击复制到剪贴板)"+ "推流码:" + channelKey  }
            </p>
            <button onClick={broadcasting?handleCloseBroadcast:handleOpenBroadcast}>{broadcasting ? "关闭直播间" : "开启直播间"}</button>
        </div>
    )
}