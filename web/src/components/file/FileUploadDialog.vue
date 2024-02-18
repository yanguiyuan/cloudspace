<script setup lang="ts">

import Dialog from "primevue/dialog";
import Badge from "primevue/badge";
import FileUpload, {
  FileUploadBeforeSendEvent,
  FileUploadSelectEvent, FileUploadUploadEvent
} from "primevue/fileupload";
import Button from "primevue/button";
import ProgressBar from "primevue/progressbar";
const fileStore=useFileStore()
import { useToast } from "primevue/usetoast";
import {getCurrentInstance, ref} from "vue";
import {useFileStore} from "../../store/file";
const toast = useToast();
const app = getCurrentInstance()
const totalSize = ref(0);
const totalSizePercent = ref(0);
const files = ref<Array<File>>([]);
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

const onTemplatedUpload = (e: FileUploadUploadEvent) => {
  const resp=JSON.parse(e.xhr.response)
  if(resp.code==0){
    fileStore.fileList.push(resp.data)
  }
  console.log(resp)
  toast.add({ severity: "success", summary: "成功", detail: "上传成功", life: 3000 });
};

const onBeforeSend = (e: FileUploadBeforeSendEvent) => {
  const formData = new FormData();
  for (const file of files.value) {
    formData.append("data", file);
  }
  e.formData=formData;
  let token=localStorage.getItem("token")
  if(token==null){
    token=""
  }else{
    token=token.substring(1,token.length-1)
  }
  let tokenHeader="Bearer ".concat(token)
  e.xhr.setRequestHeader("Authorization", tokenHeader)
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
</script>

<template>
  <Dialog v-model:visible="fileStore.dialog.upload.visible" :style="{ width: '50rem' }" modal header="上传文件">
    <Toast />
    <FileUpload  :max-file-size="1021*1021*1024" name="data" :url="'http://localhost:8888/user/file/'+fileStore.getCurrentParentID()+'/upload'" @upload="onTemplatedUpload" :multiple="true" @select="onSelectedFiles" @before-send="onBeforeSend">
      <template #header="{ chooseCallback, uploadCallback, clearCallback, files }">
        <div class="flex flex-nowrap justify-between items-center flex-1 gap-2">
          <div class="flex gap-2">
            <Button @click="chooseCallback()" icon="pi pi-file" rounded outlined></Button>
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
              <div v-if="file.type=='png'||file.type=='jpg'">
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
              <div  v-if="file.type=='png'||file.type=='jpg'">
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

<style scoped>

</style>