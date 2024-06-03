import styles from "./ChannelDisplay.module.less"
import {Room} from "../../../../../../api/room.ts";
import iconEye from "../../../../../../assets/icon/MdiEye.svg"
import iconHead from "../../../../../../assets/icon/MdiHead.svg"
export function ChannelDisplay(props:{
    baseInformation:Room,
    onClick:(room:Room)=>void
}){
    return (
        <div className={styles.content} onClick={()=>{props.onClick(props.baseInformation)}}>
            <h1 className={styles.title}>{props.baseInformation.roomName}</h1>
            <img src={props.baseInformation.displayImage} alt={"加载图片失败"}/>
            <div className={styles.row}>
                <img src={iconHead} alt={""}/>
                <p>房主:{props.baseInformation.roomOwner}</p>
            </div>
            <div className={styles.row}>
                <img src={iconEye} alt={""}/>
                <p>观看人数:{props.baseInformation.viewerNum}</p>
            </div>
        </div>
    )
}