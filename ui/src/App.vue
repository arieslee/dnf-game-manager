<template>
    <div id="app">
        <v-app>
            <router-view/>
            <loading></loading>
            <message></message>
        </v-app>
    </div>
</template>

<script>
import Message from "@/components/Message";
import Loading from "@/components/Loading";
export default {
    name: "App",
    components: {
        Message,
        Loading,
    },
    data() {
        return {
        }
    },
    async created() {
    },
    methods: {
    },
};
</script>
<style>
@import "style/app.scss";
</style>