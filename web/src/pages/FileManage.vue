<template>
  <div class="card">
    <div id="sideCard">
      <div class="menuContainer">
        <div class="menuItem">
          <RouterLink to="/">
            <svg t="1671204857123" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
              p-id="6487" width="32" height="32">
              <path
                d="M471.893333 149.333333a42.666667 42.666667 0 0 0-73.258666-29.781333l-343.893334 352.981333a42.666667 42.666667 0 0 0-0.768 58.709334l343.893334 372.352a42.666667 42.666667 0 0 0 73.984-28.928v-190.677334c56.917333-5.248 116.821333-1.365333 179.882666 11.989334 65.834667 13.994667 150.528 76.032 253.909334 202.24a42.666667 42.666667 0 0 0 75.477333-31.36c-15.445333-152.32-73.984-281.301333-176.170667-384.853334-92.757333-93.994667-204.373333-146.432-333.098666-156.586666V149.333333z"
                fill="#e6e6e6" fill-opacity=".85" p-id="6488"></path>
            </svg>
          </RouterLink>
          <div class="tooltip">返回</div>
        </div>
        <div @click="showUploadPanel" class="menuItem">
          <svg t="1671205618913" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="7544" width="32" height="32">
            <path
              d="M311.466667 64c4.693333 0 8.533333 3.84 8.533333 8.533333v46.933334a8.533333 8.533333 0 0 1-8.533333 8.533333H234.666667a106.666667 106.666667 0 0 0-106.56 102.037333L128 234.666667v554.666666a106.666667 106.666667 0 0 0 102.037333 106.56L234.666667 896h554.666666a106.666667 106.666667 0 0 0 106.56-102.037333L896 789.333333V234.666667a106.666667 106.666667 0 0 0-102.037333-106.56L789.333333 128h-55.466666a8.533333 8.533333 0 0 1-8.533334-8.533333V72.533333c0-4.693333 3.84-8.533333 8.533334-8.533333H789.333333a170.666667 170.666667 0 0 1 170.666667 170.666667v554.666666a170.666667 170.666667 0 0 1-170.666667 170.666667H234.666667a170.666667 170.666667 0 0 1-170.666667-170.666667V234.666667a170.666667 170.666667 0 0 1 170.666667-170.666667h76.8z m212.138666 0l2.517334 0.298667c4.138667 0.746667 8.106667 2.709333 11.306666 5.909333l185.386667 185.386667a8.533333 8.533333 0 0 1 2.517333 6.058666v66.368a8.533333 8.533333 0 0 1-14.570666 6.037334L554.666667 177.962667V632.96a8.533333 8.533333 0 0 1-8.533334 8.533333h-46.933333a8.533333 8.533333 0 0 1-8.533333-8.533333V177.322667l-156.096 156.096a8.533333 8.533333 0 0 1-14.570667-6.037334v-66.346666a8.533333 8.533333 0 0 1 2.496-6.037334l184.768-184.789333A21.248 21.248 0 0 1 521.088 64h2.517333z"
              fill="#e6e6e6" p-id="7545"></path>
          </svg>
          <div class="tooltip">上传文件</div>
        </div>
        <div @click="fileCreatePanelState.visible.visible=true" class="menuItem">
          <svg t="1671206543009" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="16208" width="32" height="32">
            <path
              d="M842.666667 285.866667l-187.733334-187.733334c-14.933333-14.933333-32-21.333333-53.333333-21.333333H234.666667C194.133333 74.666667 160 108.8 160 149.333333v725.333334c0 40.533333 34.133333 74.666667 74.666667 74.666666h554.666666c40.533333 0 74.666667-34.133333 74.666667-74.666666V337.066667c0-19.2-8.533333-38.4-21.333333-51.2z m-44.8 44.8H618.666667c-6.4 0-10.666667-4.266667-10.666667-10.666667V140.8l189.866667 189.866667z m-8.533334 554.666666H234.666667c-6.4 0-10.666667-4.266667-10.666667-10.666666V149.333333c0-6.4 4.266667-10.666667 10.666667-10.666666h309.333333V320c0 40.533333 34.133333 74.666667 74.666667 74.666667h181.333333V874.666667c0 6.4-4.266667 10.666667-10.666667 10.666666z"
              fill="#e6e6e6" p-id="16209"></path>
            <path
              d="M618.666667 586.666667h-74.666667V512c0-17.066667-14.933333-32-32-32s-32 14.933333-32 32v74.666667H405.333333c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h74.666667V725.333333c0 17.066667 14.933333 32 32 32s32-14.933333 32-32v-74.666666H618.666667c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32z"
              fill="#e6e6e6" p-id="16210"></path>
          </svg>
          <div class="tooltip">新建文件</div>
        </div>
        <div class="menuItem">
          <svg t="1671206682427" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="16853" width="32" height="32">
            <path
              d="M512 1024C230.4 1024 0 793.6 0 512S230.4 0 512 0s512 230.4 512 512-230.4 512-512 512z m259.2-569.6H480c-12.8 0-25.6 12.8-25.6 25.6v64c0 12.8 12.8 25.6 25.6 25.6h176c12.8 0 25.6 12.8 25.6 25.6v12.8c0 41.6-35.2 76.8-76.8 76.8h-240c-12.8 0-25.6-12.8-25.6-25.6V416c0-41.6 35.2-76.8 76.8-76.8h355.2c12.8 0 25.6-12.8 25.6-25.6v-64c0-12.8-12.8-25.6-25.6-25.6H416c-105.6 0-188.8 86.4-188.8 188.8V768c0 12.8 12.8 25.6 25.6 25.6h374.4c92.8 0 169.6-76.8 169.6-169.6v-144c0-12.8-12.8-25.6-25.6-25.6z"
              fill="#e6e6e6" p-id="16854"></path>
          </svg>
          <div class="tooltip">从gitee拉取</div>
        </div>
      </div>
    </div>

    <div class="content">

      <div class="header">
        <div @click="update(rootFileItemData)" v-if="fileBreadcrumbState.path.length" class="filePathItem filePoint">{{
          userData.username
        }}</div>
        <strong v-if="!fileBreadcrumbState.path.length" id="currentDirectory">{{ userData.username }}</strong>
        <div class="filePathContainer" v-for="item, key in fileBreadcrumbState.path">
          <div class="filePathItem divider">/</div>
          <div @click="update(item)" v-if="key != fileBreadcrumbState.path.length - 1" class="filePathItem filePoint">{{
            item.fileName }}
          </div>
          <strong v-if="key == fileBreadcrumbState.path.length - 1" id="currentDirectory">{{ item.fileName }}</strong>
        </div>
      </div>

      <div class="fileList">
        <div class="listItem" v-for="item in fileListData.value.sort((a, b) => a.id - b.id)" :key="item.id">
          <svg v-if="item.type != 'directory'" t="1671732414196" class="icon" viewBox="0 0 1024 1024" version="1.1"
            xmlns="http://www.w3.org/2000/svg" p-id="2644" width="32" height="32">
            <path
              d="M842.666667 285.866667l-187.733334-187.733334c-14.933333-14.933333-32-21.333333-53.333333-21.333333H234.666667C194.133333 74.666667 160 108.8 160 149.333333v725.333334c0 40.533333 34.133333 74.666667 74.666667 74.666666h554.666666c40.533333 0 74.666667-34.133333 74.666667-74.666666V337.066667c0-19.2-8.533333-38.4-21.333333-51.2z m-44.8 44.8c-2.133333 2.133333-4.266667 0-8.533334 0h-170.666666c-6.4 0-10.666667-4.266667-10.666667-10.666667V149.333333c0-2.133333 0-6.4-2.133333-8.533333 0 0 2.133333 0 2.133333 2.133333l189.866667 187.733334z m-8.533334 554.666666H234.666667c-6.4 0-10.666667-4.266667-10.666667-10.666666V149.333333c0-6.4 4.266667-10.666667 10.666667-10.666666h311.466666c-2.133333 4.266667-2.133333 6.4-2.133333 10.666666v170.666667c0 40.533333 34.133333 74.666667 74.666667 74.666667h170.666666c4.266667 0 6.4 0 10.666667-2.133334V874.666667c0 6.4-4.266667 10.666667-10.666667 10.666666z"
              fill="#666666" p-id="2645"></path>
            <path
              d="M640 693.333333H341.333333c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h298.666667c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32zM640 522.666667H341.333333c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h298.666667c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32zM341.333333 416h85.333334c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32h-85.333334c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32z"
              fill="#666666" p-id="2646"></path>
          </svg>
          <svg v-if="item.type == 'directory'" t="1671863859291" class="icon" viewBox="0 0 1024 1024" version="1.1"
            xmlns="http://www.w3.org/2000/svg" p-id="2669" width="32" height="32">
            <path
              d="M810.666667 85.333333a85.333333 85.333333 0 0 1 85.333333 85.333334v152.021333c36.821333 9.493333 64 42.88 64 82.645333v405.333334a128 128 0 0 1-128 128H192a128 128 0 0 1-128-128V298.666667a85.376 85.376 0 0 1 64-82.645334V170.666667a85.333333 85.333333 0 0 1 85.333333-85.333334h597.333334zM128.149333 296.170667L128 298.666667v512a64 64 0 0 0 60.245333 63.893333L192 874.666667h640a64 64 0 0 0 63.893333-60.245334L896 810.666667V405.333333a21.333333 21.333333 0 0 0-18.837333-21.184L874.666667 384H638.165333l-122.069333-101.717333a21.333333 21.333333 0 0 0-10.688-4.736l-2.986667-0.213334H149.333333a21.333333 21.333333 0 0 0-21.184 18.837334zM535.189333 213.333333l127.978667 106.666667H832V170.666667a21.333333 21.333333 0 0 0-18.837333-21.184L810.666667 149.333333H213.333333a21.333333 21.333333 0 0 0-21.184 18.837334L192 170.666667v42.666666h343.168z"
              fill="#333333" p-id="2670"></path>
          </svg>
          <div class="fileName">
            <div @click="preview(item)" v-if="!item.renameComponent.visible.visible">{{ item.fileName }}</div>
            <div class="renameDiv" v-if="item.renameComponent.visible.visible"><input
                v-model="item.renameComponent.newName" type="text">{{ item.renameComponent.newName }}
            </div>
          </div>

          <div @click="displayMenu($event,item)" class="fileMenu">
            <svg t="1671208100554" class="icon" viewBox="0 0 1026 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
              p-id="17969" width="16" height="16">
              <path d="M110.98112 510.74048m-110.98112 0a21.676 21.676 0 1 0 221.96224 0 21.676 21.676 0 1 0-221.96224 0Z"
                fill="#231815" p-id="17970"></path>
              <path d="M512.93184 510.74048m-110.98112 0a21.676 21.676 0 1 0 221.96224 0 21.676 21.676 0 1 0-221.96224 0Z"
                fill="#231815" p-id="17971"></path>
              <path d="M914.87744 510.74048m-110.98112 0a21.676 21.676 0 1 0 221.96224 0 21.676 21.676 0 1 0-221.96224 0Z"
                fill="#231815" p-id="17972"></path>
            </svg>
          </div>
        </div>
      </div>

    </div>

    <YMaskCard :visible="fileCreatePanelState.visible.visible" v-on:on-close="fileCreatePanelState.visible.visible=false">
      <div>
        当前位置：{{ userData.username + "/" + getPath() }}
      </div>
      <div>
        文件名：
        <input v-model="fileCreatePanelState.createdFileName">
      </div>
      <div @click="onConfirm" class="comfirm-btn">
        <YButton> 确定 </YButton>
      </div>
    </YMaskCard>

  </div>

  <div ref="fileTip" :style="{display:fileTipsMenuState.display}" class="settingTooltip">
    <div v-if="fileTipsMenuState.currentFile?.type != 'directory'" @click="fileTipsMenuState.deleteFile"
      class="tipMenuItem">删除</div>
    <div @click="fileTipsMenuState.renameFile" class="tipMenuItem">重命名</div>
    <div @click="fileTipsMenuState.editFile" v-if="fileTipsMenuState.currentFile?.type != 'directory'"
      class="tipMenuItem">
      编辑</div>
    <div v-if="fileTipsMenuState.currentFile?.type != 'directory'" class="tipMenuItem"><a
        v-bind:href="fileTipsMenuState.currentFile?.url">下载</a></div>
  </div>

  <Drawer :title="drawerComponentState.title" :visible="drawerComponentState.visible" @drawerClose="onClose">
    <iframe v-if="pdfPreviewComponentState.visible.visible" :src="pdfPreviewComponentState.url" frameborder="0"
      width="100%" height="100%"></iframe>
    <v-md-preview v-if="markdownPreViewState.visible.visible" :text="markdownPreViewState.markdownText"></v-md-preview>
    <v-md-editor v-if="markdownEditorState.visible.visible" v-model="markdownEditorState.markdownText" />
  </Drawer>

  <el-dialog v-model="fileUploadPanelState.visible.visible">
    <el-upload :on-success="refresh" drag :action="fileUploadPanelState.url" multiple>
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div>
        拖拽文件或者 <em>点击上传</em>
      </div>
    </el-upload>
  </el-dialog>
