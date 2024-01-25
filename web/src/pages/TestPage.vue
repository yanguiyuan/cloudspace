<template>
    <AppTemplate>
        <template #content>
            nihao
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

testb.test2()
testa.test1()
console.log(testa.prop1)
</script>
