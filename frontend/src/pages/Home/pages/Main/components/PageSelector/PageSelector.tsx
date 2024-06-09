import styles from "./PageSelector.module.less"
export function PageSelector(props:{
    pageIndex:number,
    onPreviousClick:()=>void,
    onNextClick:()=>void,
}){
    return (
        <div className={styles.content}>
            <button onClick={()=>{
                props.onPreviousClick()
            }}>{"<"}</button>
            <p>第 {props.pageIndex} 页</p>
            <button onClick={()=>{
                props.onNextClick()
            }}>{">"}</button>
        </div>
    )
}