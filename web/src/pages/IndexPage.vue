<template>
    <el-container style="height:100%">
        <el-aside width="200px" style="height:100%">
            <el-container width="200px" style="height:100%;background-color:#1f2233">
                <el-header style="color:#fff;height:10%;padding:10px;border: 1px solid;">
                    <el-dropdown>
                        <span class="el-dropdown-link">
                            <div style="display:flex">
                                <div>
                                    <el-avatar> user </el-avatar>
                                </div>
                                <div style="display: flex;align-items: center;margin: 10px;">
                                    {{ username }}
                                </div>
                            </div>
                        </span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>个人主页</el-dropdown-item>
                                <el-dropdown-item @click="logout">退出账号</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </el-header>
                <el-main width="200px" style="height:90%;padding: 0;">
                    <el-menu style="height:90%" width="200px" background-color="#1f2233" text-color="#fff"
                        default-active="1" :collapse="isCollapse">

                        <el-sub-menu index="1">
                            <template #title>
                                <el-icon>
                                    <location />
                                </el-icon>
                                <span>我的</span>
                            </template>
                            <el-menu-item-group title="group One">
                                <el-menu-item index="1-1">item one</el-menu-item>
                                <el-menu-item index="1-2">item two</el-menu-item>
                            </el-menu-item-group>
                            <el-menu-item-group title="Group Two">
                                <el-menu-item index="1-3">item three</el-menu-item>
                            </el-menu-item-group>

                        </el-sub-menu>
                        <el-menu-item index="2">
                            <el-icon><icon-menu /></el-icon>
                            <template #title>书籍管理</template>
                        </el-menu-item>
                        <el-menu-item index="3">
                            <el-icon>
                                <document />
                            </el-icon>
                            <template #title>文件管理</template>
                        </el-menu-item>
                        <el-menu-item index="4">
                            <el-icon>
                                <setting />
                            </el-icon>
                            <template #title>设置</template>
                        </el-menu-item>
                    </el-menu>
                </el-main>
            </el-container>

        </el-aside>
        <el-main style="margin-left: 10px;">
            <RouterView></RouterView>
        </el-main>
    </el-container>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router'
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'
const router = useRouter()
const username = ref<string>("未命名")
const logout=function(){
    localStorage.removeItem("user")
    router.push("/login")
}
onMounted(
    () => {
        const user = localStorage.getItem("user")
        if (!user) {
            router.push("/login")
        } else {
            const userInfo = JSON.parse(user)
            username.value = userInfo.username
        }
    }
)
const isCollapse = ref<boolean>(false)
</script>