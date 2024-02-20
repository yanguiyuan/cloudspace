<template>
  <AppTemplate :options="SideMenuOptionItems">
    <template #header>
        <FileManagementHeader></FileManagementHeader>
    </template>
    <template #content>
        <div class="relative p-1">
          <div class="flex align-center rounded p-2 hover:bg-gray-200" v-for="it in fileStore.fileList.sort((a, b) => a.id.localeCompare(b.id))" :key="it.id">
            <div class="mr-1 flex align-center justify-center" v-if="it.fileType!='directory'" v-html="iconService.getIcon('file')"> </div>
            <div class="mr-1" v-if="it.fileType=='directory'" v-html="iconService.getIcon('folder')"></div>
            <div v-if="!canRename||renameFileItem.id!=it.id" @click="onClickFileItem(it)" class="fileName">
              <div >{{ getFileName(it) }}</div>
            </div>
            <IconField id="renameInput"  iconPosition="right" v-if="canRename&&renameFileItem.id==it.id">
              <InputIcon>
                <i @click="toggle" class="pi pi-check cursor-pointer" />
              </InputIcon>
              <InputText  size="small" type="text" v-model="it.fileName" />
            </IconField>
            <div id="file-options" @click="displayMenu($event,it)" v-html="iconService.getIcon('options')"  class="absolute right-0 m-2 cursor-pointer"></div>
          </div>
        </div>
    </template>
  </AppTemplate>
  <Toast></Toast>
  <OverlayPanel ref="renameOp">
    <span class="font-medium text-900 block mb-2">是否确定重命名为{{renameFileItem.fileName}}</span>
    <Button id="renameConfirm" @click="fileRename" size="small">确定</Button>
  </OverlayPanel>
  <div ref="fileTip" id="fileTip"  class="settingTooltip hidden">
    <div @click="deleteFileOrDirectory(optionFileItem,toast,confirm)" class="tipMenuItem">删除</div>
    <div id="renameButton" @click="doRename" class="tipMenuItem">重命名</div>
    <div @click="editFile(optionFileItem,toast)" v-if="canEdit(optionFileItem)" class="tipMenuItem">编辑</div>
    <div @click="downLoad"  v-if="optionFileItem?.fileType!='directory'" class="tipMenuItem">下载</div>
  </div>
  <FileUploadDialog></FileUploadDialog>
  <FileCreateDialog></FileCreateDialog>
  <Dialog :header="fileStore.clickedFileItem.fileName"  v-model:visible="fileStore.dialog.imagePreview.visible">
    <Image preview width="500" height="600"  :src="fileStore.dialog.imagePreview.url" alt="图片预览"/>
  </Dialog>
  <Dialog :header="fileStore.clickedFileItem.fileName" v-model:visible="fileStore.dialog.pdfPreview.visible">
      <div style="width:1000px;height:800px">
          <object type="application/pdf" :data="fileStore.dialog.pdfPreview.url" width="100%" height="100%">
            <p>如果您的浏览器不支持 PDF 预览，请下载 PDF 查看。</p>
          </object>
      </div>
  </Dialog>
  <Dialog :header="fileStore.clickedFileItem.fileName" v-model:visible="fileStore.dialog.txtPreview.visible">
      <v-md-preview :text="fileStore.dialog.txtPreview.text" height="500px"></v-md-preview>
  </Dialog>
  <Dialog :header="fileStore.clickedFileItem.fileName" maximizable class="w-[50vw]" v-model:visible="fileStore.dialog.markdownPreview.visible">
      <v-md-preview :text="fileStore.dialog.markdownPreview.text" height="500px"></v-md-preview>
  </Dialog>
  <ConfirmDialog></ConfirmDialog>
  <NamespaceManagementDialog></NamespaceManagementDialog>
  <MarkdownEdit></MarkdownEdit>
  <CreateTextFile></CreateTextFile>
