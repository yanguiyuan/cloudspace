import {defineStore} from "pinia";
const EmptyFileItem={id:"",fileName:"undefined",fileType:"",createTime:"",updateTime:""}
export const useFileStore=defineStore({
    id:"file",
    state:():FileManagementState=>({
       dialog:{
           upload:{
               visible:false,
           },
           createDirectory:{
               visible:false,
           },
           imagePreview:{
               visible:false,
               url:"",
           },
           namespace:{
               visible:false
           },
           markdownEdit:{
               visible:false,
               text:"",
               editFileItem:EmptyFileItem,
           }
       },
        linkParams:DefaultLinkParams,
        namespaces:[],
        breadcrumbs:[],
        fileList:[EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem],
        urlMap:new Map<string, string>()
    }),
    actions:{
        getCurrentParentID():string{
            return this.breadcrumbs[this.breadcrumbs.length-1].id
        },
        setNamespaceList(list:Namespace[]){
            this.namespaces=list
        },
    }
})
export interface LinkParams{
    user_id:number;
    namespace_id:number;
    name:string;
    authority:number;
    token:string;
}
export const DefaultLinkParams={
    user_id:0,
    namespace_id:0,
    name:"none",
    authority:0,
    token:"",
}
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
        };
        markdownEdit:{
            visible:boolean;
            text:string;
            editFileItem:FileItem;
        }
    }
    namespaces:Namespace[];
    breadcrumbs:FileItem[];
    fileList:FileItem[];
    linkParams:LinkParams;
    urlMap:Map<string,string>
}

export interface FileItem{
    id:string;
    fileName:string;
    fileType:string;
    createTime:string;
    updateTime:string;
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