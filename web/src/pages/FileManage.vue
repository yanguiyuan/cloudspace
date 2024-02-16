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
    <Button @click="fileRename" size="small">确定</Button>
  </OverlayPanel>
  <div ref="fileTip" id="fileTip"  class="settingTooltip hidden">
    <div @click="deleteFileOrDirectory(optionFileItem,toast,confirm)" class="tipMenuItem">删除</div>
    <div id="renameButton" @click="canRename=true;renameFileItem=optionFileItem" class="tipMenuItem">重命名</div>
    <div v-if="optionFileItem?.fileType!='directory'" class="tipMenuItem">编辑</div>
    <div  v-if="optionFileItem?.fileType!='directory'" class="tipMenuItem"><a :href="fileStore.urlMap.get(optionFileItem.id)">下载</a></div>
  </div>
  <FileUploadDialog></FileUploadDialog>
  <FileCreateDialog></FileCreateDialog>
  <Dialog v-model:visible="fileStore.dialog.imagePreview.visible">
    <Image preview width="500" height="600"  :src="fileStore.dialog.imagePreview.url" alt="图片预览"/>
  </Dialog>
  <ConfirmDialog></ConfirmDialog>
  <NamespaceManagementDialog></NamespaceManagementDialog>
</template>
<script setup lang="ts">
import {AssetsIconSvgService} from "../assets/assets";
import { ref, onUnmounted, onMounted, getCurrentInstance } from 'vue'
import AppTemplate from "../components/AppTemplate.vue";
import { useToast } from "primevue/usetoast";
import FileUploadDialog from "../components/file/FileUploadDialog.vue";
import {useFileStore} from "../store/file";
import Dialog from "primevue/dialog";
import ConfirmDialog from "primevue/confirmdialog";
import {useConfirm} from "primevue/useconfirm";
import InputText from "primevue/inputtext";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import OverlayPanel from "primevue/overlaypanel";
import Button from "primevue/button";
import {
  FileItem, onClickFileItem, getRootFileItemID, initDefaultFileItemList,
  deleteFileOrDirectory, getFileItemByID,
  SideMenuOptionItems, DefaultFileItem, renameFileOrDirectory, getUserNamespaces
} from "../service/filemanage";
import FileCreateDialog from "../components/file/FileCreateDialog.vue";
import FileManagementHeader from "../components/file/FileManagementHeader.vue";
import Image from "primevue/image";
import NamespaceManagementDialog from "../components/file/NamespaceManagementDialog.vue";
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
const getFileName = function (value: FileItem) {
  if (value.fileName.length > 30) {
    return value.fileName.slice(0, 30)+"...";
  }
  return value.fileName;
}
const fileRename = async function () {
    renameOp.value.hide();
    await renameFileOrDirectory(renameFileItem.value,toast);
}
const handleClickOutside = function (e: MouseEvent) {
  const target = e.target as HTMLElement;
  if(app==null){
    return;
  }
  const tips = app.refs.fileTip as HTMLElement;
  if (!target.closest('#file-options') && tips.style.display=="block") {
    tips.style.display='none';
  }
  if (!target.closest('#renameButton')&&!target.closest('#renameInput') && canRename.value) {
    canRename.value=false;
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
