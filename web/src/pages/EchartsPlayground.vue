<template>
    <div>
        <AppTemplate>
            <template #content>
                <div ref="charts" :style="{ width: '900px', height: '300px' }">
            </div>
            </template>
        </AppTemplate>
    </div>
</template>
<script setup lang="ts">
import AppTemplate from '../components/AppTemplate.vue';
import { ref, onMounted, } from 'vue'
import * as echarts from 'echarts'
const canExpand = ref<Boolean>(false)
const expand = function () {
    canExpand.value = true
}
const recover = function () {
    canExpand.value = false
}
const charts = ref<any>(null)
onMounted(() => {
    console.log(charts.value)
    var myChart = echarts.init(charts.value as HTMLElement)
    myChart.setOption({
        title: {
            text: 'ECharts 入门示例'
        },
        tooltip: {},
        xAxis: {
            data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
        },
        yAxis: {},
        series: [
            {
                name: '销量',
                type: 'bar',
                data: [5, 20, 36, 10, 10, 20]
            }
        ]
    });
})
</script>