</template>
<script setup lang="ts">
import YButton from '../components/YButton.vue'
import YMaskCard from '../components/YMaskCard.vue';
import { ElMessage } from 'element-plus';
import { ref, onUnmounted, onMounted, reactive, getCurrentInstance } from 'vue'
import {cf as axios} from '../axios/axios';
import router from '../router/router';
import {Runtime} from "inspector";
const app = getCurrentInstance()
///用户状态数据
interface UserData {
    token: string
    username: string
    namespaceID:number
}
const userData =reactive<UserData>({
    token:"",
    username:"",
    namespaceID:0
})

interface FileLink{
    id:number
    fileID:number
    url:string
    path:string
    createTime:Date
    updateTime:Date
}
interface FileItem{
    id:number
    name:string
    type:string
    size:number
    fileLink:FileLink
    createTime:Date
    updateTime:Date
}
const fileList=ref<Array<FileItem>>([])
//控件显示状态
const getFileListByParentID=async (id:number)=>{
    const resp=await axios.get("/auth/files/"+id)
    fileList.value=resp.data
}

onMounted(()=>{
    //更新用户数据
    let str=localStorage.getItem("user")
    if (str==null){
        ElMessage.error("用户信息获取失败，请重新登录")
        return
    }
    let jsonObj=JSON.parse(str)
    userData.namespaceID=jsonObj.userIndex.namespaceID
    userData.username=jsonObj.username
    str=localStorage.getItem("token")
    if (str==null){
        ElMessage.error("Token获取失败，请重新登录")
        return
    }
    userData.token=str
    //获取用户文件项
    getFileListByParentID(userData.namespaceID)
})
class ControlVisibleState {
  visible: boolean
  constructor(visible: boolean) {
    this.visible = visible
  }
}

