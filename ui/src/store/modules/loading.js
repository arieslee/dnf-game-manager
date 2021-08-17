
const state = {
    text:'',
    show:false,
    color:'primary',
    timeout: 0,
    callback: null,
    size:64,
}

let loadingTimer = null;

const mutations = {
    SET_LOADING: (state, message) => {
        if(loadingTimer){
            clearTimeout(loadingTimer);
        }
        state = Object.assign(state, message);
        if(state.show && state.timeout > 0){
            // 自动关闭
            loadingTimer = setTimeout(()=>{
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
    show({commit}, params){
        return new Promise((resolve, reject) => {
            if(!params){
                params = {
                    show:true,
                };
            }else{
                params = Object.assign(params, {
                    show:true,
                });
            }
            commit('SET_LOADING', params);
        })
    },
    hide({commit}){
        return new Promise((resolve, reject) => {
            commit('SET_LOADING', {
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
