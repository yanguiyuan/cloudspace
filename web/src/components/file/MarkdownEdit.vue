<script setup lang="ts">
<<<<<<< HEAD
import {useFileStore} from "../../store/file";
import {saveFileContent, unlockFile} from "../../service/filemanage";
import {useToast} from "primevue/usetoast";
const fileStore=useFileStore()
const toast=useToast();
const cancel=async ()=>{
  await unlockFile(fileStore.dialog.markdownEdit.editFileItem,toast)
  console.log("解锁")
}
=======
import {ref} from "vue";
import {useFileStore} from "../../store/file";
import {saveFileContent} from "../../service/filemanage";
import {useToast} from "primevue/usetoast";
const fileStore=useFileStore()
const toast=useToast();
>>>>>>> 641946dc (test)
</script>

<template>
  <Toast></Toast>
<<<<<<< HEAD
  <Dialog @after-hide="cancel" header="编辑" maximizable class="w-[50vw]" modal v-model:visible="fileStore.dialog.markdownEdit.visible">
=======
  <Dialog header="编辑" maximizable class="w-[50vw]" modal v-model:visible="fileStore.dialog.markdownEdit.visible">
>>>>>>> 641946dc (test)
    <v-md-editor v-model="fileStore.dialog.markdownEdit.text" height="500px"></v-md-editor>
    <template #footer>
      <Button @click="fileStore.dialog.markdownEdit.visible=false">取消</Button>
      <Button @click="saveFileContent(toast)">保存</Button>
    </template>
  </Dialog>
</template>

<style scoped>

</style>