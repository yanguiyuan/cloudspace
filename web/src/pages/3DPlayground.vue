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
            </div>
        </div>
        <div ref="space" class="content">
            <div id="three"> </div>
            <div>
                <div><button @click="pause">暂停</button></div>
                <div><button @click="play">播放</button></div>
                <div><button @click="change">切换</button></div>



            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import * as THREE from 'three'
//引入轨道控制器（用来通过鼠标事件控制模型旋转、缩放、移动），没有这个需求可不引入
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls'
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
import { ref, onUnmounted, onMounted, reactive, getCurrentInstance } from 'vue'
import { AnimationAction, AnimationMixer } from 'three';
import { ElMessage } from 'element-plus';
const app = getCurrentInstance()
let el: HTMLElement
let spaceWidth: number = 0
let spaceHeight: number = 0
const canExpand = ref<Boolean>(false)
const expand = function () {
    canExpand.value = true
}
const recover = function () {
    canExpand.value = false
}
const animation = ref<string>("Dog跑步.glb")

//创建一个三维场景
const scene = new THREE.Scene()
scene.background = new THREE.Color(0xffffff);
//添加光源 这里添加了两个光源
//AmbientLight：环境光源，均匀照亮所有物体，防止有些光源照射不到呈现不出来
//PointLight：点光源，类似灯泡发出的光，可以投射阴影，使模型更加立体
const ambientLight = new THREE.AmbientLight(0xffffff, 0.5)
const pointLight = new THREE.PointLight(0xffffff, 0.4)
//设置点光源所在位置
pointLight.position.set(200, 300, 400)

//每创建一个object都需要将其添加到三维场景中
scene.add(ambientLight)
scene.add(pointLight)


//创建一个透视相机 这里的width、height是整个画布的宽高度
const width = window.innerWidth / 2, height = window.innerHeight / 2,
    camera = new THREE.PerspectiveCamera(45, width / height, 1, 1000)
//设置相机位置
camera.position.set(200, 200, 0)
//设置相机方向
camera.lookAt(0, 0, 0)


//创建辅助坐标轴
const axesHelper = new THREE.AxesHelper(100)
scene.add(axesHelper)
//创建一个物体（形状）
const geometry = new THREE.BoxGeometry(10, 10, 10)
//创建材质（外观）
const material = new THREE.MeshLambertMaterial({
    color: 0x00ffff,//设置材质颜色
    transparent: true,//开启透明度
    opacity: 1 //设置透明度
})

//创建一个网格模型对象,将上面设置好的物体及其材质注入到模型对象中
const mesh = new THREE.Mesh(geometry, material)
mesh.position.set(200, 300, 400)
scene.add(mesh)
//加载模型
const loader = new GLTFLoader()
let mixer: AnimationMixer
let action: AnimationAction
let modelScene: THREE.Group
loader.load('/src/assets/' + animation.value, gltf => {
    scene.add(gltf.scene)
    modelScene = gltf.scene
    const aa = gltf.animations[0]
    mixer = new THREE.AnimationMixer(gltf.scene)
    if (aa != undefined) {
        console.log("hh")
        action = mixer.clipAction(aa)
        action.play()
    }
    animate();

})
//创建一个WebGL渲染器
const renderer = new THREE.WebGLRenderer()
console.log(spaceHeight, spaceWidth)
// renderer.setSize(500, 500)
renderer.render(scene, camera)
//vue3的生命周期钩子函数，渲染后执行(这是为了拿到dom节点)
onMounted(() => {
    //将渲染器绘制出的canvas添加到页面中
    document.getElementById('three')?.appendChild(renderer.domElement)
})

//创建轨道控制器
const controls = new OrbitControls(camera, renderer.domElement)
//添加事件监听 当事件发生改变时触发
controls.addEventListener('change', () => {
    //重新渲染
    renderer.render(scene, camera)

})
const clock = new THREE.Clock();
function animate() {
    requestAnimationFrame(animate);
    const deltaTime = clock.getDelta();
    mixer.update(deltaTime);
    renderer.render(scene, camera);
}
const play = function () {
    action.play()
}
const pause = function () {
    action.stop()
    ElMessage.info(action.getClip().duration + "")
}
const reLoadModel = async function () {
    await loader.load('/src/assets/' + animation.value, gltf => {
        scene.add(gltf.scene)
        modelScene = gltf.scene
        const aa = gltf.animations[0]
        mixer = new THREE.AnimationMixer(gltf.scene)
        if (aa != undefined) {
            console.log("hh")
            action = mixer.clipAction(aa)
            action.play()
        }
        animate();

    })
}
const change = function () {
    if(animation.value=="Dog跑步.glb"){
        animation.value = "Dog行走.glb"
    }else if(animation.value=="Dog行走.glb"){
        animation.value="Dog摇尾巴.glb"
    }else if(animation.value=="Dog摇尾巴.glb"){
        animation.value="Dog跑步.glb"
    }
    
    scene.remove(modelScene)
    reLoadModel()
}
onMounted(function () {
    if (app != null) {
        el = app.refs.space as HTMLElement
        spaceWidth = el.offsetWidth
        spaceHeight = el.offsetHeight
        renderer.setSize(spaceHeight, spaceWidth)
    }
})
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
    display: flex;
    justify-content: center;
    align-items: center;
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
}
</style>
  