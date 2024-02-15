import {defineStore} from "pinia";
import { FileManagementState} from "../service/filemanage";
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
           }
       },
        breadcrumbs:[],
        fileList:[EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem,EmptyFileItem],
        urlMap:new Map<string, string>()
    }),
    actions:{
        getCurrentParentID():string{
            return this.breadcrumbs[this.breadcrumbs.length-1].id
        },
    }
})