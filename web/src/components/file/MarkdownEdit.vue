<script setup lang="ts">
import {ref} from "vue";
import {useFileStore} from "../../store/file";
import {saveFileContent} from "../../service/filemanage";
import {useToast} from "primevue/usetoast";
const fileStore=useFileStore()
const toast=useToast();
</script>

<template>
  <Toast></Toast>
  <Dialog header="编辑" maximizable class="w-[50vw]" modal v-model:visible="fileStore.dialog.markdownEdit.visible">
    <v-md-editor v-model="fileStore.dialog.markdownEdit.text" height="500px"></v-md-editor>
    <template #footer>
      <Button @click="fileStore.dialog.markdownEdit.visible=false">取消</Button>
      <Button @click="saveFileContent(toast)">保存</Button>
    </template>
  </Dialog>
</template>

<style scoped>

</style>