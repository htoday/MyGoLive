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

    }
}