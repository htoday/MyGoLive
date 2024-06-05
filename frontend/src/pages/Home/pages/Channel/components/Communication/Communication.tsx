import styles from "./Communication.module.less"
enum MessageType{
    TEXT,
    GIFT,
    IMAGE,
    SYSTEM
}
export class Message{
    constructor(
        public type:MessageType,
        public content:string,
        public sender:string,
    ) {
    }
}

export function Communication(props:{socket:WebSocket|null}){
    //const [message,setMessage]=useState([] as Message[])
    if(props.socket===null){
        alert("socket异常 无法启动对话栏!")
        return <></>
    }else{
        const onMessage=props.socket.onmessage!
        props.socket.onmessage=function(event){
            onMessage.apply(props.socket!,[event])
            //const data=JSON.parse(event.data)
        }
    }
    return(
        <div className={styles.communication_window}>
            <h1 className={styles.title}>
                聊天信息
            </h1>
            <div className={styles.communication_window_communication}>
            </div>
            <div className={styles.communication_window_input}>
                <textarea className={styles.communication_window_input_text}></textarea>
                <button className={styles.communication_window_input_button}>发送</button>
            </div>
        </div>
    )
}