export const baseData={
    apiServer:{
        hostname:"localhost",
        port:8889,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.port;
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
}