//文件重命名组件状态
class FileRenameComponentState {
  visible: ControlVisibleState
  newName: string
  constructor(visible: ControlVisibleState, newName: string) {
    this.visible = visible
    this.newName = newName
  }
}
// const userData = reactive<UserData>({
//   token: "",
//   username: ""
// })

//文件数据（整颗文件树）
interface FileData {
  fileNodes: Array<FileNode>
  fileMap: Map<number, FileSpecificInfo>
}
const fileData = ref<FileData>({
  fileNodes: [],
  fileMap: new Map<number, FileSpecificInfo>(),
})
const fetchFileData = async function () {
  const userStr = localStorage.getItem("user")
  if (userStr != null) {
    const userInfo = JSON.parse(userStr)
    userData.token = userInfo.token
    userData.username = userInfo.username
    await axios.get("/usr/files", {
      params: {
        token: userData.token,
      }
    }).then((res) => {
      fileData.value.fileNodes = res.data.Data.fileNodes
      fileData.value.fileMap = res.data.Data.fileMap
      console.log("www:", fileData.value)
    }).catch((err) => {
      console.log(err)
    })
  }
}
//文件项数据
class FileItemData {
  id: number
  fileName: string
  type: string | undefined
  url: string | undefined
  renameComponent: FileRenameComponentState
  constructor(fileName: string, id: number, type: string, url: string, renameComponent: FileRenameComponentState) {
    this.id = id
    this.fileName = fileName
    this.type = type
    this.url = url
    this.renameComponent = renameComponent
  }


}
const displayMenu = function (e: MouseEvent,item:FileItemData) {

  if (app != null) {
    const el = app.refs.fileTip as HTMLElement
    el.style.left = `${e.clientX + 10}px`
    el.style.top = `${e.clientY}px`
    fileTipsMenuState.currentFile = item
  }
  fileTipsMenuState.display="block"
}
const preview = function (file: FileItemData) {
  console.log("x", file)
  drawerComponentState.title = file.fileName
  if (file.type == ".pdf") {
    if (file.url != undefined) {
      pdfPreviewComponentState.url = file.url
      pdfPreviewComponentState.visible.visible=true
      drawerComponentState.visible=true
    }
  } else if (file.type == ".md" || file.type == ".txt") {
    if (file.url != undefined) {
      fetch(file.url).then((res) => {
        return res.text()
      }).then((res) => {
        console.log(res)
        markdownPreViewState.markdownText = res
        markdownPreViewState.visible.visible = true
        drawerComponentState.visible = true
      })
    }

  }
  else if (file.type == "directory") {
    
    fileBreadcrumbState.path.push(file)
    update(file)
    console.log(fileBreadcrumbState.path)
    console.log("path:",getPath())
  }
}
const rootFileItemData = reactive<FileItemData>(new FileItemData("root", 0, "directory", "", new FileRenameComponentState(new ControlVisibleState(false), "")))

