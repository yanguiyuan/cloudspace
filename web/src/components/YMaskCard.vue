<template>
    <div v-if="props.visible" class="info-card-mask">
        <div class="info-card">
            <div class="info-card-close">
                <svg @click="onClose" t="1678186892453" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="2754" width="16" height="16">
                    <path
                        d="M512 456.310154L94.247385 38.557538a39.542154 39.542154 0 0 0-55.689847 0 39.266462 39.266462 0 0 0 0 55.689847L456.310154 512 38.557538 929.752615a39.542154 39.542154 0 0 0 0 55.689847 39.266462 39.266462 0 0 0 55.689847 0L512 567.689846l417.752615 417.752616c15.163077 15.163077 40.290462 15.36 55.689847 0a39.266462 39.266462 0 0 0 0-55.689847L567.689846 512 985.442462 94.247385a39.542154 39.542154 0 0 0 0-55.689847 39.266462 39.266462 0 0 0-55.689847 0L512 456.310154z"
                        fill="#999999" p-id="2755"></path>
                </svg>
            </div>
            <div class="info-card-content">
                <slot></slot>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
// import { emit } from 'process';
import { defineProps,withDefaults,defineEmits } from 'vue';
interface Props{
    visible:boolean,
}
const props=withDefaults(defineProps<Props>(),{visible:false})
const emit=defineEmits<{(event:"onClose") :void}>()
const onClose=function(){
    emit("onClose");
}
</script>
<style scoped>
.info-card {
  position: fixed;
  bottom: 40%;
  top: 40%;
  left: 30%;
  right: 30%;
  border-radius: 15px;
  background-color: rgb(235, 235, 235);
  box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
  padding: 10px;
}

.info-card-mask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 9999;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}

.info-card-content {
  margin: 10px;
  position: absolute;
  top: 15px;
  left: 0;
  right: 0;
  bottom: 0;
}

.info-card-close {
  display: flex;
  justify-content: flex-end;
  cursor: pointer;
  z-index: 999;
}</style>