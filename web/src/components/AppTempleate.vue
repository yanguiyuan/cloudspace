<template>
    <div :class="{ card: !canExpand, expandCard: canExpand }">
        <div :class="{ sideCard: !canExpand, sideCardExpand: canExpand }">
            <div class="menuContainer">
                <div class="menuItem">
                    <RouterLink to="/">
                        <svg t="1671204857123" class="icon" viewBox="0 0 1024 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="6487" width="32" height="32">
                            <path
                                d="M471.893333 149.333333a42.666667 42.666667 0 0 0-73.258666-29.781333l-343.893334 352.981333a42.666667 42.666667 0 0 0-0.768 58.709334l343.893334 372.352a42.666667 42.666667 0 0 0 73.984-28.928v-190.677334c56.917333-5.248 116.821333-1.365333 179.882666 11.989334 65.834667 13.994667 150.528 76.032 253.909334 202.24a42.666667 42.666667 0 0 0 75.477333-31.36c-15.445333-152.32-73.984-281.301333-176.170667-384.853334-92.757333-93.994667-204.373333-146.432-333.098666-156.586666V149.333333z"
                                fill="#e6e6e6" fill-opacity=".85" p-id="6488"></path>
                        </svg>
                    </RouterLink>
                    <div class="tooltip">返回</div>
                </div>
                <div class="menuItem" @click="expand">
                    <svg t="1678623322130" class="icon" viewBox="0 0 1024 1024" version="1.1"
                        xmlns="http://www.w3.org/2000/svg" p-id="4872" width="32" height="32">
                        <path
                            d="M240.8 196l178.4 178.4-45.6 45.6-177.6-179.2-68 68V128h180.8l-68 68z m133.6 408.8L196 783.2 128 715.2V896h180.8l-68-68 178.4-178.4-44.8-44.8zM715.2 128l68 68-178.4 178.4 45.6 45.6 178.4-178.4 68 68V128H715.2z m-65.6 476.8l-45.6 45.6 178.4 178.4-68 68H896V715.2l-68 68-178.4-178.4z"
                            p-id="4873"></path>
                    </svg>

                    <div class="tooltip">全屏</div>
                </div>
                <div class="menuItem" @click="recover">
                    <svg t="1678624002040" class="icon" viewBox="0 0 1024 1024" version="1.1"
                        xmlns="http://www.w3.org/2000/svg" p-id="5081" width="32" height="32">
                        <path
                            d="M298.666667 631.466667H226.133333v-81.066667h217.6v204.8h-85.333333v-68.266667l-128 128L170.666667 759.466667l128-128z m422.4 0l128 128-59.733334 59.733333-128-128v68.266667h-85.333333V554.666667h217.6v81.066666h-72.533333zM298.666667 341.333333L187.733333 230.4 243.2 170.666667l115.2 115.2V217.6h85.333333v204.8H226.133333V341.333333H298.666667z m430.933333 0h64v81.066667h-217.6V217.6h85.333333v72.533333L780.8 170.666667l59.733333 59.733333L729.6 341.333333z"
                            fill="#444444" p-id="5082"></path>
                    </svg>

                    <div class="tooltip">取消全屏</div>
                </div>
                <slot name="option"> </slot>
            </div>
        </div>
        <div class="content">
            <slot name="content"></slot>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const canExpand = ref<Boolean>(false)
const expand = function () {
    canExpand.value = true
}
const recover = function () {
    canExpand.value = false
}



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
    background-color: rgba(0, 0, 0, 0.1);
    border-radius: 15px 45px 45px 15px;
    backdrop-filter: blur(30px);
    transition: 0.15s;
    box-shadow: 0 15px 30px rgb(0 0 0 / 25%);
    transition: 500ms;
}


.expandCard {
    position: absolute;
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
    right: 60px;
    background-color: #ccc;
    padding: 10px;
    border-radius: 5px;
    transition: 0.25s;
}</style>
  