<template>
    <div class="btn-container">
        <button :style="{height:props.height,width:props.width}"  class="btn">
            <slot></slot>
        </button>
    </div>
    
</template>
<script setup lang="ts">
import { defineProps,withDefaults } from 'vue';
const props=withDefaults(defineProps<{height:string,
    width:string
}>(),{height:"43px",width:"160px"})
</script>
<style scoped>
.btn {
user-select: none;
box-sizing: border-box;
touch-action: pan-y;
outline: 0;
border: none;
position: relative;
width: 100%;
height: 100%;
padding: 0 20px;
color: #fff;
background-color: transparent;
font-size: 14px;
font-family: var(--fonts-light);
text-align: center;
cursor: pointer;
}
.btn-container:hover {
    background-color: rgba(0, 0, 255, .50);
}
.btn-container {

font-family: var(--fonts-light);
cursor: pointer;
-webkit-tap-highlight-color: transparent;
user-select: none;
box-sizing: border-box;
outline: 0;
touch-action: pan-y;
position: absolute;
left: 50%;
transform: translateX(-50%);
/* width: 160px; */
max-width: 80%;
/* height: 43px; */
border-radius: 30px;
color: #fff;
background-color: #3b4354;
box-shadow: rgba(0, 0, 0, .2) 0 0 10px;
backdrop-filter: blur(10px);
overflow: hidden;
transition: color .25s, background-color .25s, box-shadow .25s, left .25s, opacity .25s, top .25s, width .25s;
/* top: 465px; */
opacity: 1;
}
</style>