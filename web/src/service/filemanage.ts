import {ToastServiceMethods} from "primevue/toastservice";
import axios from "../axios/axios";
import {useFileStore} from "../store/file";
import {ConfirmationOptions} from "primevue/confirmationoptions";
import {useRouter} from "vue-router";
export interface FileManagementState{
    dialog:{
        upload:{
            visible:boolean;
        };
        createDirectory:{
            visible:boolean;
        };
        imagePreview:{
            visible:boolean;
            url:string;
        };
        namespace:{
            visible:boolean;
        }
    }
    namespaces:Namespace[];
    breadcrumbs:FileItem[];
    fileList:FileItem[];
    urlMap:Map<string,string>
}
export interface FileItem{
    id:string;
    fileName:string;
    fileType:string;
    createTime:string;
    updateTime:string;
}
export interface ConfirmMethods{
    require:(option: ConfirmationOptions)=>void;
    close:()=>void;
}
export interface Namespace{
    id:number;
    name:string;
    rootID:string;
    authority:number;
    updateTime:string;
}
export const DefaultNamespace:Namespace={
    id:0,
    name:"未知",
    rootID:"",
    authority:0,
    updateTime:""
}
export const DefaultFileItem:FileItem={
    id:"",
    fileName:"",
    fileType:"",
    createTime:"",
    updateTime:"",
}
const fileStore=useFileStore();
export const SideMenuOptionItems=[
    {
        tooltip:"文件上传",
        to:"/",
        icon:"upload",
        enableRouterLink:false,
        action:()=>{
            fileStore.dialog.upload.visible=true;
        }
    },
    {
        tooltip:"创建文件夹",
        to:"/",
        icon:"create-folder",
        enableRouterLink:false,
        action:()=>{
            fileStore.dialog.createDirectory.visible=true;
        }
    },
    {
        tooltip:"创建文件",
        to:"/",
        icon:"create-file",
        enableRouterLink:false,
        action:()=>{
        }
    },
    {
        tooltip:"命名空间管理",
        to:"/",
        icon:"namespace-management",
        enableRouterLink:false,
        action:()=>{
            fileStore.dialog.namespace.visible=true;
        }
    },
];

