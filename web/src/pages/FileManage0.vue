<template>
  <AppTemplate :options="optionItems">
    <template #content>
      <div class="content">
        <div class="header p-2">
          <div v-if="breadcrumbs.length!=0"   class="filePathItem filePoint">{{namespace }}</div>
          <strong v-if="breadcrumbs.length == 0"  id="currentDirectory">{{ namespace }}</strong>
          <div class="filePathContainer" v-for="(it, key) in breadcrumbs">
            <div class="filePathItem divider">/</div>
            <div v-if="key!= breadcrumbs.length - 1"   class="filePathItem filePoint">{{it }}</div>
            <strong v-if="key == breadcrumbs.length - 1"  id="currentDirectory">{{ it }}</strong>
          </div>
        </div>

        <div class="fileList p-2">
          <div class="listItem p-1 rounded" v-for="it in fileList.sort((a, b) => a.id.localeCompare(b.id))" :key="it.id">
            <div class="mr-1" v-if="it.type!='directory'" v-html="iconService.getIcon('file')"> </div>
            <div class="mr-1" v-if="it.type=='directory'" v-html="iconService.getIcon('folder')"></div>
            <div class="fileName">
              <div >{{ it.fileName }}</div>
            </div>
            <div @click="displayMenu($event,it)" v-html="iconService.getIcon('options')"  class="fileMenu mr-2"></div>
          </div>
        </div>
      </div>
    </template>
  </AppTemplate>
  <div ref="fileTip" id="fileTip"  class="settingTooltip hidden">
    <div class="tipMenuItem">删除</div>
    <div class="tipMenuItem">重命名</div>
    <div class="tipMenuItem">编辑</div>
    <div class="tipMenuItem"><a>下载</a></div>
  </div>
  <Dialog v-model:visible="maskCardVisible" :style="{ width: '50rem' }" modal header="上传文件">
    <Toast />
    <FileUpload name="demo[]" url="/api/upload" @upload="onTemplatedUpload()" :multiple="true" accept="image/*" :maxFileSize="1024*1024" @select="onSelectedFiles">
      <template #header="{ chooseCallback, uploadCallback, clearCallback, files }">
        <div class="flex flex-nowrap justify-between items-center flex-1 gap-2">
          <div class="flex gap-2">
            <Button @click="chooseCallback()" icon="pi pi-images" rounded outlined></Button>
            <Button @click="uploadEvent(uploadCallback)" icon="pi pi-cloud-upload" rounded outlined severity="success" :disabled="!files || files.length === 0"></Button>
            <Button @click="clearCallback()" icon="pi pi-times" rounded outlined severity="danger" :disabled="!files || files.length === 0"></Button>
          </div>
          <ProgressBar :value="totalSizePercent" :showValue="false" :class="['md:w-20rem h-1rem w-full md:ml-auto', { 'exceeded-progress-bar': totalSizePercent > 100 }]">
            <span class="white-space-nowrap">{{ totalSize }}B / 1Mb</span>
          </ProgressBar>
        </div>
      </template>
      <template #content="{ files, uploadedFiles, removeUploadedFileCallback, removeFileCallback }">
        <div v-if="files.length > 0">
          <h5>Pending</h5>
          <div class="flex flex-wrap p-0 sm:p-5 gap-5">
            <div v-for="(file, index) of files" :key="file.name + file.type + file.size" class="card m-0 px-6 flex flex-column border-1 surface-border align-items-center gap-3">
              <div>
                <img role="presentation" :alt="file.name" :src="file.objectURL" width="100" height="50" />
              </div>
              <span class="font-semibold">{{ file.name }}</span>
              <div>{{ formatSize(file.size) }}</div>
              <Badge value="Pending" severity="warning" />
              <Button icon="pi pi-times" @click="onRemoveTemplatingFile(file, removeFileCallback, index)" outlined rounded  severity="danger" />
            </div>
          </div>
        </div>

        <div v-if="uploadedFiles.length > 0">
          <h5>Completed</h5>
          <div class="flex flex-wrap p-0 sm:p-5 gap-5">
            <div v-for="(file, index) of uploadedFiles" :key="file.name + file.type + file.size" class="card m-0 px-6 flex flex-column border-1 surface-border align-items-center gap-3">
              <div>
                <img role="presentation" :alt="file.name" :src="file.objectURL" width="100" height="50" />
              </div>
              <span class="font-semibold">{{ file.name }}</span>
              <div>{{ formatSize(file.size) }}</div>
              <Badge value="Completed" class="mt-3" severity="success" />
              <Button icon="pi pi-times" @click="removeUploadedFileCallback(index)" outlined rounded  severity="danger" />
            </div>
          </div>
        </div>
      </template>
      <template #empty>
        <div class="flex flex-col items-center">
          <i class="pi pi-cloud-upload border-2 flex content-center justify-center rounded-full p-5 text-8xl text-400 border-400" />
          <p class="align-center mt-4 mb-0 mx-10">
            拖动文件进行上传
          </p>
        </div>
      </template>
    </FileUpload>
  </Dialog>
