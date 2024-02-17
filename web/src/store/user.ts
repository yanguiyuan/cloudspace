import {defineStore} from "pinia";
import { User, UserState} from "../service/user";
export const DefaultUserState:UserState={
    dialog:{
        account:{
            visible:false,
        }
    },
    input:{
        email:{
            visible:false,
        },
        gender:{
            visible:false,
        }
    },
    user:{
        id:0,
        username:"未知",
        password:"",
        gender:"未知",
        email:"未绑定邮箱",
        role:"user",
        createTime:"",
        updateTime:""
    },
    linkNamespace:false
}
export const useUserStore=defineStore({
    id:"user",
    state:():UserState=>(DefaultUserState),
    actions:{
        showAccountDialog(){
            this.dialog.account.visible=true
        },
        setUser(u:User){
            if(u.email==""){
                u.email="未绑定邮箱";
            }
            this.user=u;
        }
    }
})