export async function deleteFile(file:FileItem,toast: ToastServiceMethods):Promise<any> {
    const res=axios.delete("/user/file/"+file.id+"/"+file.fileName,).then((res)=>{
        fileStore.fileList.splice(fileStore.fileList.indexOf(file),1);
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "删除文件成功",
            life: 3000
        });
        return res.data
    }).catch((e)=>{
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "删除文件失败",
            life: 3000
        });
    });
    return res;
}
export async function deleteDirectory(file:FileItem,toast: ToastServiceMethods):Promise<any> {
    const res=await axios.delete("/user/file/dir/"+file.id+"/"+file.fileName,).then((res)=>{
        fileStore.fileList.splice(fileStore.fileList.indexOf(file),1);
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "删除文件夹成功",
            life: 3000
        });
        return res.data
    }).catch((e)=>{
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "删除文件夹失败",
            life: 3000
        });
    });
    return res;
}
export async function deleteFileOrDirectory(file:FileItem,toast: ToastServiceMethods,confirm:ConfirmMethods):Promise<any> {
    confirm.require({
        message: `你确定要删除 ${file.fileName} 吗?`,
        header: '提示',
        icon: 'pi pi-exclamation-triangle',
        rejectClass: 'p-button-secondary p-button-outlined',
        rejectLabel: '取消',
        acceptLabel: '确认',
        accept: () => {
            if(file.fileType==="directory"){
                deleteDirectory(file,toast);
                return;
            }
            let res=deleteFile(file,toast);
            console.log("delete-result:",res);
        },
    });
}
export async function onClickBreadcrumbs(it:FileItem,key:number) {
    if(it.fileType=="directory"||it.fileType=="namespace") {
        const breadcrumbs=fileStore.breadcrumbs;
        fileStore.breadcrumbs.splice(key+1,breadcrumbs.length-key);
        let res=await axios.get("/user/file/"+it.id+"/list").then((res)=>{
            return res.data;
        }).catch((e)=>{
            console.log("error:",e);
        })
        fileStore.fileList=res.data.items;
    }
}
export  async function onClickFileItem (it:FileItem) {
    if(it.fileType=="directory"){
        fileStore.fileList=[];
        fileStore.breadcrumbs.push(it);
        let res=await axios.get("/user/file/"+it.id+"/list").then((res)=>{
            return res.data;
        }).catch((e)=>{
            console.log("error:",e);
        })
        fileStore.fileList=res.data.items;
    }else if(it.fileType=="png"){
        const urlMap=fileStore.urlMap;
        let url=urlMap.get(it.id);
        if(url==undefined){
            const resp=await axios.get("/user/file/"+it.id+"/url").then((res)=>{
                return res.data;
            }).catch((e)=>{
                console.log("error:",e);
            })
            if(resp&&resp.code==0){
                urlMap.set(it.id,resp.data);
                fileStore.dialog.imagePreview.visible=true;
                fileStore.dialog.imagePreview.url=resp.data;
            }
            return;
        }
        fileStore.dialog.imagePreview.visible=true;
        fileStore.dialog.imagePreview.url=url;
    }
}
export function getCurrentBreadcrumbsPath():string{
    let path="";
    for(let i=0;i<fileStore.breadcrumbs.length;i++){
        path+=fileStore.breadcrumbs[i].fileName+"/";
    }
    return path;
}
export async function createFolder(folderName:string,toast: ToastServiceMethods) {
    console.log("createFolder:",)
    const res=await axios.post("/user/directory",{
        parentID:fileStore.getCurrentParentID(),
        directoryName:folderName,
    }).then((res)=>{
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "创建文件夹成功",
            life: 3000
        });
        fileStore.dialog.createDirectory.visible=false;
        return res.data
    }).catch((e)=>{
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "创建文件夹失败",
            life: 3000
        });
    })
    fileStore.fileList.push(res.data);
}
export async function getRootFileItemID():Promise<string> {
    const router=useRouter();
    const id =await axios.get("/user/file/root").then((res) => {
        return res.data.id;
    }).catch((e) => {
        if(e.response&&e.response.status==401){
            router.push("/login");
        }
        console.log("error:",e);
    })
    console.log("id:",id);
    return id;
}
export async function getFileItemByID(id:string):Promise<FileItem> {
    const resp=await axios.get("/user/file/"+id).then((res)=>{
        return res.data.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    return resp;
}
export async function initDefaultFileItemList(id:string):Promise<void> {
    const resp=await axios.get("/user/file/"+id+"/list").then((res)=>{
        return res.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    if(!resp){
        return ;
    }
    fileStore.fileList=resp.data.items;
    for (const urlmap of resp.data.urlmaps) {
        fileStore.urlMap.set(urlmap.id,urlmap.url);
    }
}
export async function renameFileOrDirectory(file:FileItem,toast: ToastServiceMethods){
    await axios.put("/user/file/rename",{
        id:file.id,
        newName:file.fileName,
    }).then((res)=>{
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "重命名成功",
            life: 3000
        });
        return res.data
    }).catch((e)=>{
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "重命名失败",
            life: 3000
        });
    });
}
export async function getUserNamespaces():Promise<Namespace[]> {
    const resp=await axios.get("/user/namespace/list").then((res)=>{
        return res.data.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    return resp;
}
// id 命名空间的根文件项id
export async function  changeNamespace(id:string){
    const namespaceFileItem=await getFileItemByID(id)
    if(!namespaceFileItem){
        return;
    }
    fileStore.breadcrumbs.length=0;
    fileStore.breadcrumbs.push(namespaceFileItem);
    await initDefaultFileItemList(id);
}
export async function createNamespace(name:string,toast: ToastServiceMethods){
    await axios.post("/user/namespace/"+name).then((res)=>{
        console.log("namespace-create:",res);
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "创建命名空间成功",
            life: 3000
        });
    }).catch((e)=>{
        console.log("error:",e);
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "创建命名空间失败",
            life: 3000
        });
    });
}
export async function generateNamespaceJoinLink(id:number):Promise<any> {
    const resp=await axios.get("/user/namespace/"+id.toString()+"/link").then((res)=>{
        return res.data.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    return resp;
}
export function getUrlParams(url:string):any {
    let urlStr = url.split('?')[1]
    const urlSearchParams = new URLSearchParams(urlStr)
    const result = Object.fromEntries(urlSearchParams.entries())
    return result
}