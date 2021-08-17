<template>
    <v-overlay :value="overlay">
        <v-progress-circular
            indeterminate
            :color="loading.color"
            :size="loading.size"
        ></v-progress-circular>
        <p v-if="loading.text">
            {{loading.text}}
        </p>
    </v-overlay>
</template>

<script>
export default {
    computed:{
        loading(){
            return this.$store.getters.loading;
        },
        overlay(){
            return this.loading.show;
        },
    },
}
</script>