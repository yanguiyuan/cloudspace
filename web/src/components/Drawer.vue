<template>
    <div ref="mask" @click='onClick' v-if="visible"  class="drawerMask">
      <div ref="container" class="drawerContainer">
          <div>{{ title }}</div>
          
          <slot></slot>
        </div>
    </div>
    
</template>
<script setup lang="ts">
import { ref, onUnmounted, onMounted,toRefs} from 'vue'
const props=defineProps({
  visible:{
    type:Boolean,
    default:false
  },
  title:{
    type:String,
    default:""
  }
})
const mask=ref(null)
const container=ref(null)
const emit=defineEmits(['drawerClose'])
const onClick=(e:MouseEvent)=>{
  if(e.target==mask.value&&e.target!=container.value){
    emit('drawerClose')
  }
}
</script>
<style>
.drawerMask{
    position: fixed;
    right: 0;
    top: 0;
    left: 0;
    bottom: 0;
    height: 100%;
    z-index: 999999;
    background-color: rgba(0, 0, 0, 0.5);
}
.drawerContainer{
    position: fixed;
    height: 100%;
    width: 60%;
    right: 0;
    top: 0;
    border-radius: 10px 0 0 10px;
    background-color: #ffff;
    padding: 15px;
    padding-left: 35px;
    overflow: scroll;
    transition: 0.5s;
}
.drawerTitle{
  position: absolute;
  border: 1px solid;
  left: -45px;
  width: 100px;
  height: 50px;
}
.drawerContainer::-webkit-scrollbar {
  width: 5px;
}
.drawerContainer::-webkit-scrollbar-thumb {
  background-color: rgb(175,175,175);
  border-radius: 3px;
}
</style>