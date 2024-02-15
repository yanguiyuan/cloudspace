<template>
    <AppTemplate>
        <template  #header>
          <div class="bg-blue-500">
            <h1>test</h1>
          </div>

        </template>
        <template  #content>
          <div class="bg-amber-500">
            nihao
            <div class="bg-red-500">
                <button @click="testa.test1()">点击</button>
            </div>
            <div class="bg-green-500">
                <button @click="testb.test2()">点击</button>
            </div>
          </div>
          <div>{{getUrlParams(url)}}</div>
        </template>
    </AppTemplate>
</template>
<script lang="ts" setup>
import { reactive } from 'vue';
import AppTemplate from '../components/AppTemplate.vue';
interface TestB{
    one:string
    test2():void
}
const testb=reactive<TestB>({
    one:"王五",
    test2(){
        this.one="ahh"
    }
})
interface TestA{
    prop1:string
    test1():void
}
const testa=reactive<TestA>({
    prop1:"zhangsan",
    test1(){
        this.prop1="李四"
        console.log(testb.one)
        console.log(testa)
    }
})
const url=window.location.href
let URL = "http://www.baidu.com?name=Jack&age=25&sex=men&wife=Lucy"
function getUrlParams(url:string):any {
  let urlStr = url.split('?')[1]
  const urlSearchParams = new URLSearchParams(urlStr)
  const result = Object.fromEntries(urlSearchParams.entries())
  return result
}
console.log(getUrlParams(URL))
testb.test2()
testa.test1()
console.log(testa.prop1)
</script>
