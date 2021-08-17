import axios from 'axios';
import tokenHelper from "@/utils/token";
import vm from '@/main';
// 是否正在刷新的标记
let isRefreshing = false;
let unauthorizedMessage = '身份验证失败，请重新登录！';
// 重试队列，每一项将是一个待执行的函数形式
let retryRequests = [];
const INVALID_REFRESH_TOKEN_CODE = 40002;

const request = axios.create({
    timeout: 50000,
})
// http request 拦截器 Request
request.interceptors.request.use(
    (config) => {
        let accessToken = tokenHelper.get();
        if (accessToken && !config.headers['Authorization']) {
            config.headers['Authorization'] = 'Bearer ' + accessToken;
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// http response 拦截器 Response
request.interceptors.response.use(
    (response) => {
        // code == 0: 成功
        const res = response.data
        const errorCode = Number(res.error);
        if (errorCode !== 0) {
            if (errorCode === 4) {
                // 没有权限
                vm.$router.push('/error/no-authority');
            }
            return Promise.reject(res)
        } else {
            return response.data
        }
    },
    (error) => {
        if (!error.response) return Promise.reject(error);
        let status = Number(error.response.status);
        if (status === 401) {
            // refresh_token失效了，只能重新登录
            if (error.response.data.error === INVALID_REFRESH_TOKEN_CODE) {
                resetLogin();
                return Promise.reject(error);
            }
            const config = error.config
            if (!isRefreshing) {
                isRefreshing = true
                return refreshToken()
                    .then((res) => {
                        let error = Number(res.data.error);
                        if (error !== 0) {
                            resetLogin(res.message);
                            return Promise.reject(res);
                        }
                        // 重新设置token
                        tokenHelper.set(res.data);
                        let token = tokenHelper.get();
                        config.headers['Authorization'] = 'Bearer ' + token;
                        // 已经刷新了token，将所有队列中的请求进行重试
                        // @ts-ignore
                        retryRequests.forEach((cb) => cb(token))
                        // 重试完清空这个队列
                        retryRequests = []
                        // 这边不需要baseURL是因为会重新请求url，url中已经包含baseURL的部分了
                        //config.baseURL = ''
                        return request(config)
                    })
                    .catch((err) => {
                        let error;
                        error = Number(err.error);
                        // 刷新token已失效
                        if (error === INVALID_REFRESH_TOKEN_CODE){
                            err.message = unauthorizedMessage;
                            resetLogin(unauthorizedMessage);
                            return Promise.reject(err);
                        }
                        if(!err.message){
                            err.message = unauthorizedMessage;
                        }
                        return Promise.reject(err);
                    })
                    .finally(() => {
                        isRefreshing = false
                    })
            } else {
                // 正在刷新token，返回一个未执行resolve的promise
                return new Promise((resolve) => {
                    // 将resolve放进队列，用一个函数形式来保存，等token刷新后直接执行
                    // @ts-ignore
                    retryRequests.push((token) => {
                        config.baseURL = ''
                        config.headers['Authorization'] = 'Bearer ' + token;
                        resolve(request(config))
                    })
                })
            }
        } else {
            switch (status) {
                case 400:
                    error.message = '错误请求'
                    break;
                case 403:
                    error.message = '拒绝访问'
                    break;
                case 404:
                    error.message = '请求错误,未找到该资源'
                    break;
                case 405:
                    error.message = '请求方法未允许'
                    break;
                case 408:
                    error.message = '请求超时'
                    break;
                case 500:
                    error.message = '服务器内部错误'
                    break;
                case 501:
                    error.message = '网络未实现'
                    break;
                case 502:
                    error.message = '网络错误'
                    break;
                case 503:
                    error.message = '服务不可用'
                    break;
                case 504:
                    error.message = '网络超时'
                    break;
                case 505:
                    error.message = 'http版本不支持该请求'
                    break;
                default:
                    error.message = `连接错误${error.response.status}`
            }
            return Promise.reject(error)
        }
    }
)
const trimChar = (str, char, type) => {
    if (char) {
        if (type === 'left') {
            return str.replace(new RegExp('^\\' + char + '+', 'g'), '');
        } else if (type === 'right') {
            return str.replace(new RegExp('\\' + char + '+$', 'g'), '');
        }
        return str.replace(new RegExp('^\\' + char + '+|\\' + char + '+$', 'g'), '');
    }
    return str.replace(/^\s+|\s+$/g, '');
}
const getApiUrl = (uri) => {
    let url;
    uri = trimChar(uri, '/');
    if (uri.indexOf(':/') !== -1) {
        url = uri;
    } else {
        url = process.env.VUE_APP_BASE_API + process.env.VUE_APP_API_VERSION + uri;
    }
    return url;
}

// 刷新token的请求方法
function refreshToken() {
    let params = {
        refreshToken: tokenHelper.getRefreshToken() || '',
    };
    let url = process.env.VUE_APP_BASE_API+process.env.VUE_APP_API_VERSION+'refresh-token';
    return axios.post(url, params);
}

function resetLogin(msg = '') {
    if (window.location.href.indexOf('/login') === -1) {
        if (!msg) {
            msg = unauthorizedMessage;
        }
        let opt = {
            title:'系统提示',
            buttonFalseText: null,
            persistent:true,
        };
        vm.$confirm(msg, opt).then(() => {
            // 跳转登录页
            tokenHelper.clean();
            vm.$router.push({
                path: '/login',
                query: {
                    redirectUrl: window.location.href.split('#')[1] || ''
                }
            })
        });
    }
}

export const post = (uri, params, config) => {
    let url = getApiUrl(uri);
    if (config) {
        return request.post(url, params, config)
    }
    return request.post(url, params)
}
export const get = (uri, params) => {
    let url = getApiUrl(uri), config;
    if (params) {
        if (typeof params.params !== 'undefined') {
            config = params;
        } else {
            config = {params: params};
        }
    } else {
        config = null;
    }
    return request.get(url, config)
}
export default request
