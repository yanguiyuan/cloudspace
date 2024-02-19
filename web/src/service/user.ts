import axios from "../axios/axios";
import {useUserStore} from "../store/user";
import {ToastServiceMethods} from "primevue/toastservice";
import {useFileStore} from "../store/file";
import router from "../router/router";
export interface User {
    id:number;
    username:string;
    password:string;
    gender:string;
    role:string;
    email:string;
    createTime:string;
    updateTime:string;
}
export interface UserState{
    dialog:{
        account:{
            visible:boolean;
        }
    },
    input:{
        email:{
            visible:boolean;
        }
        gender:{
            visible:boolean;
        }
<<<<<<< HEAD
        password:{
            oldPassword:string;
            newPassword:string;
            confirmPassword:string;
        }
=======
>>>>>>> 641946dc (test)
    },
    user:User,
    linkNamespace:boolean
}

const fileStore=useFileStore();
const userStore=useUserStore();

export const FetchUserInfo=async function ():Promise<User> {
    const res=await axios.get("/user/info").then((res)=>{
        if(res.data.code==401){
            router.push("/login");
        }
        return res.data.data;
    }).catch(err=>{
        console.log(err);
    });
    return res;
}
const checkEmail=(s:string):boolean=>{
    if(!s.includes("@")||!s.includes(".")){
        return false;
    }else{
        return true;
    }
}
export const modifyUserInfo=async function (toast: ToastServiceMethods):Promise<void> {
    if (!checkEmail(userStore.user.email)){
<<<<<<< HEAD
        toast.add({severity:'error', summary: '修改失败', detail: '邮箱格式不正确,请修改邮箱', life: 3000});
=======
        toast.add({severity:'error', summary: '修改失败', detail: '邮箱格式不正确', life: 3000});
>>>>>>> 641946dc (test)
        return
    }
    const res=await axios.put("/user/info",userStore.user).then((res)=>{
        console.log("r:",res)
        return res.data;
    }).catch(err=>{
        toast.add({severity:'error', summary: '修改失败', detail: '修改失败', life: 3000});
        console.log(err);
    });
    if(res.code===0){
        userStore.input.email.visible=false;
        userStore.input.gender.visible=false;
        toast.add({severity:'success', summary: '修改成功', detail: '修改成功', life: 3000});
    }else{
        console.log(res)
        toast.add({severity:'error', summary: '修改失败', detail: '修改失败', life: 3000});
    }
}
export const logout=async function ():Promise<void> {
    const res=await axios.post("/user/logout").then((res)=>{
        console.log("r:",res)
        userStore.$reset();
        fileStore.$reset();
        return res.data;
    }).catch(err=>{
        console.log(err);
    });
<<<<<<< HEAD
}
export const modifyPassword=async (toast:ToastServiceMethods)=>{
    if(userStore.input.password.newPassword!==userStore.input.password.confirmPassword){
        toast.add({severity:'error', summary: '修改失败', detail: '两次输入的密码不一致', life: 3000});
        return
    }
    await axios.put("/user/"+userStore.user.id+"/password",{
        oldPassword:userStore.input.password.oldPassword,
        newPassword:userStore.input.password.newPassword,
    }).then((res)=>{
        if(res.data.code===0){
            userStore.input.password.oldPassword="";
            userStore.input.password.newPassword="";
            userStore.input.password.confirmPassword="";
            toast.add({severity:'success', summary: '修改成功', detail: '修改成功', life: 3000});
        }else{
            toast.add({severity:'error', summary: '修改失败', detail: '修改失败', life: 3000});
        }
    }).catch(err=>{
        console.log(err);
    })
=======
>>>>>>> 641946dc (test)
}