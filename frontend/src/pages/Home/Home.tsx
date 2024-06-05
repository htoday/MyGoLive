import {BottomMenu} from "./components/BottomMenu/BottomMenu.tsx";
import {useState} from "react";
import styles from "./Home.module.less"
import {Channel} from "./pages/Channel/Channel.tsx";
import {Main} from "./pages/Main/Main.tsx";
import {Room} from "../../api/room.ts";
import {Broadcast} from "./pages/Broadcast/Broadcast.tsx";
export function Home() {
    const [index, setIndex] = useState(0)
    const [joinedRoom, setJoinedRoom]=useState(null as (Room|null))
    return (
        <>
            <div className={styles.background}>
                <div className={styles.page_area}>
                    {index===0&&<Main onChannelClick={(room)=>{
                        setJoinedRoom(()=>room)
                        setIndex(1)
                    }}/>}
                    {index===1&&<Channel room={joinedRoom}/>}
                    {index===3&&<Broadcast/>}
                </div>
            </div>
            <BottomMenu
                onIndexChange={(index) => {
                    if(index===1&&joinedRoom===null){
                        alert("请先在主页选择一个房间加入!")
                        return
                    }
                    setIndex(() => index)
                }}
                menuList={["主 页","频 道","我 的","开 播"]}
                index={index}
            />
        </>

    )
}