import axios from "../axios/axios";
import {User} from "./user";
import {useAdminStore} from "../store/admin";
import router from "../router/router";
export interface AdminState {
    users:User[];
}
const adminStore=useAdminStore();
export const getUsers=async function (offset:number,limit:number):Promise<void> {
    const res=await axios.get("/user/list/"+offset+"/"+limit).then((res)=>{
        console.log("users:",res.data)
        return res.data.data;
    }).catch(err=>{
        if (err.response.status===401){
            router.push("/login");
        }
        console.log(err);
    });
    adminStore.setUsers(res);
}