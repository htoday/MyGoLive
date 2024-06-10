import styles from "./Main.module.less"
import {useEffect, useState} from "react";
import {GetRoomListRequest, JoinRoomRequest, Room} from "../../../../api/room.ts";
import {ChannelDisplay} from "./components/ChannelDisplay/ChannelDisplay.tsx";
import {baseData} from "../../../../data/BaseData.ts";
import {oneRunningAsync} from "../../../../utils/Utils.ts";
import {PageSelector} from "./components/PageSelector/PageSelector.tsx";

export function Main(props:{
    onChannelClick:(room:Room)=>void
}){
    const roomListDefault=[
         new Room(1,"test","test",1,""),
    ]as Room[]
    const handleLoadRoomListRequest=()=>{
        const url=baseData.roomApiServer.getBaseUrl()+"/getRoomList"
        fetch(url,{
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                "Authorization":localStorage.getItem("token")!
            },
            body:JSON.stringify(new GetRoomListRequest(pageIndex))
        }).then(res=>res.json()).then(res=>{
            if(res.status===200){
                if(res.roomList!==null) {
                    setChannelList(res.roomList)
                }else{
                    setChannelList([])
                }
            }
        }).catch(err=>{
            console.log(err)
        })
    }

    const handleJoinRoomRequest=oneRunningAsync(async (room:Room)=>{
        const url=baseData.roomApiServer.getBaseUrl()+"/joinRoom"
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
    //const pages=["推荐","游戏","动漫","电影"]
    const [pageIndex,setPageIndex]=useState(1)
    //const [pageName,setPageName]=useState(pages[0])
    const [channelList,setChannelList]=useState(roomListDefault as Room[])
    const channelListComponent=channelList.map((channel,index)=>{
        return <ChannelDisplay onClick={async () => {
            await handleJoinRoomRequest(channel)
            props.onChannelClick(channel)
        }} key={index} baseInformation={channel}/>
    })

    useEffect(() => {
        document.title="首页"
        handleLoadRoomListRequest()
        return ()=>{
            document.title=""
        }
    }, []);
    useEffect(() => {
        handleLoadRoomListRequest()
    }, [pageIndex]);
    return (
        <div className={styles.content}>
            {/*false && <Search/>*/}
            {/*false&&<TopMenu menuList={pages}
                      onIndexChange={(index) => {
                          setPageIndex(index)
                          setPageName(pages[index])
                      }}
                      index={pageIndex}/>*/}
            <div className={styles.row}>
                {channelListComponent}
            </div>
            <PageSelector pageIndex={pageIndex} onNextClick={()=>{}} onPreviousClick={()=>{}}/>
        </div>
    )
}