<script setup lang="ts">

import AppTemplate from "../components/AppTemplate.vue";
import {AssetsIconSvgService} from "../assets/assets";
import {onMounted} from "vue";
import {getUsers} from "../service/admin";
import {useAdminStore} from "../store/admin";
const iconService=AssetsIconSvgService.getInstance();
const adminStore=useAdminStore();
onMounted(()=>{
  getUsers(0,10)
})
</script>

<template>
  <AppTemplate>
    <template #header>
      <div class="flex items-center">
        <div class="fileIcon" v-html="iconService.getIcon('app-admin-32')"> </div>
        <div class="ml-1">用户管理</div>
      </div>
    </template>
    <template #content>
      <div v-for="it in adminStore.users">
        <div>
          {{it.username}}
        </div>
      </div>
    </template>
  </AppTemplate>
</template>

<style scoped>

</style>