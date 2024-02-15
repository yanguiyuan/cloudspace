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
    alert("请重新登录");
  }
  console.log(params);
});
</script>

<template>
<Dialog v-model:visible="visible" @after-hide="afterHide">
  你好
</Dialog>
</template>

<style scoped>

</style>