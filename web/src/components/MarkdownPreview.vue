<template>
<div>     
  <v-md-preview :text="md"></v-md-preview>
</div> 
</template>
<script setup lang="ts">
import { ref } from 'vue';
const md=ref<string>("# 你好啊")

</script>