<template>
    <AppTempleate>
        <template #content>
            <div class="chat-container">
                <div class="img-container">
                    <img  :src="imgUrl">
                </div>
            </div>
            <div class="question-input-div">
                <textarea v-model="myQuestion"></textarea>
                <svg @click="generateImage" t="1677941452978" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="3434" width="48" height="48">
                    <path
                        d="M411.776 559.957333l203.605333-170.112-269.184 119.978667-148.010666-113.152a21.333333 21.333333 0 0 1 8.64-37.845333l599.850666-123.648a21.333333 21.333333 0 0 1 25.045334 25.834666l-117.546667 492.373334a21.333333 21.333333 0 0 1-33.706667 12.010666l-268.693333-205.44z m-55.893333 170.026667l29.333333-137.301333 114.986667 88.768-112.064 70.997333a21.333333 21.333333 0 0 1-32.277334-22.485333z"
                        fill="#2A2A37" p-id="3435"></path>
                </svg>
            </div>
        </template>
    </AppTempleate>
</template>
<script lang="ts" setup>
import AppTempleate from '../components/AppTempleate.vue';
import {cf as axios} from '../axios/axios';
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
const myQuestion = ref<string>("")
const imgUrl=ref<string>("https://oaidalleapiprodscus.blob.core.windows.net/private/org-wrgyM1Dsxobdhme1NXVgAaeS/user-0QR7SXLgImhxicLK2q2fUGmW/img-HgfKglK1H05uvrHkJOV3zrU5.png?st=2023-03-15T11%3A48%3A47Z&se=2023-03-15T13%3A48%3A47Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-14T21%3A12%3A40Z&ske=2023-03-15T21%3A12%3A40Z&sks=b&skv=2021-08-06&sig=/fyZwsxXXX5lkUzi4xYzAVhawhyvFL3J8NRdxi9Nn3g%3D")
// 替换为你的 API 密钥
const apiKey = "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"

// 发送 HTTP 请求生成图片
const generateImage = async () => {
    const response = await axios.post(
        'https://api.openai.com/v1/images/generations',
        {
            prompt: myQuestion.value,
            n: 1,
            size: '1024x1024',
        },
        {
            headers: {
                Authorization: `Bearer ${apiKey}`,
                'Content-Type': 'application/json',
            },
        }
    );
    const blob = new Blob([response.data], { type: 'image/jpeg' });
    console.log(response.data.data[0])
    imgUrl.value= response.data.data[0].url
}
</script>
<style scoped>
.question-input-div {
    padding: 5px;
    position: absolute;
    bottom: 0;
    right: 0;
    left: 0;
    height: calc(15%);
    border-top: #ccc 1px solid;
    border-radius: 0 0 15px 15px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: rgb(235, 235, 235);

}
.chat-container {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: calc(85%);
    margin-top: 10px;
    overflow-y: auto;
    
}
.img-container{
    position: absolute;
    display: flex;
    justify-content: center;
    /* align-items: center; */
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}
textarea {
    font-family: inherit;
    font-size: inherit;
    color: #333;
    border: 2px solid #ccc;
    border-radius: 20px;
    padding: 10px;
    resize: horizontal;
    min-height: 70px;
    min-width: 300px;

}
</style>