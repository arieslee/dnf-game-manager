<template>
  <v-app>
      <v-main>
          <v-container fluid fill-height>
              <v-layout class="justify-center">
                  <div class="result">
                      <div class="result-title">404</div>
                      <div class="result-subtitle">抱歉，你访问的页面不存在。</div>
                      <div class="result-extra">
                          <v-btn color="primary" @click="goHome">返回首页</v-btn>
                      </div>
                  </div>
              </v-layout>
          </v-container>
      </v-main>
  </v-app>
</template>
<script>
export default {
  data(){
    return {
      homeUrl:process.env.VUE_APP_INDEX_PATH
    }
  },
    methods:{
      goHome() {
          this.$router.push(this.homeUrl);
      }
    }
}
</script>
<style lang="scss" scoped>
  .result{
    padding-top:100px;
  }
  .result-title{
    color: rgba(0,0,0,.85);
    font-size: 24px;
    line-height: 1.8;
    text-align: center;
  }
  .result-subtitle {
    color: rgba(0,0,0,.45);
    font-size: 14px;
    line-height: 1.6;
    text-align: center;
  }
  .result-extra {
    margin-top: 32px;
    text-align: center;
  }
</style>
