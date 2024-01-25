<template>
    <AppTempleate>
        <template #option>
            <div class="menuItem">
                    <div @click="save">
                        <svg t="1678518979198" class="icon" viewBox="0 0 1024 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="2770" width="32" height="32">
                            <path
                                d="M704 128H192c-35.2 0-64 28.8-64 64v640c0 35.2 28.8 64 64 64h640c35.2 0 64-28.8 64-64V320L704 128zM256 256h320v128H256V256z m256 512c-70.4 0-128-57.6-128-128s57.6-128 128-128 128 57.6 128 128-57.6 128-128 128z"
                                p-id="2771"></path>
                        </svg>
                    </div>
                    <div class="tooltip">保存到文件空间</div>
                </div>
                <div class="menuItem">
                    <div @click="restart">
                        <svg t="1677948435025" class="icon" viewBox="0 0 1025 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="11550" width="32" height="32">
                            <path
                                d="M707.584 389.312c-12.928-6.08-18.88-14.272-17.728-24.576s5.504-18.432 13.12-24.576c2.304-1.536 9.728-7.424 22.272-17.664 12.544-10.304 28.032-22.656 46.272-37.12-32-36.544-70.144-65.28-114.304-86.208-44.096-20.864-92.096-31.36-143.936-31.36-48 0-92.992 8.96-134.912 26.816-41.92 17.92-78.528 42.432-109.76 73.664-31.232 31.168-55.808 67.712-73.728 109.568s-26.88 86.784-26.88 134.72c0 47.168 8.96 91.712 26.88 133.568s42.496 78.4 73.728 109.568c31.232 31.232 67.84 55.744 109.76 73.6s86.848 26.816 134.912 26.816c39.616 0 77.184-6.272 112.576-18.816 35.456-12.544 67.84-30.272 97.152-53.056 29.312-22.848 54.464-49.664 75.456-80.512 20.928-30.848 36.416-64.512 46.272-101.056 5.312-17.472 15.232-31.936 29.696-43.392 14.464-11.392 31.616-17.152 51.456-17.152 23.616 0 43.648 8.384 60.032 25.152 16.384 16.704 24.576 36.544 24.576 59.328 0 10.624-1.92 20.928-5.696 30.848-16 53.248-39.616 102.336-70.848 147.264s-68.416 83.52-111.488 115.84c-43.072 32.32-90.88 57.664-143.488 75.904-52.544 18.368-107.776 27.52-165.696 27.52-70.848 0-137.536-13.312-200.064-39.936-62.464-26.624-116.8-63.168-162.88-109.568s-82.688-100.672-109.76-162.688c-27.072-62.016-40.576-128.448-40.576-199.232s13.504-137.344 40.576-199.808c27.072-62.4 63.616-116.608 109.76-162.688 46.08-46.016 100.416-82.56 162.88-109.568s129.152-40.512 200.064-40.512c78.528 0 151.296 16.192 218.368 48.512s124.224 76.288 171.456 131.84c17.536-12.928 32.192-24.192 44.032-33.664 11.84-9.536 19.264-15.424 22.272-17.728 4.608-3.776 9.536-6.656 14.848-8.576s10.304-1.92 14.848 0c4.544 1.92 8.576 5.888 12.032 11.968s5.888 15.232 7.424 27.392c0.768 6.848 1.536 18.432 2.304 34.816s1.536 35.392 2.304 57.088c0.768 21.696 1.152 44.16 1.152 67.328 0 23.232-0.384 45.12-1.152 65.664-0.768 25.856-6.848 44.544-18.304 55.936-6.144 6.144-15.296 9.6-27.456 10.368-12.224 0.768-25.152-0.384-38.848-3.456-20.608-4.544-43.264-9.472-68.032-14.848-24.768-5.312-48.384-10.816-70.848-16.576-22.464-5.696-42.304-11.072-59.456-16-17.152-4.864-28.032-8.512-32.64-10.752l0 0 0 0z"
                                fill="#444444" p-id="11551"></path>
                        </svg>
                    </div>
                    <div class="tooltip">清空会话</div>
                </div>
        </template>
        <template #content>
            <div class="chat-container">
                <div class="chat-div" v-for="record in records">
                    <div v-if="record.role == 'assistant'" class="answer-div">
                        <svg t="1677943933247" class="icon avator" viewBox="0 0 1024 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="7539" width="32" height="32">
                            <path
                                d="M512 64c259.2 0 469.333333 200.576 469.333333 448s-210.133333 448-469.333333 448a484.48 484.48 0 0 1-232.725333-58.88l-116.394667 50.645333a42.666667 42.666667 0 0 1-58.517333-49.002666l29.76-125.013334C76.629333 703.402667 42.666667 611.477333 42.666667 512 42.666667 264.576 252.8 64 512 64z m0 64C287.488 128 106.666667 300.586667 106.666667 512c0 79.573333 25.557333 155.434667 72.554666 219.285333l5.525334 7.317334 18.709333 24.192-26.965333 113.237333 105.984-46.08 27.477333 15.018667C370.858667 878.229333 439.978667 896 512 896c224.512 0 405.333333-172.586667 405.333333-384S736.512 128 512 128z m-157.696 341.333333a42.666667 42.666667 0 1 1 0 85.333334 42.666667 42.666667 0 0 1 0-85.333334z m159.018667 0a42.666667 42.666667 0 1 1 0 85.333334 42.666667 42.666667 0 0 1 0-85.333334z m158.997333 0a42.666667 42.666667 0 1 1 0 85.333334 42.666667 42.666667 0 0 1 0-85.333334z"
                                fill="#333333" p-id="7540"></path>
                        </svg>
                        <!-- <div class="answer-content">{{ record.content }}</div> -->
                        <div class="answer-content"><v-md-preview :text="record.content"></v-md-preview></div>
                    </div>

                    <div v-if="record.role == 'user'" class="question-div">

                        <svg t="1677944666743" class="icon avator" viewBox="0 0 1024 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="9414" width="32" height="32">
                            <path d="M336 421m-48 0a48 48 0 1 0 96 0 48 48 0 1 0-96 0Z" p-id="9415" fill="#1296db"></path>
                            <path d="M688 421m-48 0a48 48 0 1 0 96 0 48 48 0 1 0-96 0Z" p-id="9416" fill="#1296db"></path>
                            <path
                                d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64z m263 711c-34.2 34.2-74 61-118.3 79.8C611 874.2 562.3 884 512 884c-50.3 0-99-9.8-144.8-29.2-44.3-18.7-84.1-45.6-118.3-79.8-34.2-34.2-61-74-79.8-118.3C149.8 611 140 562.3 140 512s9.8-99 29.2-144.8c18.7-44.3 45.6-84.1 79.8-118.3 34.2-34.2 74-61 118.3-79.8C413 149.8 461.7 140 512 140c50.3 0 99 9.8 144.8 29.2 44.3 18.7 84.1 45.6 118.3 79.8 34.2 34.2 61 74 79.8 118.3C874.2 413 884 461.7 884 512s-9.8 99-29.2 144.8c-18.7 44.3-45.6 84.1-79.8 118.2z"
                                p-id="9417" fill="#1296db"></path>
                            <path
                                d="M664 565H360c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8z"
                                p-id="9418" fill="#1296db"></path>
                        </svg>
                        <div class="question-content">
                            {{ record.content }}
                        </div>
                    </div>

                </div>
            </div>
            <div class="question-input-div">
                <textarea v-model="myQuestion"></textarea>
                <svg @click="chat" t="1677941452978" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="3434" width="48" height="48">
                    <path
                        d="M411.776 559.957333l203.605333-170.112-269.184 119.978667-148.010666-113.152a21.333333 21.333333 0 0 1 8.64-37.845333l599.850666-123.648a21.333333 21.333333 0 0 1 25.045334 25.834666l-117.546667 492.373334a21.333333 21.333333 0 0 1-33.706667 12.010666l-268.693333-205.44z m-55.893333 170.026667l29.333333-137.301333 114.986667 88.768-112.064 70.997333a21.333333 21.333333 0 0 1-32.277334-22.485333z"
                        fill="#2A2A37" p-id="3435"></path>
                </svg>
            </div>
        </template>
    </AppTempleate>
