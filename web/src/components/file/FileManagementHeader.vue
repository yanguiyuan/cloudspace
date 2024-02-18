<script setup lang="ts">

import {
  changeNamespace,
  getUserNamespaces,
  onClickBreadcrumbs
} from "../../service/filemanage";
import {AssetsIconSvgService} from "../../assets/assets";
import {Namespace, useFileStore} from "../../store/file";
import MegaMenu from "primevue/megamenu";
import {onMounted, ref} from "vue";
import {MenuItem} from "primevue/menuitem";
const iconService=AssetsIconSvgService.getInstance();
const fileStore=useFileStore();
const items=ref<MenuItem>([{
  label:"命名空间",
  icon:"pi pi-box",
  items:[
      [
        {
          label:"所有",
          items:[]
        },
      ],
    [
      {
        label:"可写",
        items:[]
      }
    ],
    [
      {
        label:"可读",
        items:[]
      }
    ]
  ]
}])
const addOwnerNamespace=(n:Namespace)=>{
  if(n.authority==0){
    items.value[0].items[0][0].items.push({
      label:n.name,
      command:async ()=>{
        await changeNamespace(n.rootID)
      }
    });
  }else if(n.authority==1){
    items.value[0].items[1][0].items.push({
      label:n.name,
      command:async ()=>{
        await changeNamespace(n.rootID)
      }
    });
  }else if(n.authority==2){
    items.value[0].items[2][0].items.push({
      label:n.name,
      command:async ()=>{
        await changeNamespace(n.rootID)
      }
    });
  }
}
onMounted(async ()=>{
  const namespaces=await getUserNamespaces();
  console.log("namespaces:",namespaces);
  fileStore.setNamespaceList(namespaces)
  for (const namespace of namespaces) {
    addOwnerNamespace(namespace);
  }
  await changeNamespace(fileStore.namespaces[0].rootID);
});
</script>

<template>
  <div class="flex items-center">
    <div class="fileIcon" v-html="iconService.getIcon('file-management-32')"> </div>
    <div class="ml-1">文件管理</div>
  </div>
  <div class="border-b border-solid border-gray-300 p-2">
    <MegaMenu  class="inline-block ml-[-1.0rem] mr-2 bg-transparent border-0" :model="items" />
    <div class="inline-block" v-for="(it, key) in fileStore.breadcrumbs">
      <div v-if="key!=0" class="inline-block divider">/</div>
      <div @click="onClickBreadcrumbs(it,key)" v-if="key!= fileStore.breadcrumbs.length - 1"   class="inline-block filePoint">{{it.fileName }}</div>
      <strong v-if="key == fileStore.breadcrumbs.length - 1"  id="currentDirectory">{{ it.fileName }}</strong>
    </div>

  </div>
</template>

<style scoped>
.divider {
  color: rgba(0, 0, 0, 0.4);
  margin: 0 0.2rem 0;
}
.filePoint {
  color: #005980;
  cursor: pointer;
}

.filePoint:hover {
  color: #537c8d;
}
</style>