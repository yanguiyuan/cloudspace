<template>
    <AppTempleate>
        <template #content>
            <div id="map"></div>
        </template>
    </AppTempleate>
</template>
<script setup lang="ts">
import { onMounted } from 'vue';
import AppTempleate from '../components/AppTempleate.vue';

onMounted(() => {
    var map = new BMapGL.Map("map");
    var point = new BMapGL.Point(114.38904113196203,30.48748682392412 );
    map.centerAndZoom(point, 15);
    map.enableScrollWheelZoom(true);
    var locationControl = new BMapGL.LocationControl({
        // 控件的停靠位置（可选，默认左上角）
        anchor: BMAP_ANCHOR_TOP_RIGHT,
        // 控件基于停靠位置的偏移量（可选）
        offset: new BMapGL.Size(20, 20)
    });
    // 将控件添加到地图上
    map.addControl(locationControl);
    // 添加定位事件
    locationControl.addEventListener("locationSuccess", function (e: any) {
        console.log(e)
    });
    map.addEventListener('click', function (e: any) {
        alert('点击位置经纬度：' + e.latlng.lng + ',' + e.latlng.lat);
    });
});
</script>
<style>
#map {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
}
</style>