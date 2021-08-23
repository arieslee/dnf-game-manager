import {get,post} from '@/utils/request';
export const getAccountList = (params)=>{
    return get("/account",{params:params});
}
export const getRoleList = (params)=>{
    return get("/role",{params:params});
}
export const addRole = (params) => {
    return post("/role/add",params);
}

export const deleteAccount=(parmas) => {
    return post('account/delete', parmas);
}