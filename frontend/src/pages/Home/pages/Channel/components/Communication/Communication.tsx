import styles from "./Communication.module.less"
import {useState} from "react";
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
const defaultMessage=[
    new Message(MessageType.TEXT,"你好啊","系2统"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
    new Message(MessageType.TEXT,"AAA","我"),
]
export function Communication(props:{socket:WebSocket|null}){
    const [message,setMessage]=useState(defaultMessage as Message[])
    if(props.socket===null){
        alert("socket异常 无法启动对话栏!")
        return <></>
    }else{
        const onMessage=props.socket.onmessage!
        props.socket.onmessage=function(event){
            onMessage.apply(props.socket!,[event])
            const data=JSON.parse(event.data)
            switch (data.type){
                case MessageType.TEXT: {
                    setMessage([...message, new Message(MessageType.TEXT, data.content, data.sender)])
                    break;
                }
                default:{
                    setMessage([...message, new Message(MessageType.TEXT, "未知类型消息", data.sender)])
                    break;
                }
            }
        }
    }
    const messageList=message.map((msg,index)=>{
        return (<div key={index} className={styles.communication_window_communication_content}>
            <h1>{msg.sender}</h1>
            <p>{msg.content}</p>

        </div>)
    })
    return(
        <div className={styles.communication_window}>
            <h1 className={styles.title}>
                聊天信息
            </h1>
            <div className={styles.communication_window_communication}>
                {messageList}
            </div>
            <div className={styles.communication_window_input}>
                <textarea className={styles.communication_window_input_text}>
                </textarea>
                <button className={styles.communication_window_input_button}>发送</button>
            </div>
        </div>
    )
}