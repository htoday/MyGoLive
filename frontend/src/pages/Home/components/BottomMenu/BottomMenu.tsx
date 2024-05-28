import styles from "./BottomMenu.module.less"
export function BottomMenu(props:{
    menuList:string[],
    onIndexChange:(newIndex:number)=>void,
    index:number,

}){
    const list=[] as JSX.Element[]
    for(let i=0;i<props.menuList.length;i++){
        list.push(
            <li className={i===props.index?styles.selected:styles.unselected}
                onClick={()=>props.onIndexChange(i)}
            >{props.menuList[i]}</li>
        )
    }

    return(
        <div className={styles.bottom_center}>
            <ul className={styles.bottom_menu}>
                {list}
                <div className={styles.follow} style={{left:`${props.index*100/props.menuList.length}%`}}></div>
            </ul>
        </div>

    )
}