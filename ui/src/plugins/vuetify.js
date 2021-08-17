import Vue from 'vue';
import 'vuetify/dist/vuetify.min.css';
import Vuetify, { VSnackbar, VBtn, VIcon } from 'vuetify/lib';
import {LIGHT} from "@/config/theme";

Vue.use(Vuetify);
Vue.component("v-snackbar", VSnackbar);
Vue.component("v-btn", VBtn);
Vue.component("v-icon", VIcon);
const opts = {
    // 添加主题配置
    theme: {
        dark: false, // 如果指定为 true，那么整个页面主题将为下面的 dark 配色
        themes: {
            light: LIGHT,
        },
    },
};
const vuetify = new Vuetify(opts);
// https://github.com/yariksav/vuetify-confirm
import VuetifyConfirm from 'vuetify-confirm';
const cfmOpts = {
    vuetify,
    buttonTrueText: '确定',
    buttonFalseText: '取消',
    buttonTrueColor: 'primary',
    buttonFalseColor: 'grey',
    buttonTrueFlat: false,
    buttonFalseFlat: true,
    color: 'warning',
    //icon: 'warning',
    title: 'Warning',
    width: 350,
    property: '$confirm',
    persistent:true, // 阻止点击外部时自动关闭
};
Vue.use(VuetifyConfirm, cfmOpts);
export default vuetify;