//文件面包屑数据
interface FileBreadcrumbState {
  path: Array<FileItemData>
}
const fileBreadcrumbState = reactive<FileBreadcrumbState>({
  path: []
})
const getPath = function () {
  let s: string
  s = ""
  for (let p of fileBreadcrumbState.path) {
    s += p.fileName + "/"
  }
  return s
}

const getCurrentDirectory = function () {
  if (fileBreadcrumbState.path.length == 0) {
    return rootFileItemData
  } else {
    return fileBreadcrumbState.path[fileBreadcrumbState.path.length - 1]
  }
}
//当前文件列表数据
interface FileListData {
  value: Array<FileItemData>

}
const fileListData = reactive<FileListData>({
  value: [],
})
const updateFromRemote = async function (file: FileItemData) {
  await fetchFileData()
  console.log("ss:", fileData.value.fileNodes)
  update(file)
}
const refresh = function () {
  console.log("currentdir:",getCurrentDirectory())
  
  updateFromRemote(getCurrentDirectory())
  console.log("this:",userData.token)
  fileUploadPanelState.url="http://127.0.0.1:8088/usr/files?path=" + getPath() + "&token=" + userData.token

}
const showUploadPanel=function(){
  fileUploadPanelState.visible.visible=true
  fileUploadPanelState.url="http://127.0.0.1:8088/usr/files?path=" + getPath() + "&token=" + userData.token
  console.log(fileUploadPanelState.url)

}
const update = function (file: FileItemData) {
  console.log("a1:",file)
  const index=fileBreadcrumbState.path.indexOf(file)
  if(index>=0){
    fileBreadcrumbState.path.splice(index+1)
    console.log("yy",fileBreadcrumbState.path)
  }else{
    fileBreadcrumbState.path.splice(0)
    console.log("xx",fileBreadcrumbState.path)
  }
  console.log("index",index)
  fileListData.value.splice(0, fileListData.value.length)
  console.log("sxz:",fileListData.value)
  const fileList: Array<FileNode> = fileData.value.fileNodes
    console.log("a2:",fileList)
  const fileMap = new Map(Object.entries(fileData.value.fileMap))
  if (fileList != undefined) {
    for (let node of fileList) {
      if (node.ParentID == file.id) {
        console.log("a3:",node)
        if (node.IsDir) {
          console.log("a4:")
          fileListData.value.push(new FileItemData(node.FileName, node.ID, "directory", "", new FileRenameComponentState(new ControlVisibleState(false), "")))
        } else {
          fileListData.value.push(
            new FileItemData(node.FileName, node.ID, fileMap.get(node.MapID.toString())?.Type,
              fileMap.get(node.MapID.toString())?.Url, new FileRenameComponentState(new ControlVisibleState(false), "")))
        }
      }
    }
  }
  console.log("sxxxxz:",fileListData.value)
}
//------------------------------------------分隔线--------------------------------------------------



