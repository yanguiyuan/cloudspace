<script setup lang="ts">
import {onMounted, ref} from "vue";
import {getUrlParams} from "../service/filemanage";
import {useUserStore} from "../store/user";
import {useRouter} from "vue-router";
const userStore=useUserStore();
const visible=ref<boolean>(true);
const router=useRouter();
const afterHide=()=>{
  router.push("/");
};

onMounted(()=>{
  const params=getUrlParams(window.location.href);
  if(params.user_id===userStore.user.id.toString()){
    userStore.linkNamespace=true;
    router.push("/login");
  }
  console.log(params);
});
</script>

<template>
<Dialog v-model:visible="visible" @after-hide="afterHide">
  你确定加入命名空间：
  <span></span>
</Dialog>
</template>

<style scoped>

</style>