</template>
<script setup lang="ts">
import {AssetsIconSvgService} from "../assets/assets";
import { ref, onUnmounted, onMounted, reactive, getCurrentInstance } from 'vue'
import AppTemplate from "../components/AppTemplate.vue";
import Dialog  from 'primevue/dialog';
import Badge  from 'primevue/badge';
import Button  from 'primevue/button';
import ProgressBar from "primevue/progressbar";
const namespace=ref<string>("202121091391")
const breadcrumbs=reactive<string[]>(["test","test1"])
const maskCardVisible=ref<boolean>(false);
import FileUpload from 'primevue/fileupload';
import {FileUploadSelectEvent} from 'primevue/fileupload';
interface FileItem{
  id:string
  fileName:string
  type:string
}
const optionItems=[
  {
    tooltip:"文件上传",
    to:"/",
    icon:"upload",
    enableRouterLink:false,
    action:()=>{
      maskCardVisible.value=true;
    }
  },
]
const app = getCurrentInstance()
const iconService=AssetsIconSvgService.getInstance()
const fileList=reactive<FileItem[]>([
    {
      id:"D1",
      fileName:"test",
      type:"directory"
    },
    {
      id:"F2",
      fileName:"test1",
      type:"file"
    },
    {
      id:"F3",
      fileName:"test2",
      type:"file"
    },
    {
      id:"D4",
      fileName:"test",
      type:"directory"
    },
    {
      id:"F5",
      fileName:"test1",
      type:"file"
    },
    {
      id:"F6",
      fileName:"test2",
      type:"file"
    }

])
console.log(fileList.sort((a, b) => a.id.localeCompare(b.id)));
const displayMenu = function (e: MouseEvent,it:FileItem) {
  if (app != null) {
    const el = app.refs.fileTip as HTMLElement
    el.style.left = `${e.clientX + 15}px`;
    el.style.top = `${e.clientY}px`;
    el.style.display='block';
  }
}
const handleClickOutside = function (e: MouseEvent) {
  const target = e.target as HTMLElement;
  if(app==null){
    return
  }
  const tips = app.refs.fileTip as HTMLElement
  if (!target.closest('.fileMenu') && tips.style.display=="block") {
    tips.style.display='none';
  }
}
import { useToast } from "primevue/usetoast";
const toast = useToast();

const totalSize = ref(0);
const totalSizePercent = ref(0);
const files = ref([]);
const onRemoveTemplatingFile = (file, removeFileCallback, index) => {
  removeFileCallback(index);
  totalSize.value -= parseInt(formatSize(file.size));
  totalSizePercent.value = totalSize.value / 10;
};

const onClearTemplatingUpload = (clear) => {
  clear();
  totalSize.value = 0;
  totalSizePercent.value = 0;
};

const onSelectedFiles = (event:FileUploadSelectEvent) => {
  files.value = event.files;
  files.value.forEach((file) => {
    totalSize.value += parseInt(formatSize(file.size));
  });
};

const uploadEvent = (callback:()=>void) => {
  totalSizePercent.value = totalSize.value / 10;
  callback();
};

const onTemplatedUpload = () => {
  toast.add({ severity: "info", summary: "Success", detail: "File Uploaded", life: 3000 });
};

const formatSize = (bytes:number) => {
      const k = 1024;
      const dm = 3;
      const sizes = app?.appContext.config.globalProperties.$primevue.config.locale?.fileSizeTypes;
      if(!sizes){
        return "size undefined";
      }
      if (bytes === 0) {
        return `0 ${sizes[0]}`;
      }

      const i = Math.floor(Math.log(bytes) / Math.log(k));
      const formattedSize = parseFloat((bytes / Math.pow(k, i)).toFixed(dm));

      return `${formattedSize} ${sizes[i]}`;
}
onMounted(() => {
  document.addEventListener("click", handleClickOutside)
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


.fileList {
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


.header {
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  padding-bottom: 10px;
  margin-bottom: 10px;
}


</style>
