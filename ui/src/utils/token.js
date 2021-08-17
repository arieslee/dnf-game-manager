import storage from '@/utils/storage';
const tokenKey = 'accessToken';
const tokenHelper = {
    set: function (response, callback) {
        if (!response) {
            return false;
        }
        let token = response.token,
            expire = response['expire_unix'] - 120;
        storage.set(tokenKey, token)
        storage.set('accessTokenExpire', expire.toString())
        if (typeof response['refresh_token'] !== 'undefined') {
            let refresh_token = response['refresh_token'],
                refresh_token_expire = response['refresh_token_expire_unix'];
            storage.set('refreshToken', refresh_token)
            storage.set('refreshTokenExpire', refresh_token_expire)
        }
        if (callback && typeof callback === 'function') {
            return callback(response);
        }
    },
    get: function () {
        let timestamp = (new Date()).getTime() / 1000, accessTokenExpire = storage.get('accessTokenExpire');
        if (!accessTokenExpire || accessTokenExpire <= timestamp) {
            return null;
        }
        return storage.get(tokenKey);
    },
    clean: function () {
        storage.remove(tokenKey);
        storage.remove('accessTokenExpire');
        storage.remove('refreshToken');
        storage.remove('refreshTokenExpire');
        storage.remove('userInfo');
    },
    getRefreshToken: function () {
        let timestamp = (new Date()).getTime() / 1000, refreshTokenExpire = storage.get('refreshTokenExpire');
        if (!refreshTokenExpire || refreshTokenExpire <= timestamp) {
            return null;
        }
        return storage.get('refreshToken')
    },
};
export default tokenHelper;
