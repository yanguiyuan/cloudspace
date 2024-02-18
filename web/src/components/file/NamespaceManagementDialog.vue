<script setup lang="ts">
import FloatLabel from 'primevue/floatlabel';

import Dropdown from 'primevue/dropdown';
import {onMounted, ref} from "vue";
import {DefaultNamespace, Namespace, useFileStore} from "../../store/file";
import {
  createNamespace,
  generateNamespaceJoinLink,
  getUserNamespaces,
} from "../../service/filemanage";
import Menu from 'primevue/menu';
import Card from "primevue/card";
import {useToast} from "primevue/usetoast";
const toast=useToast();
const selectedNamespace=ref<Namespace>(DefaultNamespace);
const selectedAuth=ref<{ label:string,code:number }>();
const auths=ref([
  {label:"读权限（查看文件列表，浏览文件，下载文件）",code:2},
  {label:"写权限（创建，删除，重命名，上传文件）",code:1},
])
const visible=ref<string>("create")
const title=ref<string>("新建命名空间");
const fileStore = useFileStore()
const namespaceName=ref<string>("");
const copyLink=async ()=>{
    if(!selectedAuth.value){
      toast.add({ severity: 'error', summary: '错误', detail: '请选择授权类型', life: 3000});
      return
    }
    if(selectedNamespace.value.id==0){
      toast.add({ severity: 'error', summary: '错误', detail: '请选择命名空间', life: 3000});
      return
    }
    const res=await generateNamespaceJoinLink(selectedNamespace.value,selectedAuth.value?.code);
    console.log(res);
    toast.add({ severity: 'info', summary: '链接复制成功', detail: res, life: 6000});
}
const items=ref([
  { label: '新建', icon: 'pi pi-plus', command:()=>{visible.value="create";title.value="新建命名空间"} },
  { label: '授权', icon: 'pi pi-paperclip',command:()=>{visible.value="auth";title.value="命名空间授权"} }
])

</script>

<template>
  <Dialog modal v-model:visible="fileStore.dialog.namespace.visible">
    <template #header>
      <h3 class="mr-5">命名管理</h3>
    </template>
    <div class="flex border-t border-solid border-t-gray-300 p-2 ">
      <div class="mr-3 mt-3 ">
        <Menu :model="items"  :pt="{
          action: ({ props, state, context }) => ({
              class: context.focused ? 'bg-blue-500 rounded' : undefined
          })
        }"></Menu>
      </div>
      <div class=" p-3 flex   items-center border-l border-solid">
          <Card>
            <template #title>{{title}}</template>
            <template #content>
              <div v-if="visible=='create'" class="flex   items-center">
                <FloatLabel>
                  <InputText id="namespace" v-model="namespaceName" />
                  <label for="namespace">命名空间名称</label>
                </FloatLabel>
                <Button size="small" @click="createNamespace(namespaceName,toast)" class="ml-2">确定</Button>
              </div>
              <div v-if="visible=='auth'" >
                <div>
                  <div>命名空间:</div>
                  <Dropdown v-model="selectedNamespace" :options="fileStore.namespaces.filter(it=>it.authority==0)" optionLabel="name" placeholder="选择一个命名空间"  />
                </div>
                <div>
                  <div>权限:</div>
                  <Dropdown v-model="selectedAuth" :options="auths" optionLabel="label" placeholder="选择对应的权限"  />
                </div>

                <Button @click="copyLink" size="small" class="mt-2">复制链接</Button>
                <p class="text-sm text-neutral-400">请复制链接后发送给需要邀请其加入该命名空间的用户</p>
              </div>
            </template>
          </Card>

      </div>
    </div>
  </Dialog>
</template>

<style scoped>

</style>