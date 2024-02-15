import {defineStore} from "pinia";
import {User, UserState} from "../service/user";
import {AdminState} from "../service/admin";

export const useAdminStore=defineStore({
    id:"admin",
    state:():AdminState=>({
       users:[]
    }),
    actions:{
       setUsers(users:User[]){
           this.users=users
       },
    }
})