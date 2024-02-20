<script setup lang="ts">

import AppTemplate from "../components/AppTemplate.vue";
import {AssetsIconSvgService} from "../assets/assets";
import {onMounted, reactive, ref} from "vue";
import {getUsers, resetPassword} from "../service/admin";
import {useAdminStore} from "../store/admin";

import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Password from "primevue/password";
import FloatLabel from "primevue/floatlabel";
import {User} from "../service/user";
import {useToast} from "primevue/usetoast";
const iconService=AssetsIconSvgService.getInstance();
const adminStore=useAdminStore();
const toast=useToast();
const openResetPasswordDialog=(data:User)=>{
  adminStore.selectedUser=data;
  adminStore.dialog.resetPassword.visible=true
}
onMounted(()=>{
  getUsers(0,10,toast)
})
</script>

<template>
  <Toast></Toast>
  <Dialog header="重置密码" v-model:visible="adminStore.dialog.resetPassword.visible">
    <div class="mb-5">
      用户：{{adminStore.selectedUser.username}}
    </div>
    <FloatLabel class="mt-6">
      <Password toggleMask v-model="adminStore.dialog.resetPassword.newPassword" inputId="newPassword" />
      <label for="newPassword">新密码</label>
    </FloatLabel>
    <FloatLabel class="mt-6">
      <Password toggleMask v-model="adminStore.dialog.resetPassword.confirmPassword" inputId="confirmPassword" />
      <label for="confirmPassword">确认密码</label>
    </FloatLabel>
    <template #footer>
      <Button label="取消" @click="adminStore.dialog.resetPassword.visible=false" class="mr-2" />
      <Button label="确定" @click="resetPassword(toast)" class="mr-2" />
    </template>
  </Dialog>
  <AppTemplate>
    <template #header>
      <div class="flex items-center">
        <div class="fileIcon" v-html="iconService.getIcon('app-admin-32')"> </div>
        <div class="ml-1">用户管理</div>
      </div>
    </template>
    <template #content>
      <DataTable :value="adminStore.users" :pt="{
        bodyRow:{
          style:{
            background: 'rgb(235, 235, 235)'
          }
        },
      }">
        <Column :pt="{
          headerCell:{
            style:{
              background: 'rgb(235, 235, 235)'
            }
          },
        }" field="id" header="ID"></Column>
        <Column :pt="{
          headerCell:{
            style:{
              background: 'rgb(235, 235, 235)'
            }
          },
        }" field="username" header="用户名"></Column>
        <Column :pt="{
          headerCell:{
            style:{
              background: 'rgb(235, 235, 235)'
            }
          },
        }" field="gender" header="性别"></Column>
        <Column :pt="{
          headerCell:{
            style:{
              background: 'rgb(235, 235, 235)'
            }
          },
        }" field="email" header="邮箱"></Column>
        <Column :pt="{
          headerCell:{
            style:{
              background: 'rgb(235, 235, 235)'
            }
          },
        }" style="flex: 0 0 4rem">
          <template #body="{ data }">
            <Button label="重置密码" type="button" @click="openResetPasswordDialog(data)" text size="small" />
          </template>
        </Column>
      </DataTable>
    </template>
  </AppTemplate>
</template>

<style scoped>

</style>