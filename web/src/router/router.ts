import {createRouter, createWebHashHistory, createWebHistory} from "vue-router"


const router=createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:"/login",
            component:()=>import('../pages/LoginPage.vue')

        },
        {
            path:"/",
            component:()=>import('../pages/HomePage.vue'),
            children:[
                {
                    path:"",
                    component:()=>import('../pages/AppDesktop.vue')

                },
                {
                    path:"/fileManagement",
                    component:()=>import('../pages/FileManage.vue')
                },
                {
                    path:"/chatgpt",
                    component:()=>import('../pages/ChatGPT.vue')
                },
                {
                    path:"/test",
                    component:()=>import('../pages/TestPage.vue')
                },
                {
                    path:"/3DPlayground",
                    component:()=>import('../pages/3DPlayground.vue')
                },
                {
                    path:"/echarts",
                    component:()=>import('../pages/EchartsPlayground.vue')
                },
                {
                    path:"/map",
                    component:()=>import('../pages/BaiduMapPlayground.vue')
                },
                {
                    path:"/image",
                    component:()=>import('../pages/ImageGenerator.vue')
                },
                {
                    path:"/js",
                    component:()=>import('../pages/Js.vue')
                },
                {
                    path:"/admin",
                    component:()=>import('../pages/AdminPage.vue')
                },
                {
                    path:"/link",
                    component:()=>import('../pages/LinkNamespace.vue')
                }
            ]
        }
    ]
})

export default router