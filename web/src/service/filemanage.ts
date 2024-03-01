import {ToastServiceMethods} from "primevue/toastservice";
import axios from "../axios/axios";
import {EmptyNamespaceUser, FileItem, LinkParams, Namespace, useFileStore} from "../store/file";
import {ConfirmationOptions} from "primevue/confirmationoptions";
import {useUserStore} from "../store/user";
import router from "../router/router";

export interface ConfirmMethods{
    require:(option: ConfirmationOptions)=>void;
    close:()=>void;
}

const fileStore=useFileStore();
const userStore=useUserStore();
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
            fileStore.dialog.createTextFile.visible=true
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
    axios.delete("/user/file/"+file.id).then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
        fileStore.fileList.splice(fileStore.fileList.indexOf(file),1);
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "删除文件成功",
            life: 3000
        });
    }).catch((e)=>{
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "删除文件失败",
            life: 3000
        });
    });
}
export function initLinkParams(){
    if(fileStore.linkParams.user_id!=0){
        return
    }
    const userStore=useUserStore();
    fileStore.linkParams=getUrlParams(window.location.href);
    userStore.linkNamespace=true;
    if(userStore.user.id==fileStore.linkParams.user_id){
        console.log("跳转")
        router.push("/login")
    }
}
export async function deleteDirectory(file:FileItem,toast: ToastServiceMethods):Promise<any> {
    const res=await axios.delete("/user/directory/"+file.id).then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
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
    fileStore.clickedFileItem=it;
    if(it.fileType=="directory"){
        fileStore.fileList=[];
        fileStore.breadcrumbs.push(it);
        let res=await axios.get("/user/file/"+it.id+"/list").then((res)=>{
            return res.data;
        }).catch((e)=>{
            console.log("error:",e);
        })
        fileStore.fileList=res.data.items;
    }else if(it.fileType=="png"||it.fileType=="jpg"||it.fileType=="jpeg"||it.fileType=="gif"||it.fileType=="bmp"){
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
    }else if(it.fileType=="pdf"){
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
                fileStore.dialog.pdfPreview.visible=true;
                fileStore.dialog.pdfPreview.url=resp.data;
            }
            return;
        }
        fileStore.dialog.pdfPreview.visible=true;
        fileStore.dialog.pdfPreview.url=url;
    }else if(it.fileType=="md"){
            const resp=await axios.get("/user/file/"+it.id+"/content").then((res)=>{
                return res.data;
            }).catch((e)=>{
                console.log("error:",e);
            })
            if(resp&&resp.code==0){
                fileStore.dialog.markdownPreview.visible=true;
                fileStore.dialog.markdownPreview.text=resp.data;
            }
            return;
    }else if(it.fileType=="txt"||it.fileType=="html"||it.fileType=="json"||it.fileType=="xml"||it.fileType=="js"||canEdit(it)){
            const resp=await axios.get("/user/file/"+it.id+"/content").then((res)=>{
                return res.data;
            }).catch((e)=>{
                console.log("error:",e);
            })
            if(resp&&resp.code==0){
                fileStore.dialog.txtPreview.visible=true;
                fileStore.dialog.txtPreview.text=resp.data;
            }
            return;
    }else{
        console.log("fileType error")
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
    const res=await axios.post("/user/directory/"+fileStore.getCurrentParentID(),{
        directoryName:folderName,
    }).then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: "创建文件夹失败:"+res.data.message,
                life: 3000
            });
            console.log("error:",res.data.message)
            return;
        }
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
    await axios.put("/user/file/"+file.id+"/rename",{
        newName:file.fileName,
    }).then((res)=>{
        if(res.data.code!=0){
            file.fileName=fileStore.rename.oldName;
            toast.add({
                severity: 'error',
                summary: '错误',
                detail:res.data.message,
                life: 3000
            });
            return;
        }
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "重命名成功",
            life: 3000
        });
        return res.data
    }).catch((e)=>{
        console.log("rename-error:",e)
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "重命名失败",
            life: 3000
        });
    });
    await unlockFile(file,toast);
    console.log("解锁")
}
export async function getUserNamespaces():Promise<Namespace[]> {
    const resp=await axios.get("/user/namespace/list").then((res)=>{
        console.log("namespace-list:",res);
        return res.data.data;
    }).catch((e)=>{
        if(e.response.data.code==401){
            router.push("/login");
        }
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
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail:res.data.message,
                life: 3000
            });
            return;
        }
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
export async function generateNamespaceJoinLink(it:Namespace,auth:number):Promise<any> {
    const resp=await axios.get("/user/namespace/"+it.id.toString()+"/link?name="+it.name+"&authority="+auth.toString()).then((res)=>{
        return res.data.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    return resp;
}
export function getUrlParams(url:string):LinkParams {
    let urlStr = url.split('?')[1]
    const urlSearchParams = new URLSearchParams(urlStr)
    const result = Object.fromEntries(urlSearchParams.entries())
    return {
        user_id:parseInt(result.user_id),
        namespace_id:parseInt(result.namespace_id),
        name:result.name,
        authority:parseInt(result.authority),
        token:result.token,
    }
}
export async function linkNamespace(id:number,auth:number,token:string,toast: ToastServiceMethods){
    const resp=await axios.post("/user/namespace/"+id.toString()+"/link?token="+token+"&authority="+auth.toString()).then((res)=>{
        return res.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    if(resp&&resp.code==0){
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "加入命名空间成功",
            life: 3000
        });
        await router.push("/")
    }else{
        console.log("resp:",resp);
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "加入命名空间失败",
            life: 3000
        })
    }
}
export function canEdit(file:FileItem):boolean{
    const stringArray: string[] = ["go","cpp","c","vue","md","txt","ts","js","json","html","css","yml"];
    return stringArray.includes(file.fileType);
}
export async function  getFileURL(id:string){
    let url=fileStore.urlMap.get(id);
    if(url==undefined){
        const resp:string=await axios.get("/user/file/"+id+"/url").then((res)=>{
            return res.data.data;
        }).catch((e)=>{
            console.log("error:",e);
        })
        url=resp
    }
    return url
}
export async function lockFile(file:FileItem,toast: ToastServiceMethods){
    const username=userStore.user.username;
    if(username=="未知"){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "请先登录",
            life: 3000
        });
        return false
    }
    const resp=await axios.get("/user/file/"+file.id+"/check?username="+userStore.user.username).then((res)=>{
        return res.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    console.log(resp)
    if(resp.data=="locked"){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail:resp.message,
            life: 3000
        });
        return false
    }
    return true
}
export async function unlockFile(file:FileItem,toast: ToastServiceMethods){
    const username=userStore.user.username;
    if(username=="未知"){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "请先登录",
            life: 3000
        });
        return false
    }
    const resp=await axios.get("/user/file/"+file.id+"/check?username="+userStore.user.username).then((res)=>{
        return res.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
    if(resp.data!="unlock"){
        console.log("resp:",resp);
        toast.add({
            severity: 'error',
            summary: '错误',
            detail:resp.message,
            life: 3000
        })
        return false
    }
    return true
}
export async function editFile(file:FileItem,toast: ToastServiceMethods){
    const locked=await lockFile(file,toast);
    if(!locked){
        return
    }
    fileStore.dialog.markdownEdit.visible=true
    fileStore.dialog.markdownEdit.editFileItem=file;
    await axios.get("/user/file/"+file.id+"/content").then(async (res) => {
        fileStore.dialog.markdownEdit.text = res.data.data;
        // await unlockFile(file, toast);
        console.log("res:", res.data.data)
    }).catch((e)=>{
        console.log("error:",e);
    })
}
export async function saveFileContent(toast: ToastServiceMethods){
    await axios.put("/user/file/"+fileStore.dialog.markdownEdit.editFileItem.id+"/content",{
        content:fileStore.dialog.markdownEdit.text,
    }).then(async (res) => {
        fileStore.dialog.markdownEdit.visible=false;
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "保存成功",
            life: 3000
        });
    }).catch((e)=>{
        console.log("error:",e);
        if(e.response.status==401){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: "请先登录",
                life: 3000
            });
        }
    });
}
export async function createTextFile(toast: ToastServiceMethods){
    const state=fileStore.dialog.createTextFile;
    if(state.fileName==""){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "文件名不能为空",
            life: 3000
        });
        return;
    }
    if(state.text==""){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "文件内容不能为空",
            life: 3000
        });
        return;
    }
    const parentID=fileStore.getCurrentParentID()
    await axios.post("/user/file/"+parentID+"/content",{
        fileName:state.fileName,
        content:fileStore.dialog.createTextFile.text,
    }).then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "保存成功",
            life: 3000
        });
        fileStore.fileList.push(res.data.data);
        fileStore.dialog.createTextFile.visible=false;
    }).catch((e)=>{
        console.log("error:",e);
        toast.add({
            severity: 'error',
        });
    });
}
export async function downloadFile(file:FileItem,toast: ToastServiceMethods){
    if(!canEdit(file)){
        //通过url下载
        const urlMap=fileStore.urlMap;
        let url=urlMap.get(file.id);
        if(url==undefined){
            const resp=await axios.get("/user/file/"+file.id+"/url").then((res)=>{
                return res.data;
            }).catch((e)=>{
                console.log("error:",e);
            })
           url=resp.data
        }
        if(!url){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: "文件不存在",
                life: 3000
            });
            return
        }
        var a = document.createElement('a');
        a.href = url;
        a.download = file.fileName;
        a.click();
        window.URL.revokeObjectURL(url); // 释放该 url
       return

    }
    const resp=await axios.get("/user/file/"+file.id+"/download").then((res)=>{
        var blob = new Blob([res.data]);
        var url = window.URL.createObjectURL(blob); // 创建 url 并指向 blob
        var a = document.createElement('a');
        a.href = url;
        a.download = file.fileName;
        a.click();
        window.URL.revokeObjectURL(url); // 释放该 url
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "下载成功",
            life: 3000
        });
        return res.data;
    }).catch(e=>{
        console.log("error:",e);
    })
}
export async function cancelAuth(toast: ToastServiceMethods){
    if(!fileStore.selectedUser){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "请选择要取消授权的用户",
            life: 3000
        });
        return
    }
    const namespaceID=fileStore.selectedNamespace.id;
    if(!namespaceID){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "请选择对应的命名空间",
            life: 3000
        });
        return
    }
    await axios.delete("/user/namespace/"+namespaceID+"/authority/"+fileStore.selectedUser.id).then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
        toast.add({
            severity: 'success',
            summary: '成功',
            detail: "取消成功",
            life: 3000
        });
        fileStore.selectedUser=EmptyNamespaceUser;

    })
    console.log("cancelAuth:",fileStore.selectedUser.id,namespaceID)
}
export async function fetchNamespaceUsers(toast: ToastServiceMethods){
    const id =fileStore.selectedNamespace.id
    if(!id){
        toast.add({
            severity: 'error',
            summary: '错误',
            detail: "请选择对应的命名空间",
            life: 3000
        });
        return
    }
    await axios.get("/user/namespace/"+id+"/users").then((res)=>{
        if(res.data.code!=0){
            toast.add({
                severity: 'error',
                summary: '错误',
                detail: res.data.message,
                life: 3000
            });
            return
        }
        fileStore.namespaceUsers=res.data.data;
    }).catch((e)=>{
        console.log("error:",e);
    });
}