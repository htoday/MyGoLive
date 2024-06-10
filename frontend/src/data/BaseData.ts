export const baseData={
    userApiServer:{
        hostname:"localhost",
        port:8887,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port+"/v1";
        }
    },
    talkApiServer:{
        hostname:"localhost",
        port:8888,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port+"/v1";
        },
        getGiftUrl:function (){
            return this.protocol+"://"+this.hostname+":"+this.port;
        }

    },
    roomApiServer:{
        hostname:"localhost",
        port:8889,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port+"/v1";
        }
    },
    rpcServer:{
        hostname:"localhost",
        port:8081,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port;
        }
    },
    broadcastServer:{
        hostname:"localhost",
        port:7001,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port+"/live";
        }
    },
    webSocketServer:{
        hostname:"localhost",
        port:8890,
        protocol:"ws",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port;
        }
    }
}