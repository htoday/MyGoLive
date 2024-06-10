import styles from "./Video.module.less"
import React, {useEffect, useRef} from "react";
import flv from 'flv.js'
export interface VideoProps{
    title:string,
    owner:string,
    url:string,
}
export default function Video(props:VideoProps){
    const videoRef=useRef() as React.MutableRefObject<HTMLVideoElement>;
    let flvPlayer:flv.Player|null

    useEffect(() => {
        if (flvPlayer) return
        flvPlayer = flv.createPlayer({
            type: 'flv',// 指定视频类型
            url: props.url, // 指定流链接
            isLive: true,// 开启直播
            hasAudio: false,  // 开启/关闭 声音
            cors: true,  // 开启跨域访问
            hasVideo: true,
            duration: 0,
        }, {
            autoCleanupSourceBuffer: true,//对SourceBuffer进行自动清理
            autoCleanupMaxBackwardDuration: 12,//    当向后缓冲区持续时间超过此值（以秒为单位）时，请对SourceBuffer进行自动清理
            autoCleanupMinBackwardDuration: 8,//指示进行自动清除时为反向缓冲区保留的持续时间（以秒为单位）。
            enableStashBuffer: false, //关闭IO隐藏缓冲区
            isLive: true,
            lazyLoad: true,
        });
        flvPlayer.attachMediaElement(videoRef.current);
        flvPlayer.load();
        flvPlayer.play();
        return () => {
            if (flvPlayer) {
                flvPlayer.pause();
                flvPlayer.unload();
                flvPlayer.detachMediaElement();
                flvPlayer.destroy();
                flvPlayer = null;
            }
        }
    }, []);
    return(
        <div className={styles.content}>
            <div className={styles.title_menu}>
                <h1 className={styles.title}>{props.title}</h1>
            </div>
            <video className={styles.video} ref={videoRef}/>
            <div className={styles.bottom_menu}></div>

        </div>
    )
}