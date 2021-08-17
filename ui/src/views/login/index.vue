<template>
    <div class="login-wrapper">
        <div class="login-container">
            <div class="login-content">
                <div class="login-title">管理系统</div>
                <div class="login-form">
                    <v-form ref="login_form">
                        <v-text-field
                            label="登录账号"
                            name="account"
                            prepend-icon="mdi-account"
                            type="text"
                            autocomplete="false"
                            :loading="loginLoading"
                            v-model="loginForm.account"
                            :rules="[rules.required]"
                        ></v-text-field>

                        <v-text-field
                            id="password"
                            label="登录密码"
                            name="password"
                            :loading="loginLoading"
                            prepend-icon="mdi-lock"
                            autocomplete="false"
                            :type="passwordDisplay ? 'text' : 'password'"
                            v-model="loginForm.password"
                            :append-icon="passwordDisplay ? 'mdi-eye' : 'mdi-eye-off'"
                            @click:append="passwordDisplay = !passwordDisplay"
                            :rules="[rules.required]"
                        ></v-text-field>
                    </v-form>
                </div>
                <div class="login-btn">
                    <v-btn :loading="loginLoading" block color="primary" @click="userLogin">登录</v-btn>
                </div>
            </div>
            <p class="copyright"> Copyright © 2019 aries - <a href="https://blog.iw3c.com/">blog.iw3c.com</a></p>
        </div>
    </div>
</template>

<script>
export default{
    created() {
        this.bindEnter();
    },
    data(){
        return {
            passwordDisplay: false,
            loginLoading: false,
            loginForm: {
                account: "aries",
                password: "malshow.com",
                captcha:"",
                req_id:null,
            },
            rules: {
                required: value => !!value || "不能为空."
            },
        };
    },
    methods:{
        userLogin() {
            if (!this.$refs.login_form.validate()) return;
            // 表单验证成功
            this.loginLoading = true;
            this.$store
            .dispatch("auth/loginByJWT", this.loginForm)
            .then(() => {
                this.loginLoading = false;
                this.$store.dispatch('message/success','登录成功');
                console.log('process.env.VUE_APP_INDEX_PATH',process.env.VUE_APP_HOME_PATH);
                this.$router.push(process.env.VUE_APP_HOME_PATH);
            })
            .catch((err) => {
                this.loginLoading = false;
                this.$store.dispatch('message/error', err.message);
            });
        },
        bindEnter() {
            let that = this, key;
            document.onkeydown = function (e) {
                if(that.$route.path !== 'login' && that.$route.path !== '/login'){
                    return ;
                }
                if (typeof window.event === 'undefined') {
                    key = e.keyCode;
                } else {
                    key = window.event.keyCode;
                }
                if (parseInt(key) === 13) {
                    that.userLogin();
                }
            }
        }
    },
}
</script>

<style scoped>
.login-wrapper{
    text-align: center;
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
}
.login-wrapper .login-container {
    width: 320px;
    position: absolute;
    top: 50%;
    left: 50%;
    -webkit-transform: translate(-50%,-50%);
    transform: translate(-50%,-50%);
}
.login-wrapper .login-container .login-content {
    letter-spacing: 2px;
    background: #fff;
    padding: 70px 30px 20px;
    box-shadow: 0 1px 3px 0 rgb(0 0 0 / 6%);
    border-radius: 3px;
    box-sizing: border-box;
}
.login-wrapper .login-container .login-title {
    font-size: 30px;
    color: #3a3a3a;
    line-height: 1;
    margin: -16px 0 40px;
    font-weight: 200;
}
.login-wrapper .login-container .login-form{
    margin-bottom: 45px;
}
.login-wrapper .login-container .login-btn{
    margin-bottom:20px;
    margin-top:20px;
}
.copyright {
    letter-spacing: 1px;
    margin-top: 30px;
    color: #7e7e7e;
}
</style>