const storage = {
    keyPrefix(key) {
        return process.env.VUE_APP_STORAGE_PREFIX + key;
    },
    get(key) {
        return localStorage.getItem(this.keyPrefix(key));
    },
    set(key, value) {
        let data
        if (typeof value === 'string') {
            data = value;
        } else {
            data = JSON.stringify(value);
        }
        localStorage.setItem(this.keyPrefix(key), data);
    },
    remove(key) {
        return localStorage.removeItem(this.keyPrefix(key));
    },
    getUserInfo() {
        let userInfo = storage.get("userInfo");
        if (userInfo) {
            return JSON.parse(userInfo);
        }
        return null
    },
    getUID(){
        let userInfo = this.getUserInfo();
        if(!userInfo){
            return null;
        }
        return Number(userInfo['token_user_id']);
    },
    getPrivateKey(){
        let userInfo = this.getUserInfo();
        if(!userInfo){
            return null;
        }
        return userInfo['private_key'];
    }
}
export default storage
