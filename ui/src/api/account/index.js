import {get,post} from '@/utils/request';

export const getRoleList = (params)=>{
    return get("/role");
}