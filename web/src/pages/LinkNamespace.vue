<script setup lang="ts">
import {onMounted, ref} from "vue";
import {initLinkParams, linkNamespace} from "../service/filemanage";
import {useUserStore} from "../store/user";
import {useRouter} from "vue-router";
import {useFileStore} from "../store/file";
import {FetchUserInfo} from "../service/user";
import {useToast} from "primevue/usetoast";
const fileStore=useFileStore();
const userStore=useUserStore();
const visible=ref<boolean>(true);
const router=useRouter();
const toast=useToast();
const afterHide=()=>{
  router.push("/");
};

onMounted(async ()=>{
  const user=await FetchUserInfo();
  userStore.setUser(user);
 initLinkParams();
});
</script>

<template>
<Dialog v-model:visible="visible" @after-hide="afterHide">
  用户：{{userStore.user.username}}
  你确定加入命名空间：
  {{fileStore.linkParams.name}}
  吗？
  <template #footer>
    <Button  @click="linkNamespace(fileStore.linkParams.namespace_id,fileStore.linkParams.authority,fileStore.linkParams.token,toast)">确定</Button>
    <Button @click="visible=false">取消</Button>
  </template>
</Dialog>
  <Toast></Toast>
</template>

<style scoped>

</style>