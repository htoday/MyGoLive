import styles from "./TopMenu.module.less"
export function TopMenu(props:{
    menuList:string[],
    onIndexChange:(newIndex:number)=>void,
    index:number,
}){
    const items=[]
    for(let i=0;i<props.menuList.length;i++){
        items.push(
            <li key={i} className={i===props.index?styles.selected:""} onClick={()=>props.onIndexChange(i)}>
                {props.menuList[i]}
            </li>
        )
    }
    return(
        <ul className={styles.top_menu}>
            {items}
        </ul>
    )
}