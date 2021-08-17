
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
                color:'#4CAF50',
                show:true,
            });
        })
    },
    error({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:'#B71C1C',
                show:true,
            });
        })
    },
    warning({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:'#FF8F28',
                show:true,
            });
        })
    },
    info({commit}, text){
        return new Promise((resolve, reject) => {
            commit('SET_MESSAGE', {
                text:text,
                color:'#2392EF',
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
