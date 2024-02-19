import axios from "../axios/axios";
import {User} from "./user";
import {useAdminStore} from "../store/admin";
import router from "../router/router";
import {ToastServiceMethods} from "primevue/toastservice";
export interface AdminState {
    users:User[];
    selectedUser:User;
    dialog:{
        resetPassword:{
            visible:boolean
            newPassword:string
            confirmPassword:string
        }
    }
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
export const resetPassword=async function (toast:ToastServiceMethods):Promise<void> {
    if(adminStore.dialog.resetPassword.confirmPassword!==adminStore.dialog.resetPassword.newPassword){
        toast.add({severity:'error',summary:'重置密码失败',detail:'两次输入的密码不一致',life:3000});
        return;
    }
    const res=await axios.put("/admin/user/"+adminStore.selectedUser.id+"/password",{
        newPassword:adminStore.dialog.resetPassword.newPassword
    }).then((res)=>{
        console.log("users:",res.data)
        toast.add({severity:'success',summary:'重置密码成功',detail:'修改密码成功',life:3000});
        adminStore.dialog.resetPassword.visible=false;
        return res.data.data;
    }).catch(err=>{
        if (err.response.status===401){
            router.push("/login");
            return;
        }
        toast.add({severity:'error',summary:'重置密码失败',detail:'修改密码失败',life:3000});
        console.log(err);
    });

}