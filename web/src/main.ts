import { createApp } from 'vue'

import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router/router'
import ElementPlus from "element-plus";
import axios from 'axios'
import VMdEditor from '@kangc/v-md-editor'
import 'element-plus/theme-chalk/index.css'
import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';


import '@kangc/v-md-editor/lib/style/preview.css';
// VuePress主题以及样式（这里也可以选择github主题）
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
// Prism
import Prism from 'prismjs';
// 代码高亮
import 'prismjs/components/prism-json';
// 选择使用主题
VMdPreview.use(vuepressTheme, {
  Prism,
});
// highlightjs
import hljs from 'highlight.js';

VMdEditor.use(githubTheme, {
  Hljs: hljs,
});
const app=createApp(App)
app.config.globalProperties.$axios=axios
app.use(createPinia()).use(router).use(ElementPlus).use(VMdPreview).use(VMdEditor)
app.mount('#app')
