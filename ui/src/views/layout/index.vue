<template>
    <v-app>
        <v-app-bar color="grey darken-3" app dark>
            <v-menu
                bottom
                left>
                <template v-slot:activator="{ on, attrs }">
                    <v-btn dark icon v-bind="attrs" v-on="on">
                        <v-icon>mdi-menu</v-icon>
                    </v-btn>
                </template>

                <v-list>
                    <v-list-item v-for="(item,idx) in menus" :key="'menu_'+idx">
                        <v-list-item-title @click="menuClick(item)">
                            <i v-if="item.icon" :class="item.icon"></i>
                             {{item.label}}
                        </v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
            <v-toolbar-title>{{pageTitle}}</v-toolbar-title>
            <v-spacer></v-spacer>
            <span class="now-time">{{nowTime}}</span>
            <!--头像菜单-->
            <v-menu
                bottom
                min-width="200px"
                rounded
                offset-y>
                <template v-slot:activator="{ on }">
                    <v-btn
                        icon
                        x-large
                        v-on="on">
                        <v-avatar
                            color="white"
                            size="36">
                            <span class="primary--text">{{ cutNameToAvatar(account) }}</span>
                        </v-avatar>
                    </v-btn>
                </template>
                <v-card>
                    <v-list-item-content class="justify-center">
                        <div class="mx-auto text-center">
                            <v-avatar color="primary" size="36">
                                <span class="white--text">{{ cutNameToAvatar(account) }}</span>
                            </v-avatar>
                            <h3 class="pad-ver-mini">{{ account }}</h3>
                            <p class="caption mt-1">
                                {{ account }}
                            </p>
                            <v-divider class="my-3"></v-divider>
                            <v-btn @click="profile" depressed rounded text>
                                个人资料
                            </v-btn>
                            <v-divider class="my-3"></v-divider>
                            <v-btn @click="logout" depressed rounded text>
                                退出登录
                            </v-btn>
                        </div>
                    </v-list-item-content>
                </v-card>
            </v-menu>
            <!--头像菜单-->
        </v-app-bar>
        <v-main>
            <transition name="fade-transform" mode="out-in">
                <router-view :key="key" />
            </transition>
        </v-main>
    </v-app>
</template>
<script>
import tokenHelper from "@/utils/token";
import {MENU} from "@/config/menu";

export default {
    data(){
        return {
            nowTime: '',
            homeUrl: '/',
            menus:MENU,
        }
    },
    computed:{
        key() {
            return this.$route.fullPath;
        },
        qq(){
            return this.$store.getters.qq;
        },
        account(){
            return this.$store.getters.account;
        },
        pageTitle(){
            return process.env.VUE_APP_TITLE;
        },
    },
    mounted() {
        this.nowTimes();
        this.getLocalUserInfo();
    },
    methods:{
        goHome(){
            this.$router.push(this.homeUrl);
        },
        menuClick(item){
            if (typeof item==='string'){
                this.$router.push(path);
            }else{
                if(item.action){
                    if(typeof item.action === 'function'){
                        item.action();
                    }else if(typeof item.action === 'string'){
                        this.$router.push(item.action);
                    }else{
                        this.$store.dispatch('message/error', '菜单链接无效');
                    }
                }else{
                    this.$store.dispatch('message/error', '菜单链接无效');
                }
            }
        },
        getLocalUserInfo(){
            if(!this.account || !this.qq){
                this.$store.dispatch('auth/loadUserInfo');
            }
        },
        async logout(){
            const res = await this.$confirm('确定要退出登录吗？', { title: '系统提示' });
            if (res) {
                tokenHelper.clean();
                this.$store.dispatch('message/success', '已经成功退出登录');
                this.$router.push('/login');
            }
        },
        profile(){
            this.$router.push({
                path:'/profile',
            });
        },
        cutNameToAvatar(name = ''){
            if (!name) return ''
            let len = 1;
            const reg = /[\u4E00-\u9FA5]+/
            const result = name.match(reg) || []
            const str = result[0] || ''
            if (!str[len]) return str
            return str.substr(str.length - len, len)
        },
        timeFormat(timeStamp) {
            let year = new Date(timeStamp).getFullYear();
            let month =new Date(timeStamp).getMonth() + 1 < 10? "0" + (new Date(timeStamp).getMonth() + 1): new Date(timeStamp).getMonth() + 1;
            let date =new Date(timeStamp).getDate() < 10? "0" + new Date(timeStamp).getDate(): new Date(timeStamp).getDate();
            let hh =new Date(timeStamp).getHours() < 10? "0" + new Date(timeStamp).getHours(): new Date(timeStamp).getHours();
            let mm =new Date(timeStamp).getMinutes() < 10? "0" + new Date(timeStamp).getMinutes(): new Date(timeStamp).getMinutes();
            let ss =new Date(timeStamp).getSeconds() < 10? "0" + new Date(timeStamp).getSeconds(): new Date(timeStamp).getSeconds();
            this.nowTime = year + "年" + month + "月" + date +"日"+" "+hh+":"+mm+':'+ss ;
        },
        nowTimes(){
            this.timeFormat(new Date());
            setInterval(this.nowTimes,1000);
            this.clearTime();
        },
        clearTime(){
            clearInterval(this.nowTimes)
            this.nowTimes = '';
        }
    }
}
</script>
<style scoped>

@media screen and (min-width:1200px){
    .now-time{
        padding-right: 10px;
        display: inline;
    }
}
@media (max-width: 767px) {
    .now-time{
        display: none;
    }
}

</style>
