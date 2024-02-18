<template>
    <div class="bg">
        <Toast />
        <div class="login-container" :class="{ mask: isMask }">
            <div class="coner-ico"></div>
            <CurrentTime></CurrentTime>
            <div class="login-form" @mouseover="fuzzyBackgroud" @mouseleave="backgroudRestoration">
                <div class="tab-container">
                    <div class="tab-item-container-1">
                        <div class="tab-item" :class="{ active: isLoginActive }"><span @click="loginActive">登录</span>
                        </div>
                    </div>
                    <div class="tab-item-container-2">
                        <div class="tab-item" :class="{ active: isRegisterActive }"><span
                                @click="registerActive">注册</span></div>
                    </div>
                </div>
                <div style="top:100px" class="login-input-container">

                    <input v-model="user" class="login-input" type="text" placeholder="用户名">
                </div>
                <div :class="{ pwdinputlogin: isLoginActive, pwdinputregister: isRegisterActive }"
                    class="login-input-container">
                    <input v-model="pwd" class="login-input" type="password" placeholder="密码">
                </div>
                <div style="top:250px" :class="{ comfirminput: isLoginActive }" class="login-input-container">
                    <input v-model="comfirmPwd" class="login-input" type="password" placeholder="确认密码">
                </div>
                <div style="top:325px" class="login-btn-container">
                    <button @click="loginOrRegister" class="login-btn">{{ btnmessage }}</button>
                </div>
            </div>

        </div>

    </div>

</template>

<script setup lang="ts">
import axios from '../axios/axios'
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import CurrentTime from '../components/CurrentTime.vue';
import { useToast } from "primevue/usetoast";
import {useUserStore} from "../store/user";
import {getUrlParams} from "../service/filemanage";
const toast = useToast();
const userStore=useUserStore();
const btnmessage = ref<string>("登录")
const user = ref<string>("")
const pwd = ref<string>("")
const comfirmPwd = ref<string>("")
const isMask = ref<boolean>(false)
const isLoginActive = ref<boolean>(true)
const isRegisterActive = ref<boolean>(false)
const router = useRouter()
console.log("linkNamespace",userStore.linkNamespace)
let loginOrRegister = function () {
    if (btnmessage.value == "登录") {
        axios.post("/login", {
            "username": user.value,
            "password": pwd.value
        }).then((res) => {
            localStorage.setItem("token", JSON.stringify(res.data.token))
            localStorage.setItem("user", JSON.stringify(res.data.user))
            const params=getUrlParams(window.location.href);
            axios.defaults.headers.common["Authorization"]="Bearer "+res.data.token

            if(userStore.linkNamespace){
              userStore.user.id=0
              router.push("/link")
              return;
            }
            router.push("/");
        }).catch((err) => {
            console.log(err)
            ElMessage.error(err.response.data.message)
        })
    } else {
        if (pwd.value == comfirmPwd.value) {
            axios.post("/register", {
                "username": user.value,
                "password": pwd.value
            }).then((res) => {
                if(res.data.code==0){
                  ElMessage.success("注册成功")
                  loginActive();
                  return
                }
                if(res.data.code==1002){
                  toast.add({ severity: 'error', summary: '错误', detail: "用户名已经存在", life: 3000 })
                  return;
                }
                toast.add({ severity: 'error', summary: '错误', detail: res.data.msg, life: 3000 })

            }).catch((err)=>{
                console.log(err.response)
                ElMessage.error(err.response.data.error)
            })
        }else{
            ElMessage.error("密码不一致")
        }

    }
    console.log(user.value, pwd.value)
}

const registerActive = function () {
    isRegisterActive.value = true
    isLoginActive.value = false
    btnmessage.value = "注册"
}
const loginActive = function () {
    isRegisterActive.value = false
    isLoginActive.value = true
    btnmessage.value = "登录"
}
const fuzzyBackgroud = function () {
    isMask.value = true
}
const backgroudRestoration = function () {
    isMask.value = false
}

</script>
<style scoped>
.pwdinputregister {
    top: 175px;
}

.pwdinputlogin {
    top: 200px
}

.comfirminput {
    display: none;
}

.coner-ico {
    position: absolute;
    /* border: 1px solid; */
    height: 30px;
    width: 30px;
    right: 10px;
    top: 10px;
}

.tab-item-container-1 {
    position: fixed;
    margin-right: 80px;
    /* margin-bottom: 30px; */
}

.tab-item-container-2 {
    position: fixed;
    margin-left: 80px;
    /* margin-bottom: 30px; */
}

.tab-container {
    position: fixed;
    display: flex;
    align-items: center;
    justify-content: center;
    /* border: 1px solid; */
    width: 380px;
    height: 20px;
    border-bottom: 1px solid #f6f6f6;
    margin: 30px;
}

.tab-item.active:before {
    content: "";
    width: 20px;
    height: 3px;
    border-radius: 2px;
    background-color: #3b4354;
    ;
    left: 0;
    right: 0;
    margin: auto;
    bottom: 0;


}

.tab-item {
    display: inline-block;
    font-size: 18px;
    font-weight: 700;
    margin: 3px 35px;
    color: rgba(0, 0, 0, 0.5);
    cursor: pointer;

}

.tab-item.active {
    color: rgba(0, 0, 0, 1);
}

.login-form {
    height: 400px;
    width: 380px;
    /* border: 1px solid; */
    background-color: rgba(255, 255, 255, .99);
    ;
    border-radius: 8px;
    backdrop-filter: blur(20px);
    /* left: 40%; */
    display: flex;
    justify-content: center;

}

.mask {
    backdrop-filter: blur(10px);
    border: #3b4354 1px solid;
}

.login-container {
    display: flex;
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    box-sizing: border-box;
}

.bg {
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: #F2F3F5;
    background: url("https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/RE4wwtC?ver=b327") no-repeat;
    /* background: url("https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/RE4wwtC?ver=b327") no-repeat; */
    background-size: cover;
}

.login-input {
    user-select: none;
    box-sizing: border-box;
    touch-action: pan-y;
    outline: 0;
    border: none;
    width: 100%;
    height: 100%;
    padding: 0 20px;
    color: inherit;
    background-color: rgba(0, 0, 0, .05);
    font-size: 14px;
    font-family: var(--fonts-light);
    text-align: center;
}

.login-input-container {
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
    width: 230px;
    max-width: 80%;
    height: 43px;
    border-radius: 30px;
    color: rgba(0, 0, 0, 1);
    background-color: rgba(255, 255, 255, .25);
    box-shadow: rgba(0, 0, 0, .2) 0 0 10px;
    backdrop-filter: blur(10px);
    overflow: hidden;
    transition: color .25s, background-color .25s, box-shadow .25s, left .25s, opacity .25s, top .25s, width .25s;
    opacity: 1;
    border: #3b4354 1px solid;

}

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

.login-btn {

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

.login-btn-container:hover {
    background-color: rgba(0, 0, 255, .50);
}

.login-btn-container {

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
    width: 230px;
    max-width: 80%;
    height: 43px;
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