<template>
    <div id="timeContainer">
        <div id="timeText">{{ nowHour }}<span>:</span>{{ nowMinutes }}</div>
    </div>
</template>
<script setup lang="ts">
import { ref,onMounted } from 'vue';
const nowHour = ref<string | number>("00")
const nowMinutes = ref<string | number>("00")
const getNowTime = function () {
    const now = new Date()
    nowHour.value = now.getHours() >= 10 ? now.getHours() : `0${now.getHours()}`
    //0${now.getHours()} 中的 0 和 ${} 符号是字符串模板的一部分。在这个字符串模板中，0 是一个字符，表示这个字符串的开头，
    // 而 ${now.getHours()} 是一个插值表达式，用于将 now.getHours() 的返回值插入到字符串中。
    nowMinutes.value = now.getMinutes() >= 10 ? now.getMinutes() : `0${now.getMinutes()}`
}
onMounted(() => {
    getNowTime()
    setInterval(() => {
        getNowTime()
    }, 1000);
})

</script>

<style>
#timeText {
    font-family: var(--fonts-light);
    -webkit-tap-highlight-color: transparent;
    user-select: none;
    text-align: center;
    cursor: pointer;
    box-sizing: border-box;
    outline: 0;
    touch-action: pan-y;
    animation: delayedFadeIn 2s;
    color: #fff;
    font-size: 36px;
    transition: .25s;
}

#timeContainer {

    font-family: var(--fonts-light);
    -webkit-tap-highlight-color: transparent;
    user-select: none;
    box-sizing: border-box;
    outline: 0;
    touch-action: pan-y;
    position: fixed;
    left: 50%;
    transform: translateX(-50%);
    padding: 10px;
    text-align: center;
    cursor: pointer;
    transition: .25s;
    top: 50px
}
</style>