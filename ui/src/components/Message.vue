<template>
    <v-snackbar top :color="message.color" v-model="message.show">
        {{ message.text }}
        <template v-slot:action="{ attrs }">
            <v-btn v-if="message.closeBtnType==='text'" text v-bind="attrs" @click="hideMessage">关闭</v-btn>
            <v-btn icon v-else v-bind="attrs" @click="hideMessage"><v-icon>mdi-close-circle</v-icon></v-btn>
        </template>
    </v-snackbar>
</template>
<script>
export default {
    computed:{
        message(){
            return this.$store.getters.message;
        },
    },
    methods:{
        hideMessage(){
            this.$store.dispatch('message/hide');
        },
    }
}
</script>