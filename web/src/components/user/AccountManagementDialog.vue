<script setup lang="ts">
import {useUserStore} from "../../store/user";
import Dialog from "primevue/dialog";
import {onMounted, ref} from "vue";
import {FetchUserInfo, modifyUserInfo} from "../../service/user";
import {useToast} from "primevue/usetoast";
import RadioButton from "primevue/radiobutton";
import {getUserNamespaces} from "../../service/filemanage";
const toast=useToast();
const userStore = useUserStore();
const copy=async (content:string)=>{
  await navigator.clipboard.writeText(content);
  toast.add({severity:'success',summary:'复制成功',detail:'复制成功',life:3000});
}
onMounted(async ()=>{
  const user=await FetchUserInfo();
  userStore.setUser(user);
})
</script>

<template>
  <Toast></Toast>
  <Dialog modal v-model:visible="userStore.dialog.account.visible">
    <template #header>
      <h3 class="mr-5">账号管理</h3>

    </template>
    <div class="flex border-t border-solid border-t-gray-300 p-2 ">
      <div class="mr-3 mt-3 ">
        <div class="m-1 p-2 rounded hover:bg-blue-500 hover:text-gray-50">个人信息</div>
        <div class="m-1 p-2 rounded hover:bg-blue-500 hover:text-gray-50">账号设置</div>
      </div>
      <div class="bg-[rgb(235,235,235)] p-3 rounded-2xl shadow">
        <div class="border-b border-solid border-b-gray-400 flex p-2">
          <div class="w-20 font-bold">
            用户ID:
          </div>
          <div  class="w-48">{{userStore.user.id}}</div>
          <div @click="copy(userStore.user.id.toString())" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-copy "></i>
            <span class="font-thin text-sm ml-1">复制</span>
          </div>
        </div >
        <div class="border-b border-solid border-b-gray-400 flex p-2">
          <div class="w-20 font-bold">
            用户名:
          </div>
          <div class="w-48">{{userStore.user.username}}</div>
          <div @click="copy(userStore.user.username)" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-copy "></i>
            <span class="font-thin text-sm ml-1">复制</span>
          </div>
        </div>
        <div class="border-b border-solid border-b-gray-400 flex p-2">
          <div class="w-20 font-bold">
            性别:
          </div>
          <div class="flex w-48 align-items-center" v-if="userStore.input.gender.visible">
            <div class="flex mr-2 align-items-center">
              <RadioButton v-model="userStore.user.gender" inputId="ingredient1" name="gender" value="男" />
              <label for="ingredient1" class="ml-2">男</label>
            </div>
            <div class="flex align-items-center">
              <RadioButton v-model="userStore.user.gender" inputId="ingredient2" name="gender" value="女" />
              <label for="ingredient2" class="ml-2">女</label>
            </div>
          </div>

          <div v-if="!userStore.input.gender.visible" class="w-48">{{userStore.user.gender}}</div>
          <div v-if="!userStore.input.gender.visible" @click="userStore.input.gender.visible=true" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-user-edit "></i>
            <span class="font-thin text-sm ml-1">修改</span>
          </div>
          <div @click="modifyUserInfo(toast)" v-if="userStore.input.gender.visible" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-check "></i>
            <span  class="font-thin text-sm ml-1">确定</span>
          </div>
        </div>
        <div class="border-b border-solid border-b-gray-400 flex p-2">
          <div class="w-20 font-bold">
            邮箱:
          </div>
          <IconField v-if="userStore.input.email.visible" iconPosition="right">
            <InputIcon>
              <i @click="userStore.input.email.visible=false" class="pi pi-times cursor-pointer" />
            </InputIcon>
            <InputText :pt="{
            root: { class: 'border-teal-400 h-6 w-48' }
          }" v-model="userStore.user.email" ></InputText>
          </IconField>
          <div v-if="!userStore.input.email.visible" class="w-48">{{userStore.user.email}}</div>
          <div @click="userStore.input.email.visible=true" v-if="!userStore.input.email.visible" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-user-edit "></i>
            <span  class="font-thin text-sm ml-1">修改</span>
          </div>
          <div @click="modifyUserInfo(toast)" v-if="userStore.input.email.visible" class="flex-auto w-14 ml-3  hover:text-blue-500 cursor-pointer">
            <i class="pi pi-check "></i>
            <span  class="font-thin text-sm ml-1">确定</span>
          </div>
        </div>
      </div>
    </div>
  </Dialog>
</template>

<style scoped>

</style>