<script setup lang="ts">

import {createFolder, getCurrentBreadcrumbsPath} from "../../service/filemanage";
import {useFileStore} from "../../store/file";
import {ref} from "vue";
import Dialog from "primevue/dialog";
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import {useToast} from "primevue/usetoast";
const fileStore=useFileStore();
const folderName=ref<string>("");
const toast = useToast();
</script>

<template>
  <Toast></Toast>
  <Dialog modal v-model:visible="fileStore.dialog.createDirectory.visible">
    <template #header>
      <h4 class="mr-3">新建文件夹</h4>
    </template>
    <div>
      <span>当前所在文件路径：</span>
      <p>{{getCurrentBreadcrumbsPath()}}</p>
    </div>
    <div class="flex flex-col items-center">
      <InputText type="text" class="input" placeholder="请输入文件夹名" v-model="folderName" />
    </div>
    <template #footer>
      <Button label="取消" @click="fileStore.dialog.createDirectory.visible=false" class="mr-2" />
      <Button label="确定" @click="createFolder(folderName,toast)" class="mr-2" />
    </template>
  </Dialog>
</template>

<style scoped>

</style>