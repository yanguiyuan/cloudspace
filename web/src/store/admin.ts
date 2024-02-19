import {defineStore} from "pinia";
import {User, UserState} from "../service/user";
import {AdminState} from "../service/admin";

export const useAdminStore=defineStore({
    id:"admin",
    state:():AdminState=>({
<<<<<<< HEAD
       users:[],
        selectedUser:{
            id:0,
            username:"未知",
            password:"",
            gender:"未知",
            email:"未绑定邮箱",
            role:"user",
            createTime:"",
            updateTime:""
        },
        dialog:{
           resetPassword:{
               visible:false,
               newPassword:"",
               confirmPassword:""
           }
        }
=======
       users:[]
>>>>>>> 641946dc (test)
    }),
    actions:{
       setUsers(users:User[]){
           this.users=users
       },
    }
})