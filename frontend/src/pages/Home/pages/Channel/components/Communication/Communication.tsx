import styles from "./Communication.module.less"
import {Message, MessageType} from "../../../../../../api/communication.ts";
import {useState} from "react";


export function Communication(props:{messages:Message[],websocket:WebSocket|null}){

    const messageList=props.messages.map((msg,index)=>{
        return (<div key={index} className={styles.communication_window_communication_content}>
            <h1>{msg.name}</h1>
            <p>{msg.content}</p>

        </div>)
    })
    const [inputContent,setInputContent]=useState("")
    const handleSendMessage=()=>{
        if(props.websocket){
            const message=new Message(MessageType.TEXT,inputContent,localStorage.getItem("username")!)
            props.websocket.send(JSON.stringify(message))
            setInputContent("")
        }else{
            alert("websocket 未加载 请刷新网页重试!")
        }
    }
    return(
        <div className={styles.communication_window}>
            <h1 className={styles.title}>
                聊天信息
            </h1>
            <div className={styles.communication_window_communication}>
                {messageList}
            </div>
            <div className={styles.communication_window_input}>
                <textarea
                    className={styles.communication_window_input_text}
                    onChange={(e)=>{
                        setInputContent(()=>e.target.value)
                    }}
                    value={inputContent}>
                </textarea>
                <button className={styles.communication_window_input_button} onClick={handleSendMessage}>发送</button>
            </div>
        </div>
    )
}