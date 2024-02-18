<script setup lang="ts">
import {useFileStore} from "../../store/file";
import {useToast} from "primevue/usetoast";
import FloatLabel from "primevue/floatlabel";
import{createTextFile} from "../../service/filemanage";
const fileStore=useFileStore();
const toast=useToast()
</script>

<template>
  <Toast></Toast>
  <Dialog maximizable header="新建文件" class="w-[50vw]" modal v-model:visible="fileStore.dialog.createTextFile.visible">
    <FloatLabel class="my-6">
      <InputText id="filename" v-model="fileStore.dialog.createTextFile.fileName" />
      <label for="filename">文件名</label>
    </FloatLabel>
    <v-md-editor v-model="fileStore.dialog.createTextFile.text" height="500px"></v-md-editor>
    <template #footer>
      <Button @click="fileStore.dialog.createTextFile.visible=false">取消</Button>
      <Button @click="createTextFile(toast)">保存</Button>
    </template>
  </Dialog>
</template>

<style scoped>

</style>