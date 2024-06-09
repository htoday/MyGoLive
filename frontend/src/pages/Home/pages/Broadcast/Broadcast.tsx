import styles from "./Broadcast.module.less"
import {useEffect, useState} from "react";
import {oneRunningAsync} from "../../../../utils/Utils.ts";
import {baseData} from "../../../../data/BaseData.ts";
import {
    CreateRoomRequest,
    CreateRoomResponse,
    GetRoomPushAddressRequest,
    GetRoomPushAddressResponse
} from "../../../../api/room.ts";
export function Broadcast() {
    const [roomName, setRoomName] = useState("")
    const [broadcastUrl, setBroadcastUrl] = useState("")
    const [channelKey, setChannelKey] = useState("")
    const isBroadcasting = broadcastUrl !== "" && channelKey !== ""
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
                body: JSON.stringify(new GetRoomPushAddressRequest(0, localStorage.getItem("username")!))
            }).then((res) => {
                if(res.status !== 200){
                    alert("获取房间状态失败:" + res.status)
                    return
                }
                return res.json() as Promise<GetRoomPushAddressResponse>
        }).then(
            (data) => {
                if (data!.status === 200) {
                    setBroadcastUrl(data!.pushAddress)
                    setChannelKey(data!.channelKey)
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
            <input value={roomName} disabled={isBroadcasting} onChange={(e) => {
                setRoomName(() => e.target.value)
            }} placeholder={"直播间名称"}/>
            <p onClick={async () => {
                if(!isBroadcasting) {
                    alert("请先开播 再进行复制")
                    return
                }
                editClipboard(broadcastUrl)}}>
                {channelKey === "" ? "开启房间后 即可获取直播间推流地址" : "推流地址:" + channelKey + "(点击复制到剪贴板)"}
            </p>
            <p onClick={async () => {
                if(!isBroadcasting) {
                    alert("请先开播 再进行复制")
                    return
                }
                editClipboard(channelKey)}}>
                {channelKey === "" ? "开启房间后 即可获取直播间推流码" : "推流码:" + channelKey + "(点击复制到剪贴板)"}
            </p>
            <button onClick={handleOpenBroadcast}>{isBroadcasting ? "关闭直播间" : "开启直播间"}</button>
        </div>
    )
}