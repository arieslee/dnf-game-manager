import {LIGHT} from "@/config/theme";
const state = {
    text:'',
    show:false,
    color:'',
    closeBtn: true,
    closeBtnType:'icon', // icon or text
    timeout: 5000,
    callback: null,
}

let messageTimer = null;

const mutations = {
    SET_MESSAGE: (state, message) => {
        if(messageTimer){
            clearTimeout(messageTimer);
        }
        state = Object.assign(state, message);
        if(state.show && state.timeout > 0){
            // 自动关闭
            messageTimer = setTimeout(()=>{
                state = Object.assign(state, {
                    show:false,
                });
                if(state.callback && typeof state.callback === 'function'){
                    state.callback(state);
                }
            }, state.timeout);
        }
    },
}


const actions = {
    success({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:LIGHT.success,
                show:true,
            });
        })
    },
    error({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:LIGHT.error,
                show:true,
            });
        })
    },
    warning({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:LIGHT.warning,
                show:true,
            });
        })
    },
    info({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:LIGHT.info,
                show:true,
            });
        })
    },
    hide({commit}){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                show:false,
            });
        });
    }
};


export default {
    namespaced: true,
    state,
    mutations,
    actions
}
