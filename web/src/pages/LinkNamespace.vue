<script setup lang="ts">
import {ref} from "vue";

const visible=ref<boolean>(true)
</script>

<template>
<Dialog v-model:visible="visible">
  你好
</Dialog>
</template>

<style scoped>

</style>