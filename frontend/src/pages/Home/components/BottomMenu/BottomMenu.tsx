import styles from "./BottomMenu.module.less"
import React, {useEffect, useRef} from "react";
let lastIndex=0
export function BottomMenu(props:{
    menuList:string[],
    onIndexChange:(newIndex:number)=>void,
    index:number,

}){
    const followRef:React.RefObject<HTMLDivElement>=useRef() as React.RefObject<HTMLDivElement>
    useEffect(() => {
        followRef.current!.style.left=props.index/props.menuList.length*100+"%"
    }, []);
    useEffect(() => {
        if(lastIndex!==props.index){

            //console.log("changed to "+props.index+" from "+lastIndex)
            let lastPoint=lastIndex
            let targetPoint=props.index
            let total=20
            let count=0
            let nowPoint=lastPoint
            const move = () => {
                if(count>=total) {
                    followRef.current!.style.left=targetPoint/props.menuList.length*100+"%"
                    return
                }
                nowPoint=lastPoint/props.menuList.length+(targetPoint-lastPoint)/props.menuList.length/total*count
                followRef.current!.style.left=nowPoint*100+"%"
                count++
                requestAnimationFrame(move)
            };
            move()
            lastIndex=props.index

        }
    }, [props.index]);
    const list=[] as JSX.Element[]
    for(let i=0;i<props.menuList.length;i++){
        list.push(
            <li key={i} className={i===props.index?styles.selected:styles.unselected}
                onClick={()=>props.onIndexChange(i)}
            >{props.menuList[i]}</li>
        )
    }

    return(
        <div className={styles.bottom_center}>
            <ul className={styles.bottom_menu}>
                {list}
                <div ref={followRef} className={styles.follow} style={{left:0}}></div>
            </ul>
        </div>

    )
}