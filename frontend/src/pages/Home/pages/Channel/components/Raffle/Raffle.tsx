import styles from "./Raffle.module.less"
import React, {JSX} from "react";

export function Raffle(props:{
    visible:boolean,
    content:string,
    endTime:number,
    joined:boolean,
    winnerNumber:number,
}) {

    const [expand, setExpand] = React.useState(false)
    let component: JSX.Element
    const [left,setLeft]=React.useState(0)
    const [top,setTop]=React.useState(0)
    if(!props.visible){
        return<></>
    }
    const handleMouseDown=(e:React.MouseEvent)=>{
        let nowLeft=left
        let nowTop=top
        let nowMouseX=e.clientX
        let nowMouseY=e.clientY
        document.onmousemove=(e)=>{
            let deltaX=e.clientX-nowMouseX
            let deltaY=e.clientY-nowMouseY
            setLeft(nowLeft+deltaX)
            setTop(nowTop+deltaY)
        }
        document.onmouseup=()=>{
            document.onmousemove=null
            document.onmouseup=null
        }
    }
    if (expand) {
        component = <div style={{left:left,top:top}} onMouseDown={handleMouseDown} className={styles.expand_content} onDoubleClick={()=>{
            setExpand(false)
        }}>
            <h1>直播间抽奖</h1>
            <p>{props.content}</p>
            <button className={!props.joined? styles.join:styles.joined} disabled={props.joined}>{!props.joined?"加入":"已加入"}</button>
            <p>人数:{props.winnerNumber}</p>
            <p>截止时间:{props.endTime}</p>
        </div>
    }else{
        component=<div style={{left:left,top:top}} onMouseDown={handleMouseDown} className={styles.not_expand} onDoubleClick={()=>{
            setExpand(true)
        }}>
            <h1>直播间抽奖</h1>
        </div>
    }
    return component
}