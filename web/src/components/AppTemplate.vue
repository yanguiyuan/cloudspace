<template>
    <div :class="{ card: !canExpand, expandCard: canExpand }">
        <div :class="{ sideCard: !canExpand, sideCardExpand: canExpand }">
            <div class="menuContainer flex justify-center align-center flex-col">
                <div @click="it.action" v-for="it in optionItems" class="menuItem flex justify-center">
                    <RouterLink v-if="it.enableRouterLink" :to="it.to">
                       <div v-html="getIcon(it.icon)"></div>
                    </RouterLink>
                    <div v-html="getIcon(it.icon)" v-if="!it.enableRouterLink"></div>
                    <div class="tooltip">{{it.tooltip}}</div>
                </div>
              <div @click="it.action" v-for="it in props.options" class="menuItem flex justify-center">
                <RouterLink v-if="it.enableRouterLink" :to="it.to">
                  <div v-html="getIcon(it.icon)"></div>
                </RouterLink>
                <div v-html="getIcon(it.icon)" v-if="!it.enableRouterLink"></div>
                <div class="tooltip">{{it.tooltip}}</div>
              </div>
            </div>
        </div>
        <div  class="p-3">
          <slot name="header"></slot>
        </div>
        <div class="relative flex-auto  overflow-auto" :style="expandPadding">
            <slot name="content"></slot>
        </div>
    </div>
</template>
<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {AssetsIconSvgService} from "../assets/assets";
const props=defineProps({
  options:{
    type:Array<OptionItem>,
    default:[]
  },
})
interface OptionItem{
    tooltip:string
    to:string
    icon:string
    enableRouterLink:boolean
    action:()=>void
}
const expandPadding = ref<string>("padding: 0;transition: 500ms;")
const canExpand = ref<Boolean>(false)
const iconAsset = AssetsIconSvgService.getInstance()
const getIcon = (icon:string)=>{
    return iconAsset.getIcon(icon)
}
const optionItems:OptionItem[]=[
    {
        tooltip:"返回",
        to:"/",
        icon:"back",
        enableRouterLink:true,
        action:()=>{
            console.log("返回")
        }
    },
    {
        tooltip:"全屏",
        to:"/",
        icon:"fullscreen",
        enableRouterLink:false,
        action:()=>{
            canExpand.value = true;
            expandPadding.value="padding: 32px 64px 32px 64px;transition: 500ms;";
        }
    },
    {
        tooltip:"取消全屏",
        to:"/",
        icon:"fullscreen-close",
        enableRouterLink:false,
        action:()=>{
          canExpand.value = false;
          expandPadding.value="padding: 0;transition: 500ms;";
        }
    }
]



</script>
  
<style scoped>
input {
    border-radius: 8px;
}
.card {
    position: absolute;
    bottom: 80px;
    top: 120px;
    left: 20%;
    right: 20%;
    display:flex;
    flex-direction: column;
    border-radius: 15px;
    background-color: rgb(235, 235, 235);
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
    transition: 500ms;
}

.sideCard {
    position: absolute;
    width: 90px;
    left: -100px;
    top: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.1);
    border-radius: 15px;
    backdrop-filter: blur(30px);
    transition: 0.15s;
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
    transition: 500ms;
}
.sideCardExpand{
    position: absolute;
    width: 48px;
    left: 0px;
    z-index: 999;
    top: 25%;
    bottom: 25%;
    background-color: rgba(0, 0, 0, 0.2);
    border-radius: 15px 45px 45px 15px;
    backdrop-filter: blur(30px);
    transition: 0.15s;
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
    transition: 500ms;
}


.expandCard {
    position: absolute;
    display:flex;
    flex-direction: column;
    bottom: 0px;
    top: 0px;
    left: 0%;
    right: 0%;
    border-radius: 15px;
    background-color: rgb(235, 235, 235);
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
    transition: 500ms;
}


.content {
    position: absolute;
    bottom: 0;
    right: 0;
    left: 0;
    top: 0;
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);

    border-radius: 15px;
    background-color: rgb(235, 235, 235);
    overflow: auto;
    padding: 20px;
}

.menuContainer {
    margin: 10px;
    text-align: center;
}



.menuItem {
    margin: 10px 0 10px 0;
    position: relative;
}

.menuItem:hover {
    margin: 10px 0 10px 0;
    transform: 0.25s;
    transform: translateX(10px);
}

.menuItem:hover .tooltip {
    display: block;
}

.tooltip {
    display: none;
    position: absolute;
    top: -15px;
    /* left: 55px; */
    right: 90px;
    background-color: #ccc;
    padding: 10px;
    border-radius: 5px;
    transition: 0.25s;
}</style>
  