//文件上传面板状态
interface FileUploadPanelState {
  visible: ControlVisibleState
  url: string
}
const fileUploadPanelState = reactive<FileUploadPanelState>({
  visible: new ControlVisibleState(false),
  url: "http://127.0.0.1:8088/usr/files?path=" + getPath() + "&token=" + userData.token,
})


//文件操作提示菜单状态
interface FileTipsMenuState {
  display:string
  currentFile: FileItemData | undefined
  deleteFile(): void
  editFile(): void
  renameFile(): void
}
const fileTipsMenuState = reactive<FileTipsMenuState>({
  display:"none",
  currentFile: undefined,
  deleteFile() {
    axios.delete("/usr/files?token=" + userData.token, {
      params: {
        namespace: userData.username,
        path: getPath(),
        fileName: this.currentFile?.fileName
      }
    }).then(res => {
      refresh()
      console.log("呵呵呵",fileListData.value.length)
      if (fileListData.value.length <= 1) {
        console.log("啊哈哈")
        fileBreadcrumbState.path.pop()
        update(getCurrentDirectory())
      }
      ElMessage.success("删除成功");
    }).catch(res => {
      ElMessage.error("删除失败")
    })
  },
  editFile() {
    if (this.currentFile != undefined) {
      if (this.currentFile.type == ".md") {
        if (this.currentFile.url != undefined) {
          fetch(this.currentFile.url).then((res) => {
            return res.text()
          }).then((res) => {
            drawerComponentState.visible=true
            markdownEditorState.visible.visible=true
            markdownEditorState.markdownText = res
          })
        }

      }
    }

  },
  renameFile() {
    // this.currentFile?.renameComponent.visible.visible=true
  }
})

