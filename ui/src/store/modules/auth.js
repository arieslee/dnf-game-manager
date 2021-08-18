// eslint-disable-next-line no-unused-vars
import {findPwdVerifyAccount, login} from '@/api/auth';
import tokenHelper from '@/utils/token';
import storage from '@/utils/storage';

const state = {
    token: tokenHelper.get(),
    account: '',
    qq: '',
    userInfo: {},
    uid: 0,
}

const mutations = {
    SET_TOKEN: (state, token) => {
        state.token = token
    },
    SET_ACCOUNT: (state, account) => {
        state.account = account
    },
    SET_QQ: (state, qq) => {
        state.qq = qq
    },
    SET_USER_INFO: (state, info) => {
        state.userInfo = info
        storage.set('userInfo', JSON.stringify(info));
    },
    SET_UID: (state, uid) => {
        state.uid = uid;
    },
}

const actions = {
    loginByJWT({commit}, params) {
        return new Promise((resolve, reject) => {
            login(params).then((resp) => {
                tokenHelper.set(resp, function (res) {
                    commit('SET_USER_INFO', res);
                    commit('SET_UID', res['token_uid']);
                    commit('SET_ACCOUNT', res['token_account']);
                    commit('SET_QQ', res['token_qq']);
                    resolve(res);
                });
            }).catch((err) => {
                reject(err);
            })
        });
    },
    setUid({commit}, params){
        commit('SET_UID', params);
    },
    loadUserInfo({commit}, params) {
        return new Promise((resolve, reject) => {
            let userInfo = storage.getUserInfo();
            if(userInfo){
                commit('SET_UID', userInfo['uid']);
                commit('SET_ACCOUNT', userInfo['account']);
                commit('SET_QQ', userInfo['qq']);
            }
        });
    },
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}
