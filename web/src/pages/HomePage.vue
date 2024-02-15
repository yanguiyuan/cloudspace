<template>
    <div id="bg">
        <div class="mask">
            <div @click="showMenu" class="menuList">
                <svg t="1671208478479" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="18985" width="32" height="32">
                    <path
                        d="M550.4 74.666667c25.6 0 46.933333 19.2 53.333333 44.8l14.933334 85.333333 38.4 17.066667L727.466667 170.666667c19.2-14.933333 46.933333-12.8 66.133333 4.266666l2.133333 2.133334 53.333334 53.333333c19.2 19.2 21.333333 46.933333 6.4 68.266667l-49.066667 70.4 17.066667 38.4 85.333333 14.933333c23.466667 4.266667 42.666667 25.6 44.8 49.066667v78.933333c0 25.6-19.2 46.933333-44.8 53.333333l-85.333333 14.933334-17.066667 38.4 49.066667 70.4c14.933333 19.2 12.8 46.933333-4.266667 66.133333l-2.133333 2.133333-53.333334 53.333334c-19.2 19.2-46.933333 21.333333-68.266666 6.4l-70.4-49.066667-38.4 17.066667-14.933334 85.333333c-4.266667 23.466667-25.6 42.666667-49.066666 44.8h-78.933334c-25.6 0-46.933333-19.2-53.333333-44.8l-14.933333-85.333333-38.4-17.066667-72.533334 46.933333c-19.2 14.933333-46.933333 12.8-66.133333-4.266666l-2.133333-2.133334-53.333334-53.333333c-19.2-19.2-21.333333-46.933333-6.4-68.266667l49.066667-70.4-17.066667-38.4-85.333333-14.933333c-23.466667-4.266667-42.666667-25.6-44.8-49.066667v-78.933333c0-25.6 19.2-46.933333 44.8-53.333333l85.333333-14.933334 17.066667-38.4L170.666667 296.533333c-14.933333-19.2-12.8-46.933333 2.133333-64l2.133333-2.133333 53.333334-53.333333c19.2-19.2 46.933333-21.333333 68.266666-6.4l70.4 49.066666 38.4-17.066666 14.933334-85.333334c4.266667-23.466667 25.6-42.666667 49.066666-44.8H550.4z m-38.4 320c-64 0-117.333333 53.333333-117.333333 117.333333s53.333333 117.333333 117.333333 117.333333 117.333333-53.333333 117.333333-117.333333-53.333333-117.333333-117.333333-117.333333z"
                        fill="#666666" p-id="18986"></path>
                </svg>
            </div>
            <div v-bind:style="{ display: menuTipVisible }" class="settingTooltip">
                <div @click="userStore.showAccountDialog" class="tipMenuItem">账号管理</div>
                <div class="tipMenuItem">
                    <RouterLink  @click="logout" to="/login">
                        退出登录
                    </RouterLink>
                </div>
            </div>
            <CurrentTime></CurrentTime>
            <RouterView></RouterView>
        </div>
      <AccountManagementDialog></AccountManagementDialog>
    </div>
</template>
<script setup lang="ts">
import { ref, onUnmounted, onMounted, getCurrentInstance } from 'vue'
import {useUserStore} from "../store/user";
import AccountManagementDialog from "../components/user/AccountManagementDialog.vue";
import {logout} from "../service/user";
const menuTipVisible = ref<string>("none")
const showMenu = function () {
    menuTipVisible.value = "block"
}
const userStore=useUserStore();
const handleClickOutside = function (e: MouseEvent) {

    const target = e.target as HTMLElement;
    if (!target.closest('.menuList') && menuTipVisible.value == "block") {
        menuTipVisible.value = "none"
    }

}
onMounted(() => {
    document.addEventListener("click", handleClickOutside)
})
onUnmounted(() => {
    document.removeEventListener("click", handleClickOutside)
})
</script>
<style>
#bg {
    position: fixed;
    background: url('https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/RE4wwtC?ver=b327');
    bottom: 0;
    top: 0;
    right: 0;
    left: 0;
}

.mask {
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    //backdrop-filter: blur(10px);
}

.settingTooltip {
    position: absolute;
    background-color: white;
    left: 5px;
    padding: 5px;
    border-radius: 5px;
}

.tipMenuItem {
    padding: 10px 15px;
    border-radius: 5px;
    transition: 0.25s;
    font-size: 12px;
}

.tipMenuItem:hover {
    background-color: rgb(205, 205, 205);
}
</style>