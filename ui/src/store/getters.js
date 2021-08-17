
const getters = {
    account: state => state.auth.account,
    qq: state => state.auth.qq,
    loading: state => state.loading,
    message: state => state.message,
    uid: state => state.auth.uid,
}
export default getters;