//文件创建面板状态
interface FileCreatePanelState {
  visible: ControlVisibleState
  createdFileName: string
}
const fileCreatePanelState = reactive<FileCreatePanelState>({
  visible: new ControlVisibleState(false),
  createdFileName: "",
})
const onConfirm = function () {
  const fileName = fileCreatePanelState.createdFileName
  if (fileName.endsWith("/")) {
    console.log("创建文件夹：", fileName)
    axios.post("/usr/directory?token=" + userData.token, {
      path: getPath(),
      dirName: fileName,
    }).then(res => {
      ElMessage.success("创建成功")
      fileCreatePanelState.createdFileName = ""
      fileCreatePanelState.visible.visible=false
      refresh()
    })
  } else {
    console.log("创建文件：", fileName)
  }
}
//pdf预览组件状态
interface PDFPreviewComponentState {
  visible: ControlVisibleState
  url: string
}
const pdfPreviewComponentState = reactive<PDFPreviewComponentState>({
  visible: new ControlVisibleState(false),
  url: ""
})

//markdown预览和编辑组件状态
class MarkdownComponentState {
  visible: ControlVisibleState
  markdownText: string
  constructor(visible: ControlVisibleState, text: string) {
    this.visible = visible
    this.markdownText = text
  }
}
const markdownPreViewState = reactive<MarkdownComponentState>(new MarkdownComponentState(new ControlVisibleState(false), ""))
const markdownEditorState = reactive<MarkdownComponentState>(new MarkdownComponentState(new ControlVisibleState(false), ""))

//抽屉组件状态
interface DrawerComponentState {
  visible: boolean
  title: string
}
const drawerComponentState = reactive<DrawerComponentState>({
  visible: false,
  title: "",
})
const onClose = function () {
  drawerComponentState.visible=false
  console.log("llll")
  console.log(drawerComponentState.visible)
  pdfPreviewComponentState.visible.visible=false
  markdownPreViewState.visible.visible=false
  if (markdownEditorState.visible.visible) {
    axios.put("/usr/files?token=" + userData.token, {
      "path": getPath(),
      "fileName": fileTipsMenuState.currentFile?.fileName,
      "content": markdownEditorState.markdownText,
    }).then((res) => {
      if (fileTipsMenuState.currentFile != undefined) {
        refresh()
        markdownEditorState.visible.visible=false
      }
    })
  }
}


