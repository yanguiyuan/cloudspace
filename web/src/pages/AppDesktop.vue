<template>
    <div class="navContainer">
        <div  v-for="it in appList.filter(it => it.enable)"  class="navItem">
            <RouterLink :to="it.link">
                <div class="navIcon" v-html="getIcon(it.icon)">
                </div>
                <div class="navTitle">
                    {{it.title}}
                </div>
            </RouterLink>
        </div>
    </div>
</template>
<script setup lang="ts">
import {AssetsIconSvgService} from "../assets/assets";
import {reactive} from "vue";
interface AppProp {
  icon:string,
  title:string,
  link:string,
  enable:boolean
}
const iconAsset = AssetsIconSvgService.getInstance()
const getIcon = (icon:string)=>{
  return iconAsset.getIcon(icon)
}
const appList =reactive<Array<AppProp>>( [
  {
    icon: 'file-management',
    title: '文件管理',
    link: '/fileManagement',
    enable: true
  },
  {
    icon: 'test',
    title: '测试',
    link: '/test',
    enable: true
  },
  {
    icon: 'js',
    title: 'NT竞赛',
    link: '/js',
    enable: false
  },
  {
    icon: 'chatgpt',
    title: 'ChatGPT',
    link: '/chatGPT',
    enable: false
  },
  {
    icon: 'map',
    title: '百度地图',
    link: '/map',
    enable: false
  },
  {
    icon: 'echarts',
    title: 'Echarts',
    link: '/echarts',
    enable: false
  },
  {
    icon: '3d-playground',
    title: '3D-Playground',
    link: '/3DPlayground',
    enable: true
  },
  {
    icon: 'image',
    title: '图片生成',
    link: '/image',
    enable: false
  },
  {
    icon:'app-admin',
    title:'用户管理',
    link:'/admin',
    enable:true
  },
  {
    icon: 'add-app',
    title: '添加应用',
    link: '/',
    enable: true
  },
])
</script>
<style>
.navContainer {
    position: absolute;
    left: 20%;
    margin-top: 200px;
    right: 20%;
    height: calc(100vh - 330px);
    overflow: hidden;
    transition: .25s;
}

.navItem {
    position: relative;
    margin-top: 10px;
    margin-bottom: 40px;
    display: inline-block;
    width: 80px;
    height: 80px;
    cursor: pointer;
    margin: 25px;
    text-align: center;
    border-radius: 15px;
    background-color: rgb(235, 235, 235) !important;
}

.navTitle {
    bottom: -25px;
    color: #fff;
    font-size: 12px;
    left: 5%;
    overflow: hidden;
    position: absolute;
    text-align: center;
    text-overflow: ellipsis;
    transition: .25s;
    white-space: nowrap;
    width: 90%;
}

.navIcon {
    margin-top: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>