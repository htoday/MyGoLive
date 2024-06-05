import styles from "./Main.module.less"
import {TopMenu} from "./components/TopMenu/TopMenu.tsx";
import {useEffect, useState} from "react";
import {Search} from "./components/Search/Search.tsx";
import {GetRoomListRequest, JoinRoomRequest, Room} from "../../../../api/room.ts";
import {ChannelDisplay} from "./components/ChannelDisplay/ChannelDisplay.tsx";
import {baseData} from "../../../../data/BaseData.ts";
import {oneRunningAsync} from "../../../../utils/Utils.ts";

export function Main(props:{
    onChannelClick:(room:Room)=>void
}){
    const roomListDefault=[
        new Room(1,"测试房间","测试用户",100,""),
        new Room(2,"测试房间","测试用户",100,""),
        new Room(3,"测试房间","测试用户",100,""),
        new Room(4,"测试房间","测试用户",100,""),
        new Room(5,"测试房间","测试用户",100,""),
        new Room(6,"测试房间","测试用户",100,""),
        new Room(7,"测试房间","测试用户",100,""),
        new Room(8,"测试房间","测试用户",100,""),
        new Room(9,"测试房间","测试用户",100,""),
        new Room(10,"测试房间","测试用户",100,""),
    ]
    const handleLoadRoomList=()=>{
        const url=baseData.userApiServer.getBaseUrl()+"/getRoomList"
        fetch(url,{
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify(new GetRoomListRequest(pageName))
        }).then(res=>res.json()).then(res=>{
            if(res.status===200){
                setChannelList(res.roomList)
            }
        }).catch(err=>{
            console.log(err)
        })
    }
    const handleJoinRoomRequest=oneRunningAsync(async (room:Room)=>{
        const url=baseData.userApiServer.getBaseUrl()+"/joinRoom"
        await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": localStorage.getItem("token")!
            },
            body: JSON.stringify(new JoinRoomRequest(localStorage.getItem("username")!,room.roomId))
        }).then( r =>{
            if(r.status===200){
                return r.json()
            }else{
                throw new Error("加入房间失败")
            }
        }).then(()=>{
            setPageIndex(1)
            props.onChannelClick(room)
        }).catch(err=>{
            alert("加入房间失败:"+err.message)
        })
    },()=>{
        alert("正在加入房间 请勿重复点击")
    })
    const pages=["推荐","游戏","动漫","电影"]

    const [pageIndex,setPageIndex]=useState(0)
    const [pageName,setPageName]=useState(pages[0])
    const [channelList,setChannelList]=useState(roomListDefault as Room[])
    const channelListComponent=channelList.map((channel,index)=>{
        return <ChannelDisplay onClick={async () => {
            await handleJoinRoomRequest(channel)
            props.onChannelClick(channel)
        }} key={index} baseInformation={channel}/>
    })

    useEffect(() => {
        document.title="首页"
        handleLoadRoomList()
        return ()=>{
            document.title=""
        }
    }, []);
    useEffect(() => {
        handleLoadRoomList()
       // console.log("page name changes")
    }, [pageName]);
    return (
        <div className={styles.content}>
            <Search/>
            <TopMenu menuList={pages}
                     onIndexChange={(index)=>{
                         setPageIndex(index)
                         setPageName(pages[index])
                     }}
                     index={pageIndex}/>
            <div className={styles.row}>
                {channelListComponent}
            </div>
        </div>
    )
}