</template>
<script setup lang="ts">
import { ElMessage } from 'element-plus';
import AppTemplate from '../components/AppTemplate.vue';
import { ref, onUnmounted, onMounted, reactive, getCurrentInstance } from 'vue'
import axios from 'axios';

interface Record {
    role: string
    content: string
}
interface UserData {
    token: string
    username: string
}
const records = reactive<Record[]>([])
const userData=reactive<UserData>({
    token:"",
    username:""
})
const myQuestion = ref<string>("")
const api_key = "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"

const getContent=function(){
    let str=""
    records.forEach(e => {
        if(e.role=="user"){
            str+="## "+e.content+"##\n"
        }else{
            str+=e.content+"\n"
        }

    });
    return str
}
const restart = function () {
    records.length = 0
}
const chat = function () {
    records.push({role: "user", content: myQuestion.value})
    const url = 'https://api.openai.com/v1/chat/completions';
    const params = {
        messages: records,
        max_tokens: 1000,
        temperature: 0.8,
        model: "gpt-3.5-turbo",
        top_p: 0.9
    };
    axios.post(url, params, {
        headers: {
            'Authorization': 'Bearer ' + api_key,
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            records.push({role: "assistant", content: response.data.choices[0].message.content})
            console.log(response.data.choices[0].message.content);
            myQuestion.value = ""
        })
        .catch(error => {
            console.error(error);
        });
    ElMessage.success(myQuestion.value)
}
const save = function () {
    axios.post("/usr/create-file?token=" + userData.token, {
        path: "chatgpt/",
        fileName: Date.now()+".md",
        content:getContent()
    }).then(res => {
        ElMessage.success("创建成功")
    })
}
onMounted(function () {
    const userStr = localStorage.getItem("user")
    if (userStr != null) {
        const user = JSON.parse(userStr)
        userData.token = user.token
        userData.username = user.username
    }

})

</script>

<style scoped>
input {
    border-radius: 8px;
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

.question-div {
    display: flex;

}

.chat-div {
    top: 0;
    overflow: auto
}

.question-content {
    background-color: #e7f2f8;
    /* 浅灰色 */
    color: #333;
    /* 黑色文本 */
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 10px;
    margin-right: 5px;

}

.answer-div {
    position: relative;
    display: flex;
}

.answer-content {
    background-color: #fff;
    /* 深蓝色 */
    color: #fff;
    /* 白色文本 */
    border-radius: 5px;
    padding: 10px;
    margin-bottom: 10px;
    margin-right: 5px;
}

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



.avator {
    margin: 10px;
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