</template>
<script setup lang="ts">
import {AssetsIconSvgService} from "../assets/assets";
import { ref, onUnmounted, onMounted, getCurrentInstance } from 'vue'
import AppTemplate from "../components/AppTemplate.vue";
import { useToast } from "primevue/usetoast";
import FileUploadDialog from "../components/file/FileUploadDialog.vue";
import {DefaultFileItem, FileItem, useFileStore} from "../store/file";
import Dialog from "primevue/dialog";
import ConfirmDialog from "primevue/confirmdialog";
import {useConfirm} from "primevue/useconfirm";
import InputText from "primevue/inputtext";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import OverlayPanel from "primevue/overlaypanel";
import Button from "primevue/button";
import {
  onClickFileItem,
  deleteFileOrDirectory,
  SideMenuOptionItems, renameFileOrDirectory, editFile, canEdit, getFileURL, unlockFile, lockFile
} from "../service/filemanage";
import FileCreateDialog from "../components/file/FileCreateDialog.vue";
import FileManagementHeader from "../components/file/FileManagementHeader.vue";
import Image from "primevue/image";
import NamespaceManagementDialog from "../components/file/NamespaceManagementDialog.vue";
import MarkdownEdit from "../components/file/MarkdownEdit.vue";
import CreateTextFile from "../components/file/CreateTextFile.vue";
const renameOp = ref();
const toast = useToast();
const confirm = useConfirm();
const fileStore=useFileStore();
const optionFileItem=ref<FileItem>(DefaultFileItem)
const renameFileItem=ref<FileItem>(DefaultFileItem)
const app = getCurrentInstance();
const iconService=AssetsIconSvgService.getInstance();
const canRename=ref(false);
const displayMenu = function (e: MouseEvent,it:FileItem) {
  if (app != null) {
    const el = app.refs.fileTip as HTMLElement
    el.style.left = `${e.clientX + 15}px`;
    el.style.top = `${e.clientY}px`;
    el.style.display='block';
  }
  optionFileItem.value=it;
}
const doRename = async function () {

  renameFileItem.value=optionFileItem.value;
  const b=await lockFile(optionFileItem.value,toast);
  if(b){
    canRename.value=true;
  }
}
const downLoad =async function () {
    const url = await getFileURL(optionFileItem.value.id);
    var downloadLink = document.createElement('a');

    // 设置链接的 href 属性为文件 URL
    downloadLink.href = url;

    // 设置下载的文件名，如果服务器没有提供文件名，则需要手动设置
    downloadLink.download = optionFileItem.value.fileName;

    // 将链接添加到 DOM 中
    document.body.appendChild(downloadLink);

    // 模拟点击链接以触发下载
    downloadLink.click();

    // 删除添加到 DOM 的链接元素
    document.body.removeChild(downloadLink);
    console.log("xiazai")
}
const getFileName = function (value: FileItem) {
  if (value.fileName.length > 30) {
    return value.fileName.slice(0, 30)+"...";
  }
  return value.fileName;
}
const fileRename = async function () {
    renameOp.value.hide();
    canRename.value=false;
    await renameFileOrDirectory(renameFileItem.value,toast);
}
const handleClickOutside = async function (e: MouseEvent) {
  const target = e.target as HTMLElement;
  if (app == null) {
    return;
  }
  const tips = app.refs.fileTip as HTMLElement;
  if (!target.closest('#file-options') && tips.style.display == "block") {
    tips.style.display = 'none';
  }
  if (!target.closest('#renameButton') &&!target.closest('#renameConfirm')&& !target.closest('#renameInput') && canRename.value) {
      canRename.value = false;
      await unlockFile(renameFileItem.value, toast)
      console.log("解锁")
      await renameFileOrDirectory(renameFileItem.value, toast);
  }
}

const toggle = (event:MouseEvent) => {
  renameOp.value.toggle(event);
}

onMounted(async () => {
  document.addEventListener("click", handleClickOutside);
});
onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
})
</script>

<style scoped>
.fileName {
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}
input {
  border-radius: 8px;
}
</style>
