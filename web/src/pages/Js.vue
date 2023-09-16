<template>
    <AppTempleate>
        <template #content>
            <el-table row-click style="border-radius: 5px;width: 100%;" :data="tableData" >
                <el-table-column prop="id" label="序号"/>
                <el-table-column prop="grade" label="等级"  />
                <el-table-column prop="name" label="名称" />
                <el-table-column prop="responsibleGroup" label="负责组"  />
                <el-table-column fixed="right" label="操作" width="120">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="openURL(scope.row.url)">官网</el-button>
                        <el-button link type="primary" size="small" @click="modify(scope.row)">修改</el-button>
                        <el-button link type="primary" size="small" @click="OpenDeleteJsComfirm(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <Drawer title="修改信息" :visible="drawerVisible" @drawerClose="onClose">
                <el-form  label-width="120px">
                    <el-form-item label="序号">
                       {{ currentRow.id }}
                    </el-form-item>
                    <el-form-item label="等级">
                        <el-radio-group v-model="postData.grade">
                            <el-radio label="A" />
                            <el-radio label="B" />
                            <el-radio label="C" />
                            <el-radio label="D" />
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="名称">
                        <el-input v-model="postData.name" />
                    </el-form-item>
                    <el-form-item label="负责组">
                        <el-input v-model="postData.responsibleGroup" />
                    </el-form-item>
                    <el-form-item label="负责同学">
                        <el-input v-model="postData.responsiblePerson" />
                    </el-form-item>
                    <el-form-item label="负责老师">
                        <el-input v-model="postData.ResponsibleTeacher" />
                    </el-form-item>
                    <el-form-item label="URL">
                        <el-input v-model="postData.url" />
                    </el-form-item>
                    <el-form-item>
                    <el-button type="primary" @click="submitForm">
                        修改
                    </el-button>
                    <el-button @click="resetForm">重设</el-button>
                    </el-form-item>
                </el-form>
            </Drawer>
            <Drawer title="添加信息" :visible="maskVisible" @drawerClose="maskClose">
                <el-form  label-width="120px">
                    <el-form-item label="等级">
                        <el-radio-group v-model="postData.grade">
                            <el-radio label="A" />
                            <el-radio label="B" />
                            <el-radio label="C" />
                            <el-radio label="D" />
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="名称">
                        <el-input v-model="postData.name" />
                    </el-form-item>
                    <el-form-item label="负责组">
                        <el-input v-model="postData.responsibleGroup" />
                    </el-form-item>
                    <el-form-item label="负责同学">
                        <el-input v-model="postData.responsiblePerson" />
                    </el-form-item>
                    <el-form-item label="负责老师">
                        <el-input v-model="postData.ResponsibleTeacher" />
                    </el-form-item>
                    <el-form-item label="URL">
                        <el-input v-model="postData.url" />
                    </el-form-item>
                    <el-form-item>
                    <el-button type="primary" @click="SubmitCreateJs">
                        添加
                    </el-button>
                    </el-form-item>
                </el-form>
            </Drawer>
            <YMaskCard :visible="comfirmVisible" v-on:on-close="closeComfirm">
                <div>
                    确定删除?
                </div>
                
                <el-button style="position: absolute;left: 0;right: 0;bottom: 0;margin-left: 50px;margin-right: 50px;" type="primary" @click="deleteJs">
                       确定
                </el-button>
            </YMaskCard>
        </template>
        <template #option>
            <div class="menuItem" @click="createJs">
                <svg t="1687613922291" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2380" width="64" height="64">
                    <path d="M810.666667 128H213.333333c-47.146667 0-85.333333 38.186667-85.333333 85.333333v597.333334c0 47.146667 38.186667 85.333333 85.333333 85.333333h597.333334c47.146667 0 85.333333-38.186667 85.333333-85.333333V213.333333c0-47.146667-38.186667-85.333333-85.333333-85.333333z m-85.333334 426.666667h-170.666666v170.666666h-85.333334v-170.666666h-170.666666v-85.333334h170.666666v-170.666666h85.333334v170.666666h170.666666v85.333334z" fill="#999999" p-id="2381"></path>
                </svg>
                <div class="tooltip">创建</div>
            </div>
        </template>
    </AppTempleate>
</template>
<script lang="ts" setup>
import AppTempleate from '../components/AppTempleate.vue';
import {js} from'../axios/axios.js';
import { reactive,ref } from 'vue';
import YMaskCard from '../components/YMaskCard.vue';
import { ElButton } from 'element-plus';
type Js={
    id:number,
    grade:string,
    name:string,
    responsibleGroup:string,
    responsiblePerson:string,
    ResponsibleTeacher:string,
    url:string,
    CreateTime:string,
    UpdateTime:string,
    html:string
}
const jsDafault={
        grade:"",
        name:"",
        responsibleGroup:"",
        responsiblePerson:"",
        ResponsibleTeacher:"",
        url:"",
        id:0,
        html:"",
        CreateTime:"",
        UpdateTime:""
    }
const drawerVisible=ref<boolean>(false);
const tableData=reactive<Array<Js>>([]);
const currentRow=ref<Js>(tableData[0]);
const postData=ref<Js>(jsDafault);
const maskVisible=ref<boolean>(false)
const comfirmVisible=ref<boolean>(false)
const deleteID=ref<number>(0)
js.get("/js").then(resp=>{
    console.log(resp.data);
    tableData.concat(resp.data);
    var data=resp.data as Array<Js>;
    data.forEach(e => {
        tableData.push(e)
    });
    console.log(tableData.length);
})
const modify=(row:Js)=>{
    drawerVisible.value=true;
    console.log(row);
    currentRow.value=row
    postData.value=row
}
const onClose=()=>{
    drawerVisible.value=false;
}
const submitForm=()=>{
    js.put("/js",postData.value)
}
const resetForm=()=>{
    postData.value=currentRow.value;
}
const createJs=()=>{
    maskVisible.value=true;
}
const maskClose=()=>{
    maskVisible.value=false;
}
const  SubmitCreateJs=async ()=>{
    let res=await js.post("/js",postData.value);
    if (res.data="OK"){
        maskVisible.value=false;
        alert("添加成功");
    }else{
        alert("添加失败");
    }
}
const deleteJs=async ()=>{
    comfirmVisible.value=false;
    let res=await js.delete("/js/"+deleteID.value)
    if (res.data==deleteID.value){
        alert("删除成功");
    }else{
        alert("删除失败")
    }
}
const OpenDeleteJsComfirm=(row:Js)=>{
    comfirmVisible.value=true;
    deleteID.value=row.id
}
const closeComfirm=()=>{
    comfirmVisible.value=false;
}
const openURL=(url:string)=>{
    window.location.href=url
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