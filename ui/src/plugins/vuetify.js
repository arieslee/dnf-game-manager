import Vue from 'vue';
import 'vuetify/dist/vuetify.min.css';
import Vuetify, { VSnackbar, VBtn, VIcon } from 'vuetify/lib';

Vue.use(Vuetify);
Vue.component("v-snackbar", VSnackbar);
Vue.component("v-btn", VBtn);
Vue.component("v-icon", VIcon);
const opts = {
    // 添加主题配置
    theme: {
        dark: false, // 如果指定为 true，那么整个页面主题将为下面的 dark 配色
        themes: {
            light: {
                primary: '#07c160',
                secondary: '#424242',
                accent: '#82B1FF',
                error: '#FF5252',
                info: '#2196F3',
                success: '#4CAF50',
                warning: '#FFC107',
            },
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