interface FileNode {
  ID: number
  Deep: number
  FileName: string
  IsDir: boolean
  MapID: number
  ParentID: number
}
interface SqlTime {
  Time: string
  Valid: boolean
}
interface SqlNumber {
  Int64: number
  Valid: boolean
}
interface FileSpecificInfo {
  Author: string
  CreateTime: SqlTime
  FileName: string
  Id: SqlNumber
  Namespace: string
  Path: string
  Size: number
  Type: string
  UpdateTime: SqlTime
  Url: string
}







const handleClickOutside = function (e: MouseEvent) {
  const target = e.target as HTMLElement;
  if (!target.closest('.fileMenu') && fileTipsMenuState.display=="block") {
    fileTipsMenuState.display="none"
  }
  // if (!target.closest(".renameDiv") && !target.closest('.settingTooltip') && !target.closest('.fileMenu')) {
  //   if (fileTipsMenuState.currentFile?.fileName != fileTipsMenuState.currentFile?.renameComponent.newName) {
  //     axios.post("/usr/rename?token=" + userData.token, {
  //       "path": getPath(),
  //       "oldName": fileTipsMenuState.currentFile?.fileName,
  //       "newName": fileTipsMenuState.currentFile?.renameComponent.newName
  //     }).then((res) => {
  //       refresh()
  //     })
  //     console.log("命令成功")
  //   }
  // }
}

onMounted(() => {
  document.addEventListener("click", handleClickOutside)
  refresh()
  
}
)
onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside)
})
</script>

<style scoped>
.fileName {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.uploadWindow {
  background-color: rgb(0, 235, 235) !important;
  border-radius: 25px !important;
}

.divider {
  color: rgba(0, 0, 0, 0.4);
  margin: 0 0.2rem 0;
}

.filePathContainer {
  display: inline-block;
}

.filePathItem {
  display: inline-block;
}

.comfirm-btn {
  position: absolute;
  bottom: 50px;
  right: 0;
  left: 0;
}

.filePoint {
  color: #005980;
  cursor: pointer;
}

.filePoint:hover {
  color: #537c8d;
}

input {
  border-radius: 8px;
}

.card {
  position: fixed;
  bottom: 20%;
  top: 20%;
  left: 20%;
  right: 20%;
  border-radius: 15px;
  background-color: rgb(235, 235, 235);
  box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
}

.info-card {
  position: fixed;
  bottom: 40%;
  top: 40%;
  left: 30%;
  right: 30%;
  border-radius: 15px;
  background-color: rgb(235, 235, 235);
  box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
  padding: 10px;
}

.info-card-mask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 9999;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}

.info-card-content {
  margin: 10px;
  position: absolute;
  top: 15px;
  left: 0;
  right: 0;
  bottom: 0;
}

.info-card-close {
  display: flex;
  justify-content: flex-end;
  cursor: pointer;
  z-index: 999;
}

#sideCard {
  position: absolute;
  height: calc(100vh - 300px);
  width: 90px;
  left: -100px;
  background-color: rgba(0, 0, 0, 0.1);
  border-radius: 15px;
  backdrop-filter: blur(30px);
  transition: 0.15s;
  box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
}

.fileList {
  position: fixed;
  bottom: 20%;
  top: 25%;
  left: 20%;
  right: 20%;
  border-radius: 0 0 15px 15px;
  background-color: rgb(235, 235, 235);
  overflow: auto;
}

.listItem {
  display: flex;
  position: relative;
  margin: 5px 0 5px 0;


}

.listItem:hover {
  background-color: rgb(200, 200, 200);
}

.fileMenu {
  position: absolute;
  right: 0;
  top: 8px;
  cursor: pointer;
}


.content {
  height: calc(100vh - 300px);
}

.menuContainer {
  margin: 10px;
  text-align: center;
}

.header {
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  padding-bottom: 10px;
  margin-bottom: 10px;
}

.menuItem {
  margin: 10px 0 10px 0;
  position: relative;
}

.menuItem:hover {
  margin: 10px 0 10px 0;
  transform: translateX(10px);
}

.menuItem:hover .tooltip {
  display: block;
}

.tooltip {
  display: none;
  position: absolute;
  top: -15px;
  /* left: 55px; */
  right: 60px;
  background-color: #ccc;
  padding: 10px;
  border-radius: 5px;
  transition: 0.25s;
}
</style>
