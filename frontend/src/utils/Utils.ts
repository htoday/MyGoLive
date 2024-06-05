export function oneRunningAsync(func: Function, onCancel:()=>void=()=>{}){
    let running = false;
    return async function (...parameters:any[]) {
        if (running) {
            onCancel();
            return;
        } else {
            running = true;
            //console.log(running)
            await func(...parameters)
            running = false;
            //console.log(running)
        }
    }
}
export function isPhone() {
    //获取浏览器navigator对象的userAgent属性（浏览器用于HTTP请求的用户代理头的值）
    const info = navigator.userAgent;
    //通过正则表达式的test方法判断是否包含“Mobile”字符串
    //如果包含“Mobile”（是手机设备）则返回true
    return /mobile/i.test(info);
}